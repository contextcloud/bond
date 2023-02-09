id = "bond-organization"

provider "aws" {
  region = "us-east-1"
}

resource "aws_organization" "bond" {
}

resource "aws_organizational_unit" "cloud_unit" {
  organization_name = "cloud"
  accounts = [{
    "name"  = "cloud-control"
    "email" = "aws+cloud-control@getnoops.com"
    }, {
    "name"  = "cloud-nonprod"
    "email" = "aws+cloud-nonprod@getnoops.com"
    }, {
    "name"  = "cloud-prod"
    "email" = "aws+cloud-prod@getnoops.com"
  }]
  depends_on = ["module.bond"]
}