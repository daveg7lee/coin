package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/daveg7lee/kangaroocoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func handleHome(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.gohtml"))
	blocks := blockchain.GetBlockchain().AllBlocks()
	data := homeData{"Home", blocks}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", handleHome)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
