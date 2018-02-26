package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/namsral/flag"
)

const version = "v1"

func registerService(consul Client, serviceName string, address string, port int) {
	for {
		if err := consul.Register(serviceName, address, port); err != nil {
			log.Printf("Error unable to register %s (%s:%d) at consul: %s\n", serviceName, address, port, err.Error())
			log.Println("Waiting for 20 sec")
			time.Sleep(time.Second * 20)
		} else {
			log.Printf("Sucessfully registered %s %s:%d\n", serviceName, address, port)
			break
		}
	}
}

func main() {

	var localPort = flag.Int("p", 8080, "The port where the application instance listens to. Defaults to 8080.")
	var serviceName = flag.String("service-name", "foo", "The name of the consumer service instance (this application instance). Defaults to foo.")
	var nameOfProvider = flag.String("provider", "", "The service_name of the provider (another instance of this application). Defaults to \"\".")
	var addrOfProvider = flag.String("provider-addr", "", "The address of the provider (another instance of this application). Defaults to \"\".")
	var registrationIP = flag.String("registration-ip", "127.0.0.1", "The ip to register this service at consul. Defaults to 127.0.0.1")
	var registrationPort = flag.Int("registration-port", 8080, "The port to register this service at consul. Defaults to 8080")
	var registerAsService = flag.String("registration-name", "ping-service", "The name that is used to register at consul. Defaults to ping-service.")
	var addrOfConsul = flag.String("consul-server-addr", "127.0.0.1:8500", "The addr of the consul-server. Defaults to 127.0.0.1:8500.")
	flag.Parse()

	consul, err := NewConsulClient(*addrOfConsul)
	if err != nil {
		log.Println("Error unable to create consul at: ", *addrOfConsul)
	}

	// register at consul in background
	go registerService(consul, *registerAsService, *registrationIP, *registrationPort)

	http.Handle("/ping", &PingService{Name: *serviceName, ProviderAddr: *addrOfProvider, ProviderName: *nameOfProvider, Version: version, ConsulClient: consul})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Path '/' is not implemented")
		http.Error(w, "Path '/' is not implemented", http.StatusInternalServerError)
	})

	//start the web server
	log.Printf("%s starts listening at %d.\n", *serviceName, *localPort)

	provider := *nameOfProvider
	if len(provider) == 0 {
		provider = *addrOfProvider
	}

	if len(provider) > 0 {
		log.Printf("The provider at %s is used.\n", provider)
	} else {
		log.Println("No provider is used.")
	}
	if err := http.ListenAndServe(":"+strconv.Itoa(*localPort), nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}

	log.Println("Exiting")
}
