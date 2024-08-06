resource "aws_instance" "instance_api" {
  ami                         = var.ami_id
  instance_type               = var.instance_type
  security_groups             = [aws_security_group.instances.name]
  associate_public_ip_address = false

  tags = {
    Name = "${var.app_name}-${var.environment}-instance"
  }
}
