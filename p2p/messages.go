package p2p

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/daveg7lee/kangaroocoin/blockchain"
	"github.com/daveg7lee/kangaroocoin/utils"
)

type MessageKind int

const (
	NewestBlockMessage MessageKind = iota
	AllBlocksRequestMessage
	AllBlockResponseMessage
	NewBlockNotifyMessage
	NewTxNotifyMessage
	NewPeerNotifyMessage
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
	fmt.Printf("Sending newest block to %s\n", p.key)
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	m := makeMessage(NewestBlockMessage, b)
	p.inbox <- m
}

func requestAllBlocks(p *peer) {
	m := makeMessage(AllBlocksRequestMessage, nil)
	p.inbox <- m
}

func sendAllBlocks(p *peer) {
	m := makeMessage(AllBlockResponseMessage, blockchain.Blocks(blockchain.Blockchain()))
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case NewestBlockMessage:
		fmt.Printf("Received the newest block from %s\n", p.key)
		var payload blockchain.Block
		json.Unmarshal(m.Payload, &payload)
		b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
		utils.HandleErr(err)
		if payload.Height >= b.Height {
			fmt.Printf("Requesting all blocks from %s\n", p.key)
			requestAllBlocks(p)
		} else {
			fmt.Printf("Sending newest block to %s\n", p.key)
			sendNewestBlock(p)
		}
	case AllBlocksRequestMessage:
		fmt.Printf("%s wants all the blocks\n", p.key)
		sendAllBlocks(p)
	case AllBlockResponseMessage:
		fmt.Printf("Received all the blocks from %s\n", p.key)
		var payload []*blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		blockchain.Blockchain().Replace(payload)
	case NewBlockNotifyMessage:
		var payload *blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		blockchain.Blockchain().AddPeerBlock(payload)
	case NewTxNotifyMessage:
		var payload *blockchain.Tx
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		blockchain.Mempool().AddPeerTx(payload)
	case NewPeerNotifyMessage:
		var payload string
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		parts := strings.Split(payload, ":")
		AddPeer(parts[0], parts[1], parts[2], false)
	}
}

func notifyNewBlock(b *blockchain.Block, p *peer) {
	m := makeMessage(NewBlockNotifyMessage, b)
	p.inbox <- m
}

func notifyNewTx(tx *blockchain.Tx, p *peer) {
	m := makeMessage(NewTxNotifyMessage, tx)
	p.inbox <- m
}

func notifyNewPeer(address string, p *peer) {
	m := makeMessage(NewPeerNotifyMessage, address)
	p.inbox <- m
}
