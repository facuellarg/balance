resource "aws_s3_bucket" "transactions" {
  bucket = local.stori_bucket
  tags = {
    Name        = "transactions-bucket"
    Environment = "dev"
  }
}


resource "aws_s3_object" "object" {
  bucket = aws_s3_bucket.transactions.id
  key    = local.stori_file
  source = "${path.module}/../${local.stori_file}}"
  etag   = filemd5("${path.module}/../${local.stori_file}")

}
