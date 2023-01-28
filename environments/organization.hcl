id = "cc-organization"

provider "aws" {
  region = "us-east-1"
}

resource "aws_organization" "contextcloud" {
}