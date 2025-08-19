output "instance_ip" {
  value = aws_instance.microservice_instance.public_ip
}