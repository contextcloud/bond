id = "bond-domains"

provider "aws" {
  region = "us-east-1"

  assume_role {
    role_arn = "arn:aws:iam::497099817429:role/OrganizationAccountAccessRole"
  }
}

resource "aws_route53_zones" "zones_root" {
  zones = [{
    name    = "contextcloud.io",
    comment = "Root zone for contextcloud.io",
  }]
  tags = {
    "ManagedBy" = "Bond"
  }
  providers = {
    "aws.ns" = "aws"
  }
}

resource "aws_route53_zones" "zones_orgs" {
  zones = [{
    name         = "orgs.contextcloud.io",
    comment      = "Root zone for contextcloud.io",
    ns_zone_name = "contextcloud.io"
  }]
  tags = {
    "ManagedBy" = "Bond"
  }
  providers = {
    "aws.ns" = "aws"
  }
  depends_on = ["zones_root"]
}