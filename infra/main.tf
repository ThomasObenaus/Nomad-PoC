provider "aws" {
  profile = "${var.deploy_profile}"
  region  = "${var.region}"
}

module "networking" {
  source = "networking"
  region = "${var.region}"
}

module "nomad" {
  source = "nomad"
  region = "${var.region}"
}