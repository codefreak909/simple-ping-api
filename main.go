package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// This defines the endpoint /ping
	http.HandleFunc("/ping", pingHandler)

	// Starting HTTP server at port 8000
	http.ListenAndServe(":8000", nil)
}

// This functions contains what a ping has to do
func pingHandler(w http.ResponseWriter, req *http.Request) {
	// Checking for GET request
	if req.Method == "GET" {
		// Extracting the query parameter from URL /ping?name=whatever
		name := req.URL.Query().Get("name")

		// Returns reponse to request
		fmt.Fprintf(w, "pong! Mr."+name)

		// logging URL Path
		log.Println(req.URL.Path)
	} else if req.Method == "POST" {
		// Getting data from POST request body
		decoder := json.NewDecoder(req.Body)

		// Structure for the POST request json body
		type test_struct struct {
			Name string
		}

		var t test_struct

		// Decoding JSON into test struct
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "post pong! Mr."+t.Name)
	} else {
		// This is what it does from every other request method
		fmt.Fprintf(w, "not pong!")
	}
}
