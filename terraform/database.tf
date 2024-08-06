resource "aws_db_instance" "db_instance" {
  storage_encrypted       = true
  allocated_storage       = 20
  storage_type            = "gp2"
  engine                  = "postgres"
  engine_version          = "12.5"
  instance_class          = "db.t2.micro"
  db_name                 = var.db_name
  username                = var.db_user
  password                = var.db_password
  skip_final_snapshot     = true
  backup_retention_period = 7
  multi_az                = false
  publicly_accessible     = false

  tags = {
    Name        = "MyRDSInstance"
    Environment = "Development"
  }
}
