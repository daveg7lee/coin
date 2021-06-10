package explorer

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/daveg7lee/kangaroocoin/blockchain"
)

// struct to send data to page
type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

const templateDir string = "explorer/templates/"

var templates *template.Template

// handle '/'
func handleHome(rw http.ResponseWriter, r *http.Request) {
	// get all blocks
	blocks := blockchain.Blockchain().AllBlocks()
	// make data
	data := homeData{"Home", blocks}
	// execute template "home" with data
	templates.ExecuteTemplate(rw, "home", data)
}

// handle '/add'
func handleAdd(rw http.ResponseWriter, r *http.Request) {
	// check request's method
	switch r.Method {
	case "POST":
		// get input value
		r.ParseForm()
		data := r.Form.Get("blockData")
		// add block with data
		blockchain.Blockchain().AddBlock(data)
		// redirect to home
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	case "GET":
		// render template 'add'
		templates.ExecuteTemplate(rw, "add", nil)
	}
}

func Start(port int) {
	//make own handler
	handler := http.NewServeMux()
	// init templates
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	// handle routes
	handler.HandleFunc("/", handleHome)
	handler.HandleFunc("/add", handleAdd)
	// run server
	fmt.Printf("HTML Explorer is Listening on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
