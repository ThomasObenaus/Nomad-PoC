# NW Infrastructure and Services (MS1)

## 2018-02-10

* added basic networking setup with terraform

```bash
cd infra
terraform init
terraform plan
terraform apply
```

* added simple service in golang

```bash
cd service
# build and start service
go build -o service_bin && ./service_bin

#in another terminal, use the service
watch -x curl -s localhost:8080/ping
```

## 2018-02-21

* added chaining to service

Now it is possible to connect several instances of the service with each another.

Example: Service_1 -> Service_2 -> Service_3.

* The left side of each service combination is seen as consumer and the right side as provider (consumer->provider). So in the example above Service_1 is the consumer of Service_2, which is the consumer of Service_3.

Build the service

```bash
cd service
# build the service
go build -o service_bin
```

Start multiple instances of them s1->s2->s3

```bash
./service_bin -n s1 -consumer :8080 -provider localhost:8081 & \
./service_bin -n s2 -consumer :8081 -provider localhost:8082 & \
./service_bin -n s3 -consumer :8082
```

Watch the service-chain

```bash
#in another terminal, use the service
watch -x curl -s localhost:8080/ping
# you should see something like {"message":"/[s1,v1]/[s2,v1]/[s3,v1](PONG)","name":"s1","version":"v1"}
```

For cleanup call

```bash
killall -e service_bin
```
