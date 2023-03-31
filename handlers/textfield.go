package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"
)

func TextfieldHandler(w http.ResponseWriter, r *http.Request) {

	// Read in template for textfield page
	t, err := template.ParseFiles("templates/textfield.html")
	if err != nil {
		log.Fatal("error parsing template.")

	}

	// Render the template, writing to `w`.
	t.Execute(w, nil)

	// Done.
	log.Println("Finished HTTP request at", r.URL.Path)
}

func TextfieldMessageHandler(w http.ResponseWriter, r *http.Request) {

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
