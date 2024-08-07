resource "aws_default_vpc" "default" {
  tags = {
    Name = "${var.app_name}-${var.environment}-default-vpc"
  }
}

resource "aws_default_subnet" "default_subnet" {
  availability_zone = var.region

  tags = {
    Name = "${var.app_name}-${var.environment}-default-subnet"
  }
}


resource "aws_security_group" "instances" {
  name = "${var.app_name}-${var.environment}-instance-security-group"
}

resource "aws_security_group_rule" "allow_http_inbound" {
  type              = "ingress"
  security_group_id = aws_security_group.instances.id

  from_port   = 8080
  to_port     = 8080
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_route53_zone" "primary" {
  name = "dieg0code-api-players.duckdns.org"
}

resource "aws_acm_certificate" "cert" {
  domain_name       = "dieg0code-api-players.duckdns.org"
  validation_method = "DNS"

  tags = {
    Name = "${var.app_name}-${var.environment}-certificate"
  }
}

resource "aws_route53_record" "cert_validation" {
  for_each = {
    for dvo in aws_acm_certificate.cert.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      type   = dvo.resource_record_type
      value  = dvo.resource_record_value
    }
  }

  name    = each.value.name
  type    = each.value.type
  zone_id = aws_route53_zone.primary.zone_id
  records = [each.value.value]
  ttl     = 60
}

resource "aws_acm_certificate_validation" "cert_validation" {
  certificate_arn         = aws_acm_certificate.cert.arn
  validation_record_fqdns = [for record in aws_route53_record.cert_validation : record.fqdn]
}

resource "aws_lb_listener" "https" {
  load_balancer_arn = aws_lb.load_balancer.arn
  port              = 443
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  certificate_arn   = aws_acm_certificate.cert.arn

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.instances.arn

    fixed_response {
      content_type = "text/plain"
      message_body = "Hello, world"
      status_code  = "200"
    }
  }
}

resource "aws_lb_target_group" "instances" {
  name     = "${var.app_name}-${var.environment}-target-group"
  port     = 8080
  protocol = "HTTP"
  vpc_id   = aws_default_vpc.default.id

  health_check {
    path                = "/health"
    protocol            = "HTTP"
    matcher             = "200"
    interval            = 30
    timeout             = 5
    healthy_threshold   = 2
    unhealthy_threshold = 2
  }
}

resource "aws_lb_target_group_attachment" "api_instance" {
  target_group_arn = aws_lb_target_group.instances.arn
  target_id        = aws_instance.instance_api.id
  port             = 8080
}

resource "aws_lb_listener_rule" "instances" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 100

  condition {
    path_pattern {
      values = ["*"]
    }
  }

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.instances.arn
  }
}

resource "aws_security_group" "lb_sg" {
  name        = "${var.app_name}-${var.environment}-lb-sg"
  description = "Allow HTTPS traffic"
  vpc_id      = aws_default_vpc.default.id

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.app_name}-${var.environment}-lb-sg"
  }
}

resource "aws_lb" "load_balancer" {
  name               = "${var.app_name}-${var.environment}-load-balancer"
  load_balancer_type = "application"
  subnets            = [aws_default_subnet.default_subnet.id]
  security_groups    = [aws_security_group.instances.id]

  enable_deletion_protection = false

  access_logs {
    bucket = aws_s3_bucket.lb_logs.bucket
    prefix = "lb"
    enabled = true
  }
}
