package main

import (
	"github.com/daveg7lee/kangaroocoin/cli"
	"github.com/daveg7lee/kangaroocoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
