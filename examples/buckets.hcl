id = "bond-buckets"

provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "test-bucket" {
  bucket_name = "contextcloud-bond-test-bucket"
}

resource "aws_s3_bucket" "content-bucket" {
  bucket_name = "contextcloud-bond-content-bucket"
}