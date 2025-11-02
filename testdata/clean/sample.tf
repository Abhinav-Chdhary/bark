# Production Terraform configuration
variable "environment" {
  description = "Environment name"
  type        = string
  default     = "production"
}

variable "api_key" {
  description = "API key from secret manager"
  type        = string
  sensitive   = true
}

resource "aws_instance" "app" {
  ami           = var.ami_id
  instance_type = var.instance_type
  
  # Standard tagging
  tags = {
    Name        = "ProductionInstance"
    Environment = var.environment
    ManagedBy   = "Terraform"
  }
}

