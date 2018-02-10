package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

type pingResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

const version = "v1"

func main() {

	var addr = flag.String("addr", ":8080", "The addr of the application")
	var serviceName = flag.String("n", "foo", "The name of this service instance")
	flag.Parse()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

		response := &pingResponse{
			Message: "pong",
			Name:    *serviceName,
			Version: version,
		}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	//start the web server
	log.Println("Start listening at ", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}

	log.Println("Exiting")
}
