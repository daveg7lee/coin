package p2p

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
