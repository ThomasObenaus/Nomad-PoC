package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type pingResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type PingService struct {
	Name         string
	Version      string
	ProviderName string
	ConsulClient Client
}

const maxHop = 10

func (s *PingService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var message string

	hopStr := r.URL.Query().Get("hop")
	hop, err := strconv.Atoi(hopStr)
	if err != nil {
		hop = 0
	}
	hop++

	if len(s.ProviderName) == 0 || hop > maxHop {
		message = "(PONG)"
	} else {
		var err error
		message, err = s.getMessage(s.ProviderName, hop)
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

func (s *PingService) getMessage(providerName string, hop int) (string, error) {

	addr, err := s.ConsulClient.FindProvider(providerName)
	if err != nil {
		return "(BOING)", err
	}

	url := "http://" + addr + "/ping?hop=" + strconv.Itoa(hop)
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
