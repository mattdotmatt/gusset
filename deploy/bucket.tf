resource "aws_s3_bucket" "my_bucket" {
    bucket_prefix = "my-tf-test-bucket"
    acl    = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
