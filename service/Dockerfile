FROM golang:alpine AS builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
ADD . /app

# build the service
RUN go build -o service_bin


# nothing needed for golang
FROM alpine

WORKDIR /app

COPY --from=builder /app/service_bin ./service_bin

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Define environment variable
ENV NAME=World

# Run the service
#n -n s1 -consumer :8080
CMD ["echo $NAME"]