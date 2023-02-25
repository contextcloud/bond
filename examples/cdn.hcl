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
  subject_alternative_names = [
    "www.dev01.portal.contextcloud.n-cc.net",
    "*.dev01.portal.contextcloud.n-cc.net"
  ]
  tags = {
    "ManagedBy" = "Bond"
  }
  providers = {
    "aws"     = "aws.non-prod",
    "aws.dns" = "aws.control"
  }
}

resource "aws_cloudfront_distribution" "distribution" {
  origin = [{
    origin_id   = "www"
    domain_name = "dev01.portal.contextcloud.n-cc.net"
    custom_origin_config = {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "match-viewer"
      origin_ssl_protocols   = ["TLSv1", "TLSv1.1", "TLSv1.2"]
    }
    custom_header = [{
      name  = "X-Forwarded-Scheme"
      value = "https"
      }, {
      name  = "X-Frame-Options"
      value = "SAMEORIGIN"
    }]
    origin_shield = {
      enabled              = true
      origin_shield_region = "us-east-1"
    }
  }]

  viewer_certificate = {
    acm_certificate_arn = "arn:aws:acm:us-east-1:578958694144:certificate/9b91e9f3-e772-488b-bdd5-485f51faa38c"
  }

  geo_restriction = {
    restriction_type = "whitelist"
    locations        = ["AU"]
  }

  tags = {
    "ManagedBy" = "Bond"
  }
  
  default_cache_behavior = {
    target_origin_id       = "www"
    viewer_protocol_policy = "allow-all"
    allowed_methods        = ["GET", "HEAD", "OPTIONS"]
    cached_methods         = ["GET", "HEAD"]
    compress               = true
    query_string           = true
  }
}