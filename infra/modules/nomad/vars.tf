variable "nomad_ami_id" {
  description = "The ID of the AMI to be used for the nomad nodes."
}

variable "consul_ami_id" {
  description = "The ID of the AMI to be used for the consul nodes."
}

variable "vpc_id" {
  description = "Id of the vpc where to place in the instances."
}

variable "nomad_server_subnet_ids" {
  description = "Ids of the subnets to deploy the nomad servers into."
  type        = "list"
}

variable "aws_region" {
  description = "The AWS region to deploy into (e.g. us-east-1)."
  default     = "eu-central-1"
}

variable "nomad_cluster_name" {
  description = "What to name the Nomad cluster and all of its associated resources"
  default     = "nomad-example"
}

variable "consul_cluster_name" {
  description = "What to name the Consul cluster and all of its associated resources"
  default     = "consul-example"
}

variable "num_nomad_servers" {
  description = "The number of Nomad server nodes to deploy. We strongly recommend using 3 or 5."
  default     = 3
}

variable "num_nomad_clients" {
  description = "The number of Nomad client nodes to deploy. You can deploy as many as you need to run your jobs."
  default     = 3
}

variable "num_consul_servers" {
  description = "The number of Consul server nodes to deploy. We strongly recommend using 3 or 5."
  default     = 3
}

variable "cluster_tag_key" {
  description = "The tag the Consul EC2 Instances will look for to automatically discover each other and form a cluster."
  default     = "consul-servers"
}

variable "ssh_key_name" {
  description = "The name of an EC2 Key Pair that can be used to SSH to the EC2 Instances in this cluster. Set to an empty string to not associate a Key Pair."
  default     = ""
}
