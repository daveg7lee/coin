package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/daveg7lee/kangaroocoin/explorer"
	"github.com/daveg7lee/kangaroocoin/rest"
)

func usage() {
	fmt.Printf("Welcome to Kangaroo Coin\n\n")
	fmt.Printf("Please use the follow flags:\n\n")
	fmt.Printf("-port:	Set the PORT of the server (default 4000)\n")
	fmt.Printf("-mode:	Choose between 'html', 'rest', and 'all' (rest is recommended)\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "all":
		go rest.Start(*port + 1000)
		explorer.Start(*port)
	default:
		usage()
	}
}
