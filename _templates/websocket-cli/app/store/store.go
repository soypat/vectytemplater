package store

import (
	"vecty-templater-project/app/store/actions"
	"vecty-templater-project/model"
)

var (
	Ctx         actions.Context
	Items       []model.Item
	ServerReply string
	Listeners   = newListenerRegistry()
)

func OnAction(action interface{}) {
	switch a := action.(type) {
	case *actions.NewItem:
		Items = append(Items, a.Item)

	case *actions.PageSelect:
		oldCtx := Ctx
		Ctx = actions.Context{
			Page:     a.Page,
			Referrer: &oldCtx,
		}

	case *actions.Back:
		Ctx = *Ctx.Referrer
	case *actions.Refresh:
		// do nothing, just fire listeners to refresh page.
	default:
		panic("unknown action selected!")
	}

	Listeners.Fire(action)
}
