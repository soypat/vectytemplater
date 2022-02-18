package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/user/vecty-project/dispatcher"
	"github.com/user/vecty-project/store"
	"github.com/user/vecty-project/store/actions"
)

type Body struct {
	vecty.Core
	Ctx actions.Context `vecty:"prop"`
}

func (b *Body) Render() vecty.ComponentOrHTML {
	var mainContent vecty.MarkupOrChild
	switch b.Ctx.Page {
	case actions.Landing:
		mainContent = &LandingView{
			Items: store.Items,
		}

	}
	return elem.Body(
		vecty.If(b.Ctx.Referrer != nil, elem.Button(
			vecty.Markup(event.Click(b.backButton)),
			vecty.Text("Back"),
		)),
		mainContent,
	)
}

func (b *Body) backButton(*vecty.Event) {
	dispatcher.Dispatch(&actions.Back{})
}
