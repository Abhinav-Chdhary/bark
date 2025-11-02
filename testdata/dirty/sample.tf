# BARK: Remove test configuration before deploy
variable "environment" {
  default = "test"
}

# BARK Use proper secret management
variable "api_key" {
  default = "hardcoded-key-123"
}

resource "aws_instance" "example" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"
  
  # Regular comment
  tags = {
    Name = "ExampleInstance"
    # BARK: Update with production values
    Environment = var.environment
  }
}

# BARK! Remove this test resource
resource "aws_s3_bucket" "test" {
  bucket = "my-test-bucket"
}

