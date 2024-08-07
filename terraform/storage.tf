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

resource "aws_s3_bucket_policy" "lb_logs_policy" {
  bucket = aws_s3_bucket.lb_logs.id

  policy = jsonencode({
    Version = "2012-10-17"
    Id      = "lb_logs_policy"
    Statement = [
      {
        Sid       = "HTTPSOnly"
        Effect    = "Deny"
        Principal = "*"
        Action    = "s3:*"
        Resource = [
          aws_s3_bucket.lb_logs.arn,
          "${aws_s3_bucket.lb_logs.arn}/*",
        ]
        Condition = {
          Bool = {
            "aws:SecureTransport" = "false"
          }
        }
      },
    ]
  })
}
