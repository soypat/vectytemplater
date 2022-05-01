package model

const (
	HTTPServerAddr = ":8080"
	// Websocket sub protocol.
	WSSubprotocol = "todo"
	// Websocket TCP port number.
	WSServerAddr = ":5757"
)

// Item contains information of a TODO item.
type Item struct {
	Title       string
	Description string
}

type ServerReply struct {
	Info string
}
