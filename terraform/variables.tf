# General variables

variable "region" {
  description = "AWS region"
  type        = string
  default     = "sa-east-1"
}

variable "app_name" {
  description = "Name of the app"
  type        = string
  default     = "player-profile"
}

variable "environment" {
  description = "Deployment environment"
  type        = string
  default     = "dev"
}

# EC2 variables

variable "ami_id" {
  description = "Amazon machine image to use in EC2 instance"
  type        = string
  default     = "ami-080111c1449900431"  # Ubuntu 20.04 LTS sa-east-1
}

variable "instance_type" {
  description = "Type of instance to create"
  type        = string
  default     = "t2.micro"
}

# RDS variables

variable "db_name" {
  description = "Name of the RDS database"
  type        = string
}

variable "db_user" {
  description = "Username for the RDS database"
  type        = string
}

variable "db_password" {
  description = "Password for the RDS database"
  type        = string
  sensitive   = true
}