id = "bond-cdn"

provider "aws" {
  region = "us-east-1"
}

resource "aws_cloudfront_distribution" "distribution" {
}