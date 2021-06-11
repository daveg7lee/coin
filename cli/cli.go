package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/daveg7lee/kangaroocoin/explorer"
	"github.com/daveg7lee/kangaroocoin/rest"
)

// show usage of cli
func usage() {
	fmt.Printf("Welcome to Kangaroo Coin\n\n")
	fmt.Printf("Please use the follow flags:\n\n")
	fmt.Printf("-port:	Set the PORT of the server (default 4000)\n")
	fmt.Printf("-mode:	Choose between 'html', 'rest', and 'all' (rest is recommended)\n")
	runtime.Goexit()
}

// start cli
func Start() {
	// check len of arg
	if len(os.Args) == 1 {
		// show usage is there are no flag
		usage()
	}

	// make flags
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	// parse flags
	flag.Parse()

	// check mode
	switch *mode {
	case "rest":
		// start rest api
		rest.Start(*port)
	case "html":
		// start html explorer
		explorer.Start(*port)
	case "all":
		// start both of them
		go rest.Start(*port + 1000)
		explorer.Start(*port)
	default:
		// show usage if mode is wrong
		usage()
	}
}
