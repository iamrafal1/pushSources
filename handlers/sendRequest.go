package handlers

import (
	"bytes"
	"log"
	"net/http"
)

// Helper function for sending a request
func SendRequest(key string, token string, body []byte) error {
	responseBody := bytes.NewBuffer(body)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/top", responseBody)
	if err != nil {
		log.Print("Failed to create request")
	}
	req.Header.Add("Push-Key", key)
	req.Header.Add("Push-Token", token)

	// Post request to server
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}
