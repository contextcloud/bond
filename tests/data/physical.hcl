id = "cc-cloud"

provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "flow-logs" {
  bucket_name = "contextcloud-bond-flow-logs"
}

resource "aws_vpc" "non-prod" {
  cidr_block  = ""
  flow_logs   = "contextcloud-bond-flow-logs"
}

resource "aws_eks" "non-prod" {
}

resource "aws_route53_zone" "non-prod" {
  name = "non-prod.example.com"
}