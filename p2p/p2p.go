package p2p

import (
	"net/http"

	"github.com/daveg7lee/kangaroocoin/utils"
	"github.com/gorilla/websocket"
)

var conns []*websocket.Conn
var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	conns = append(conns, conn)
	utils.HandleErr(err)
	for {
		_, payload, err := conn.ReadMessage()
		if err != nil {
			break
		}
		for _, aConn := range conns {
			if aConn != conn {
				utils.HandleErr(aConn.WriteMessage(websocket.TextMessage, payload))
			}
		}
	}
}
