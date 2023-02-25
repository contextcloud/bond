id = "bond-domains"

provider "aws" {
  region = "us-east-1"

  assume_role {
    role_arn = "arn:aws:iam::497099817429:role/OrganizationAccountAccessRole"
  }
}


resource "aws_route53_zones" "zones" {
  zones = [{
    name    = "contextcloud.io",
    comment = "Root zone for contextcloud.io",
  }]
  tags = {
    "ManagedBy" = "Bond"
  }
}

resource "aws_route53_records" "records" {
  zone_name = "contextcloud.io"
  records = [{
    name    = "orgs.contextcloud.io",
    type    = "NS",
    comment = "Delegate to portal",
    ttl     = 300,
    records = [
      "ns-1158.awsdns-16.org",
      "ns-1537.awsdns-00.co.uk",
      "ns-289.awsdns-36.com",
      "ns-645.awsdns-16.net"
    ]
  }]
  tags = {
    "ManagedBy" = "Bond"
  }
  depends_on = ["zones"]
}