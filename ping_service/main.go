package main

import (
	"log"
	"net/http"

	"github.com/namsral/flag"
)

const version = "v1"

func main() {

	var addrOfConsumer = flag.String("consumer", ":8080", "The addr of the consumer (this application instance)")
	var serviceName = flag.String("service_name", "foo", "The name of the consumer service instance (this application instance)")
	var addrOfProvider = flag.String("provider", "", "The addr of the provider (another instance of this application)")
	flag.Parse()

	http.Handle("/ping", &PingService{Name: *serviceName, ProviderAddress: *addrOfProvider, Version: version})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Path '/' is not implemented")
		http.Error(w, "Path '/' is not implemented", http.StatusInternalServerError)
	})

	//start the web server
	log.Printf("%s starts listening at %s.\n", *serviceName, *addrOfConsumer)

	if len(*addrOfProvider) > 0 {
		log.Printf("The provider at %s is used.\n", *addrOfProvider)
	} else {
		log.Println("No provider is used.")
	}
	if err := http.ListenAndServe(*addrOfConsumer, nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}

	log.Println("Exiting")
}
