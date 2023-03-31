package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	h "github.com/iamrafal1/pushSources/handlers"
)

func main() {
	http.HandleFunc("/textfield", h.TextfieldHandler)
	http.HandleFunc("/textfieldmessage", h.TextfieldMessageHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func testing() {
	for i := 0; ; i++ {
		// Encode the data
		postBody, _ := json.Marshal(map[string]string{
			"time": time.Now().String(),
		})
		err := h.SendRequest("1", "3", postBody)
		if err != nil {
			log.Fatal(err)
		}
		// Print log message and sleep for 5 seconds.
		log.Printf("Sent message %d ", i)
		time.Sleep(5e9)
	}
}
