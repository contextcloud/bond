id = "bond-cdn"

provider "aws" {
  region = "us-east-1"

  assume_role {
    role_arn = "arn:aws:iam::578958694144:role/OrganizationAccountAccessRole"
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
    acm_certificate_arn = "arn:aws:acm:us-east-1:024171573596:certificate/12c9337b-185a-4665-b369-40706c46f13e"
  }
}