env {
  AWS_ACCESS_KEY_ID     = "..."
  AWS_SECRET_ACCESS_KEY = "..."
  AWS_SESSION_TOKEN     = "..."
}

provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "test-bucket" {
  bucket_name = "contextcloud-bond-test-bucket"
}

resource "aws_s3_bucket" "content-bucket" {
  bucket_name = "contextcloud-bond-content-bucket"
}