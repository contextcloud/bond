id = "bond-organization"

provider "aws" {
  region = "us-east-1"
}

resource "aws_organization" "bond" {
}

resource "aws_organizational_unit" "cloud" {
  organization_name = "cloud"
}

resource "aws_organization_account" "cloud_control" {
  name                = "cloud-control"
  email               = "aws+cloud-control@getnoops.com"
  organizational_unit = aws_organizational_unit.cloud.id
}