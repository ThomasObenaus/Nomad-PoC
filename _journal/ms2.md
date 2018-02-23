# First Nomad Cluster (MS 2)

## 2018-02-22

### Started with the Nomad module from terraform-module registry

#### Create an AMI with Consul and Nomad

Docu: https://github.com/hashicorp/terraform-aws-nomad/tree/master/examples/nomad-consul-ami
Note: It is possible to use aws credential profiles for packer

```bash
"builders": [
    {
      "profile": "home",
      ..
    }
    ...
```

ami: ami-adcba8c2

#### Set up the cluster

##### Build the infrastructure

```bash
git clone git@github.com:ThomasObenaus/terraform-aws-nomad.git
cd terraform-aws-nomad
terraform init
terraform plan -out file.plan
terraform apply file.plan
```

##### Show the examples to interact with the nomad-cluster

Define in ```examples/nomad-examples-helper/nomad-examples-helper.sh``` the name of your aws-profile by adjusting ```readonly AWS_PROFILE_NAME="<your profile-name>"```.

```bash
#call the script to show the example nomad instructions
examples/nomad-examples-helper/nomad-examples-helper.sh

# assumption one of your nomad server ip's is: 18.197.84.170

# show the servers
nomad server-members -address=http://18.197.84.170:4646

# open the nomad ui
firefox http://18.197.84.170:4646 &
```

### Installed nomad from https://www.nomadproject.io/intro/getting-started/install.html

Commands for vagrant:

```bash
#enter the folder containing the VagrantFile
cd nomad

#create the vagrant box
vagrant up

#enter the vagrant box
vagrant ssh

#destroy the box afterwards
vagrant destroy
```

Commands for nomad:

```bash
#start nomad agent based on given config (server or client)
#for details see: https://www.nomadproject.io/intro/getting-started/cluster.html
nomad agent -config <xyz>.hcl

#plan job
nomad plan <xyz>.nomad

#apply job
nomad run <xyz>.nomad

#stop job
nomad stop <xyz>.nomad

#show status of job
nomad status <jobname>

#show status of nodes
nomad node-status
```

Nomad UI can be visited at http://localhost:4646/