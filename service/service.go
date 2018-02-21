package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type pingResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Service struct {
	Name            string
	Version         string
	ProviderAddress string
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var message string

	if len(s.ProviderAddress) == 0 {
		message = "(PONG)"
	} else {
		var err error
		message, err = s.getMessage(s.ProviderAddress)
		if err != nil {
			log.Println("[" + s.Name + "," + s.Version + "] Error: " + err.Error())
		}
	}

	message = "/[" + s.Name + "," + s.Version + "]" + message
	response := &pingResponse{
		Message: message,
		Name:    s.Name,
		Version: s.Version,
	}
	log.Printf("[%s,%s] responds: %s\n", s.Name, s.Version, message)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Service) getMessage(providerAddress string) (string, error) {
	url := "http://" + providerAddress + "/ping"
	client := &http.Client{
		Timeout: time.Second * 1,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "(BOING)", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "(BOING)", err
	}
	defer resp.Body.Close()

	response := &pingResponse{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return "(BOING)", err
	}

	message := response.Message
	return message, nil
}
