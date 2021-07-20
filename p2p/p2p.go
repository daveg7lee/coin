package p2p

import (
	"fmt"
	"net/http"

	"github.com/daveg7lee/kangaroocoin/utils"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	for {
		_, payload, err := conn.ReadMessage()
		utils.HandleErr(err)
		fmt.Printf("%s\n", payload)
	}
}
