# NW Infrastructure and Services (MS1)

## 2018-02-21

### Dockerized the ping_service

Now you can simply build and run the ping_service using docker.

#### Build the docker image and start a container

```bash
cd ping_service
docker build -t ping_service  .

# maps the services internal port 8080 to 80
docker run -p 80:8080 ping_service
```

#### Start a container with parameters

```bash
docker run -p 80:8080 ping_service -service_name s1 -consumer :8080
```

#### Chainging the ping_services

```bash
export DOCKER_HOST_IP=$(ip route|awk '/docker/ { print $9 }') &&\
docker run -d -p 80:8080 ping_service -service_name s1 -provider $DOCKER_HOST_IP:81 &&\
docker run -d -p 81:8080 ping_service -service_name s2 -provider $DOCKER_HOST_IP:82 &&\
docker run -d -p 82:8080 ping_service -service_name s3
```

#### Cleanup

```bash
docker stop $(docker container ls | awk '/ping_service/{print $1}') &&\
docker rm $(docker container ls -a | awk '/ping_service/{print $1}')
```

### Added chaining to ping_service

Now it is possible to connect several instances of the ping_service with each another.

Example: Service_1 -> Service_2 -> Service_3.

* The left side of each ping_service combination is seen as consumer and the right side as provider (consumer->provider). So in the example above Service_1 is the consumer of Service_2, which is the consumer of Service_3.

#### Build the ping_service

```bash
cd ping_service
# build the service
go build -o ping_service_bin
```

#### Start multiple instances of them s1->s2->s3

```bash
./ping_service_bin -service_name s1 -consumer :8080 -provider localhost:8081 & \
./ping_service_bin -service_name s2 -consumer :8081 -provider localhost:8082 & \
./ping_service_bin -service_name s3 -consumer :8082
```

#### Watch the service-chain

```bash
#in another terminal, use the service
watch -x curl -s localhost:8080/ping
# you should see something like {"message":"/[s1,v1]/[s2,v1]/[s3,v1](PONG)","name":"s1","version":"v1"}
```

#### For cleanup call

```bash
killall -e ping_service_bin
```

## 2018-02-10

### Added basic networking setup with terraform

```bash
cd infra
terraform init
terraform plan
terraform apply
```

### Added simple service in golang

```bash
cd service
# build and start service
go build -o ping_service_bin && ./ping_service_bin

#in another terminal, use the service
watch -x curl -s localhost:8080/ping
```
