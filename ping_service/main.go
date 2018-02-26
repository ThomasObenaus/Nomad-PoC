package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/namsral/flag"
)

const version = "v1"

func registerService(consul Client, serviceName string, port int) {
	for {
		if err := consul.Register(serviceName, port); err != nil {
			log.Printf("Error unable to register %s at consul: %s\n", serviceName, err.Error())
			log.Println("Waiting for 20 sec")
			time.Sleep(time.Second * 20)
		} else {
			log.Println("Sucessfully registered")
			break
		}
	}
}

func main() {

	var portOfConsumer = flag.Int("consumer", 8080, "The port where the consumer shall listen to (this application instance). Defaults to 8080.")
	var serviceName = flag.String("service_name", "foo", "The name of the consumer service instance (this application instance). Defaults to foo.")
	var nameOfProvider = flag.String("provider", "", "The service_name of the provider (another instance of this application). Defaults to \"\".")
	var addrOfProvider = flag.String("provider_addr", "", "The address of the provider (another instance of this application). Defaults to \"\".")
	var addrOfConsul = flag.String("consul", "127.0.0.1:8500", "The addr of the consul server. Defaults to 127.0.0.1:8500.")
	flag.Parse()

	consul, err := NewConsulClient(*addrOfConsul)
	if err != nil {
		log.Println("Error unable to create consul at: ", *addrOfConsul)
	}

	// register at consul in background
	go registerService(consul, "ping-service", *portOfConsumer)

	http.Handle("/ping", &PingService{Name: *serviceName, ProviderAddr: *addrOfProvider, ProviderName: *nameOfProvider, Version: version, ConsulClient: consul})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Path '/' is not implemented")
		http.Error(w, "Path '/' is not implemented", http.StatusInternalServerError)
	})

	//start the web server
	log.Printf("%s starts listening at %d.\n", *serviceName, *portOfConsumer)

	provider := *nameOfProvider
	if len(provider) == 0 {
		provider = *addrOfProvider
	}

	if len(provider) > 0 {
		log.Printf("The provider at %s is used.\n", provider)
	} else {
		log.Println("No provider is used.")
	}
	if err := http.ListenAndServe(":"+strconv.Itoa(*portOfConsumer), nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}

	log.Println("Exiting")
}
