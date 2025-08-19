variable "aws_region" {
  default = "us-east-1"
}

variable "instance_type" {
  default = "t2.micro"
}

variable "ami_id" {
  default = "ami-0c94855ba95c71c99"
}

variable "key_name" {
  description = "AWS Key Pair name for SSH access"
  default     = "my-key-pair"
}
