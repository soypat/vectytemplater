package main

import (
	"context"
	"encoding/json"
	"fmt"
	"syscall/js"
	"time"

	"vecty-templater-project/app/store"
	"vecty-templater-project/app/store/actions"
	"vecty-templater-project/model"

	"vecty-templater-project/app/views"

	"github.com/hexops/vecty"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {
	// OnAction must be registered before any storage manipulation.
	actions.Register(store.OnAction)

	attachItemsStorage()

	body := &views.Body{
		Ctx:  store.Ctx,
		Info: "Welcome!",
	}
	store.Listeners.Add(body, func(interface{}) {
		body.Ctx = store.Ctx
		body.Info = store.ServerReply
		vecty.Rerender(body)
	})
	vecty.RenderBody(body)
}

// attachItemsStorage provides persistent local storage saved on edits so
// no data is lost due to bad connection or refreshed page.
func attachItemsStorage() {
	const key = "vecty_items"
	defer initWSConn()
	store.Listeners.Add(nil, func(action interface{}) {
		if _, ok := action.(*actions.NewItem); !ok {
			// Only save state upon adding an item
			return
		}
		store.ServerReply = "" // reset server reply
		// After item addition save items locally.
		b, err := json.Marshal(&store.Items)
		if err != nil {
			panic(err)
		}
		js.Global().Get("localStorage").Set(key, string(b))
		var reply model.ServerReply
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		err = wsjson.Read(ctx, wsConn, &reply)
		if err != nil {
			fmt.Println("reading server reply from websocket:", err)
			go initWSConn()
			return
		}
		// store items remotely too.
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = wsjson.Write(ctx, wsConn, &store.Items)
		if err != nil {
			fmt.Println("writing items to websocket:", err)
			go initWSConn()
			return
		}
		store.ServerReply = reply.Info
		actions.Dispatch(&actions.Refresh{})
	})

	if data := js.Global().Get("localStorage").Get(key); !data.IsUndefined() {
		// Old session data found, initialize store data.
		err := json.Unmarshal([]byte(data.String()), &store.Items)
		if err != nil {
			panic(err)
		}
	}
}

var wsConn *websocket.Conn

func initWSConn() {
	if wsConn != nil {
		wsConn.Close(websocket.StatusAbnormalClosure, "client wanted to reinitialize")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://localhost"+model.HTTPServerAddr+"/ws", &websocket.DialOptions{
		Subprotocols: []string{model.WSSubprotocol},
	})
	if err != nil {
		fmt.Println("websocket initialization failed:", err.Error())
	}
	wsConn = c
}
