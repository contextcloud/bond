id = "bond-certs"

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