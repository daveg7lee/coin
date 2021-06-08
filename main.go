package main

import (
	"github.com/daveg7lee/kangaroocoin/explorer"
	"github.com/daveg7lee/kangaroocoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
