resource "aws_s3_bucket" "b" {
  bucket = vars.bucket_name

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}