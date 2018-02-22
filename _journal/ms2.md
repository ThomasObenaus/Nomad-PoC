# First Nomad Cluster (MS 2)

## 2018-02-22

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