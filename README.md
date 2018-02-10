# Set up a Nomad COS

## Goals

* Set up Nomad-Cluster.
* Deploy 3 simple dummy services talking with each other.
* Support:
  * Logging with fluentd and ELK.
  * Monitoring with Prometheus and Grafana.
  * Tracing with Zipkin.
  * Persistant storage.
  * RBAC 
  * SecretManagement with Vault.
  * ServiceDiscovery with Consul.
* Evaluate/ implement envoy/ istio.

## Preconditions

* Infrastructure parts should be defined with terraform.
* On AWS.
* Setup on top of networking infrastructure (terraform) that is already in place.

## Milestones

1. __MS1__ NW Infrastructure and Services.

  * Set up networking infrastructure in terraform.
  * Create the sample service.

2. __MS2__ First Nomad Cluster.
3. __MS3__ Deploy Sample Services.
4. __MS4__ ServiceDiscovery with Envoy and Consul.