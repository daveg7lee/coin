package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daveg7lee/kangaroocoin/blockchain"
	"github.com/daveg7lee/kangaroocoin/utils"
	"github.com/gorilla/mux"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	// make url's form like http://localhost${port}${url}
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	// return url
	return []byte(url), nil
}

// Struct to make description of routes
type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

type addBlockBody struct {
	Data string
}

// handle '/' route
func documentation(rw http.ResponseWriter, r *http.Request) {
	// make description of route '/'
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "GET",
			Description: "Get a Block",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{id}"),
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
		var addBlockBody addBlockBody
		// get data from body and decode to go value
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		data := addBlockBody.Data
		// add block to blockchain
		blockchain.GetBlockchain().AddBlock(data)
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
}

func Start(portNum int) {
	// make own handler
	router := mux.NewRouter()
	// init port number
	port = fmt.Sprintf(":%d", portNum)
	// handle routes
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{id:[0-9]+}", block).Methods("GET")
	// run server
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
