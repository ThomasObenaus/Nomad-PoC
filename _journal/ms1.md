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