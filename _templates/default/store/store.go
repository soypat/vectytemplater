package store

import (
	"github.com/user/vecty-project/store/actions"
	"github.com/user/vecty-project/store/storeutil"
)

var (
	Ctx   actions.Context
	Items []string

	Listeners = storeutil.NewListenerRegistry()
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

	default:
		panic("unknown action selected!")
	}

	Listeners.Fire(action)
}
