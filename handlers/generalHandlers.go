package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"
)

func HandlerWrapper(htmlFile string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Read in template for textfield page
		t, err := template.ParseFiles(htmlFile)
		if err != nil {
			log.Fatal("error parsing template.")

		}

		// Render the template, writing to `w`.
		t.Execute(w, nil)

		// Done.
		log.Println("Finished HTTP request at", r.URL.Path)
	}
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	// make it separate for random and for textfield somehow
	// Read in data
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		w.Write([]byte(err.Error()))
		return
	}
	if string(b) == "" {
		log.Print("No Data")
		w.Write([]byte("No Data!"))
		return
	}
	// Put data into json
	postBody, err := json.Marshal(map[string]string{
		"message": string(b),
	})
	if err != nil {
		log.Print(err)
		w.Write([]byte(err.Error()))
		return
	}

	// Send the request to the server
	err = SendRequest("1", "3", postBody)
	if err != nil {
		log.Print(err)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Success"))
}

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
