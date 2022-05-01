package main

import (
	"context"
	"log"
	"net/http"
	"time"
	"vecty-templater-project/model"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type todoServer struct {
	list []model.Item
}

// ServeHTTP is a basic websocket implementation for reading/writing a TODO list
// from a websocket.
func (s *todoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols:       []string{model.WSSubprotocol},
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	if c.Subprotocol() != model.WSSubprotocol {
		c.Close(websocket.StatusPolicyViolation, "client must speak the "+model.WSSubprotocol+" subprotocol")
		return
	}

	for {
		err = s.getTODO(r.Context(), c)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}

		if err != nil {
			log.Printf("failed to echo with %v: %v\n", r.RemoteAddr, err)
			return
		}
	}
}

// getTODO reads from the WebSocket connection and then writes
// the received message into the server's TODO list.
// The entire function has 200s to complete.
func (t *todoServer) getTODO(ctx context.Context, c *websocket.Conn) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*200)
	defer cancel()
	reply := &model.ServerReply{
		Info: "I got the data!",
	}
	err := wsjson.Write(ctx, c, reply)
	if err != nil {
		return err
	}

	log.Printf("got TODO list:\n%s\n", t.list)
	ctx, cancel = context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	err = wsjson.Read(ctx, c, &t.list)
	if err != nil {
		return err
	}
	log.Println("data exchange succesful with reply: ", reply)
	return nil
}
