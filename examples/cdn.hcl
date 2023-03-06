id = "bond-cdn"

provider "aws" {
  alias  = "control"
  region = "us-east-1"

  assume_role {
    role_arn = "arn:aws:iam::024171573596:role/OrganizationAccountAccessRole"
  }
}

provider "aws" {
  alias  = "non-prod"
  region = "us-east-1"

  assume_role {
    role_arn = "arn:aws:iam::578958694144:role/OrganizationAccountAccessRole"
  }
}

resource "aws_acm" "cdn_certs" {
  zone_id     = "Z0135217102XSUXM7CWGH"
  domain_name = "dev01.portal.contextcloud.n-cc.net"
  tags = {
    "ManagedBy" = "Bond"
  }
  providers = {
    "aws"     = "aws.non-prod",
    "aws.dns" = "aws.control"
  }
}

resource "aws_cloudfront_distribution" "distribution" {
  aliases = ["dev01.portal.contextcloud.n-cc.net"]
  origin = [{
    origin_id   = "edge"
    domain_name = "edge.nonprod.contextcloud.n-cc.net"
    custom_origin_config = {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "match-viewer"
      origin_ssl_protocols   = ["TLSv1", "TLSv1.1", "TLSv1.2"]
    }
    origin_shield = {
      enabled              = true
      origin_shield_region = "us-east-1"
    }
  },{
    origin_id   = "www"
    s3_origin_config = {
      origin_access_identity = "origin-access-identity/cloudfront/E1XZQXQXQXQXQ"
    }
  }]

  default_cache_behavior = {
    allowed_methods = ["GET", "HEAD", "OPTIONS"]
    cached_methods  = ["GET", "HEAD"]
    cache_policy = {
      query_string_behavior = "all"
      header_behavior       = "none"
      cookie_behavior       = "all"
    }
    compress               = true
    target_origin_id       = "edge"
    viewer_protocol_policy = "allow-all"
  }
  
  ordered_cache_behavior = [{
    path_pattern         = "/*"
    allowed_methods = ["GET", "HEAD", "OPTIONS"]
    cached_methods  = ["GET", "HEAD"]
    cache_policy = {
      query_string_behavior = "all"
      header_behavior       = "none"
      cookie_behavior       = "all"
    }
    compress               = true
    target_origin_id       = "www"
    viewer_protocol_policy = "allow-all"
  }]

  geo_restriction = {
    restriction_type = "none"
  }

  tags = {
    "ManagedBy" = "Bond"
  }
  providers = {
    "aws"     = "aws.non-prod"
  }
}