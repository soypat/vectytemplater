package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/hexops/vecty"
	"github.com/user/vecty-project/components"
	"github.com/user/vecty-project/dispatcher"
	"github.com/user/vecty-project/store"
	"github.com/user/vecty-project/store/actions"
)

func main() {
	// OnAction must be registered before any storage manipulation.
	dispatcher.Register(store.OnAction)

	attachItemsStorage()
	body := &components.Body{
		Ctx: store.Ctx,
	}
	store.Listeners.Add(body, func(interface{}) {
		body.Ctx = store.Ctx // Renew context on rerender.
		vecty.Rerender(body)
	})
	vecty.RenderBody(body)
}

// attachItemsStorage provides persistent local storage saved on edits so
// no data is lost due to bad connection or refreshed page.
func attachItemsStorage() {
	const key = "items"
	store.Listeners.Add(nil, func(action interface{}) {
		if _, ok := action.(*actions.NewItem); !ok {
			// Only save state upon adding an item
			return
		}
		// After item addition save state.
		b, err := json.Marshal(&store.Items)
		if err != nil {
			panic(err)
		}
		js.Global().Get("localStorage").Set(key, string(b))
	})

	if data := js.Global().Get("localStorage").Get(key); !data.IsUndefined() {
		// Old session data found, initialize store data.
		err := json.Unmarshal([]byte(data.String()), &store.Items)
		if err != nil {
			panic(err)
		}
	}
}
