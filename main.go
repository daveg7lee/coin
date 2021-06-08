package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daveg7lee/kangaroocoin/blockchain"
	"github.com/daveg7lee/kangaroocoin/utils"
)

// port number
const port string = ":4000"

type URL string

func (u URL) MarshalText() ([]byte, error) {
	// make url's form like http://localhost${port}${url}
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	// return url
	return []byte(url), nil
}

// Struct to make description of routes
type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type AddBlockBody struct {
	Data string
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
		{
			URL:         URL("/blocks/{id}"),
			Method:      "GET",
			Description: "See a Block",
		},
	}
	// add content-type to header
	rw.Header().Add("Content-Type", "application/json")
	// encode data to JSON and write to response
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	// check request type
	switch r.Method {
	case "GET":
		// add content-type to header
		rw.Header().Add("Content-Type", "application/json")
		// encode blocks' data to json
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		// make var to store data from user
		var addBlockBody AddBlockBody
		// get data from body and decode to go value
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		data := addBlockBody.Data
		// add block to blockchain
		blockchain.GetBlockchain().AddBlock(data)
		rw.WriteHeader(http.StatusCreated)
	}
}

func main() {
	// handle routes
	http.HandleFunc("/", documentation)
	http.HandleFunc("/blocks", blocks)
	// run server
	fmt.Printf("Server running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
