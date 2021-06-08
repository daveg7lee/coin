package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// port number
const port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

// Struct to make description of routes
type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

// handle '/' route
func documentation(rw http.ResponseWriter, r *http.Request) {
	// make description of route '/'
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
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
	fmt.Printf("Server running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
