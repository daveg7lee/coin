package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/daveg7lee/kangaroocoin/blockchain"
)

const (
	port        string = ":4000"
	templateDir string = "templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func handleHome(rw http.ResponseWriter, r *http.Request) {
	blocks := blockchain.GetBlockchain().AllBlocks()
	data := homeData{"Home", blocks}
	templates.ExecuteTemplate(rw, "home", data)
}

func main() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", handleHome)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
