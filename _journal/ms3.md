# Service Discovery with Consul HTTP-API (MS 3)

## 2018-02-28

Consul provides two ways for service-discovery.

1. DNS Interface.

    * If Consul is configured as dns-server it's entries (services) can be obtained over usual dns-requests.
    * Advantage:
      * Includes loadbalancing. Consul knows about the state of the registered services (i.e. health) and automatically balances between them as needed.
      * In the service itself a dns name can be used, so it's transparent to the service.
    * Disadvantage:
      * The port has to be prior known. Or two DNS queries, one for the address and one for the port has to be made.
      * Single point of failure.

2. HTTP API

    * Consul provides an http-api that can be used to query the provider-service address, port, health-status etc.
    * Based on this information the service can decide to which provider to call next.
    * Advantage:
      * No DNS problems involved (caching, propagation, etc.).
    * Disadvantage:
      * Direct couplig in the service to consul.
      * No loadbalancing provided. The decisicion to which provider-instance to call has to be made in the service itself.

### Implemented service-discovery over Consul HTTP-API

In general it's no good idea to implement the service-discovery logic using the consul HTTP API directly in the service. But anyway I'll try it here.

#### Start a consul-server and ping_services (with docker-compose)

```bash
cd consul
docker-compose up
```

#### Start a consul-server and -client (separate calls)

```bash
cd consul

# Create folders for consul-data
mkdir -p /tmp/consul_data_dir_client &&\
mkdir -p /tmp/consul_data_dir_server

# Start server and client
consul agent -config-file=server.json &\
sleep 5 &&\
consul agent -config-file=client.json
```

#### Start an instance and register at consul SD

```bash
go build -o ping_service && ./ping_service -consul-server-addr 127.0.0.1:8500
```

#### Query consul informations

* Consul-Web UI: http://127.0.0.1:8500
* Consul-Port for DNS queries: 8600

```bash
# list servers and clients
consul members
consul info
```

#### Clean up

```bash
killall consul &&\
rm -rf /tmp/consul_data_dir_client &&\
rm -rf /tmp/consul_data_dir_server
```