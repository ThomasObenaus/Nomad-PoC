variable "deploy_profile" {
  description = "Specify the local AWS profile configuration to use."
  default     = "home"
}

variable "aws_region" {
  description = "The AWS region to deploy into (e.g. us-east-1)."
  default     = "eu-central-1"
}
