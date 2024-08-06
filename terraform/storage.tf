resource "aws_s3_bucket" "lb_logs" {
  bucket = "${var.app_name}-${var.environment}-lb-logs"

  tags = {
    Name        = "${var.app_name}-${var.environment}-lb-logs"
    Environment = var.environment
  }

}

resource "aws_s3_bucket_public_access_block" "lb_logs_public_access_block" {
  bucket = aws_s3_bucket.lb_logs.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_versioning" "lb_bucket_versioning" {
  bucket = aws_s3_bucket.lb_logs.id
  versioning_configuration {
    status     = "Enabled"
    mfa_delete = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "lb_bucket_crypto_conf" {
  bucket = aws_s3_bucket.lb_logs.bucket
  rule {
    apply_server_side_encryption_by_default {
      sse_algorithm = "AES256"
    }
  }
}
