provider "aws" {
  profile = "${var.deploy_profile}"
  region  = "${var.aws_region}"
}

module "nomad" {
  source     = "../../modules/nomad"
  aws_region = "${var.aws_region}"
}
