package main

import (
	"encoding/json"
	"log"
	"time"

	h "github.com/iamrafal1/pushSources/handlers"
)

func main() {
	// http.HandleFunc("/textfield", h.HandlerWrapper("templates/textfield.html"))
	// http.HandleFunc("/random", h.HandlerWrapper("templates/random.html"))
	// http.HandleFunc("/textfieldmessage", h.TextfieldMessageHandler)
	// log.Fatal(http.ListenAndServe("localhost:9090", nil))
	testing()
}

func testing() {
	for i := 0; ; i++ {
		// Encode the data
		postBody, _ := json.Marshal(map[string]string{
			"message": time.Now().String(),
		})
		err := h.SendRequest("10361b536", "36a349f1aed7a5e2767d37908b2a489", postBody)
		if err != nil {
			log.Fatal(err)
		}
		// Print log message and sleep for 5 seconds.
		log.Printf("Sent message %d ", i)
		time.Sleep(5e9)
	}
}
