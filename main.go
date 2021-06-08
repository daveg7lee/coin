package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// port number
const port string = ":4000"

// Struct to make description of routes
type URLDescription struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

// handle '/' route
func documentation(rw http.ResponseWriter, r *http.Request) {
	// make description of route '/'
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         "/blocks",
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
	}
	// add content-type to header
	rw.Header().Add("Content-Type", "application/json")
	// encode data to JSON and write to response
	json.NewEncoder(rw).Encode(data)
}

func main() {
	// handle route
	http.HandleFunc("/", documentation)
	// run server
	log.Fatal(http.ListenAndServe(port, nil))
	fmt.Printf("Server running on http://localhost%s\n", port)
}
