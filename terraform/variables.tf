variable "aws_region" {
  description = "The AWS region to deploy resources"
  type        = string
  default     = "us-east-1"
}

variable "instance_type" {
  description = "The AWS instance type to use"
  type        = string
  default     = "t3.micro"
}

variable "ami_id" {
  description = "The AWS AMI to use"
  type        = string
  default     = "ami-0c94855ba95c71c99"
}

variable "key_name" {
  description = "AWS Key Pair name for SSH access"
  type        = string
  default     = "microservice-key-pair"
}
