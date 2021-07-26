package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/daveg7lee/kangaroocoin/blockchain"
	"github.com/daveg7lee/kangaroocoin/utils"
)

type MessageKind int

const (
	NewestBlockMessage MessageKind = iota
	AllBlockRequestMessage
	AllBlockResponseMessage
)

type Message struct {
	Kind    MessageKind
	Payload []byte
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message{
		Kind:    kind,
		Payload: utils.ToJSON(payload),
	}
	return utils.ToJSON(m)
}

func sendNewestBlock(p *peer) {
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(NewestBlockMessage, b)
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case NewestBlockMessage:
		var payload blockchain.Block
		json.Unmarshal(m.Payload, &payload)
		fmt.Println(payload)
	}
}
