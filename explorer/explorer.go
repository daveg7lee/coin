package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/daveg7lee/kangaroocoin/blockchain"
)

// data struct to send data to page
type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

const (
	port        string = ":4000"
	templateDir string = "explorer/templates/"
)

var templates *template.Template

// handle '/'
func handleHome(rw http.ResponseWriter, r *http.Request) {
	// get all blocks
	blocks := blockchain.GetBlockchain().AllBlocks()
	// make data
	data := homeData{"Home", blocks}
	// execute template "home" with data
	templates.ExecuteTemplate(rw, "home", data)
}

// handle '/add'
func handleAdd(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	}
}

func Start() {
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/add", handleAdd)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
