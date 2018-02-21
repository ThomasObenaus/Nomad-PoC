################### BUILDER IMAGE ################################
FROM golang:alpine AS builder

# install git
RUN set -ex &&\ 
  apk update &&\ 
  apk add --no-cache git

# set up go dep
RUN go get -u github.com/golang/dep/cmd/dep

COPY Gopkg.lock Gopkg.toml /go/src/service/
WORKDIR /go/src/service/
# Install library dependencies
RUN dep ensure -vendor-only

# copy service files and build the service
COPY . /go/src/service/
RUN go build -o /bin/service

#################### SERVICE IMAGE ################################
## nothing needed for golang
FROM alpine
COPY --from=builder /bin/service /bin/service

ENTRYPOINT ["/bin/service"]