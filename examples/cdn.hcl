id = "bond-cdn"

provider "aws" {
  region = "us-east-1"

  assume_role {
    role_arn = "arn:aws:iam::024171573596:role/OrganizationAccountAccessRole"
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
}

resource "aws_cloudfront_distribution" "distribution" {
  viewer_certificate = "arn:aws:acm:us-east-1:024171573596:certificate/2598e4ef-b9d8-4f22-88c1-ee2f94575bfd"
}