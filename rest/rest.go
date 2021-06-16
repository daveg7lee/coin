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

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

type balanceResponse struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
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
			URL:         url("/status"),
			Method:      "GET",
			Description: "See the Status of the Blockchain",
		},
		{
			URL:         url("/balance/{address}"),
			Method:      "GET",
			Description: "See the Balance of the Address",
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
			URL:         url("/blocks/{hash}"),
			Method:      "GET",
			Description: "See a Block",
		},
	}
	// encode data to JSON and write to response
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	// check request type
	switch r.Method {
	case "GET":
		// encode blocks' data to json
		json.NewEncoder(rw).Encode(blockchain.Blockchain().Blocks())
	case "POST":
		blockchain.Blockchain().AddBlock()
		rw.WriteHeader(http.StatusCreated)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	block, err := blockchain.FindBlock(hash)
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

func status(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(blockchain.Blockchain())
}

func balance(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	total := r.URL.Query().Get("total")
	switch total {
	case "true":
		amount := blockchain.Blockchain().BalanceByAddress(address)
		utils.HandleErr(json.NewEncoder(rw).Encode(balanceResponse{address, amount}))
	default:
		utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.Blockchain().TxOutsByAddress(address)))
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(portNum int) {
	// make own handler
	router := mux.NewRouter()
	// init port number
	port = fmt.Sprintf(":%d", portNum)
	// use middleware
	router.Use(jsonContentTypeMiddleware)
	// handle routes
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/balance/{address}", balance).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods("GET")
	// run server
	fmt.Printf("REST API is Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
