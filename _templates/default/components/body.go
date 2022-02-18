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
	case actions.PageLanding:
		mainContent = elem.Div(
			elem.Div(elem.Button(
				vecty.Markup(event.Click(b.newItem)),
				vecty.Text("New item"),
			)),
			&LandingView{
				Items: store.Items,
			},
		)
	case actions.PageNewItem:
		mainContent = &NewItemView{}
	default:
		panic("unknown Page")
	}
	return elem.Body(
		vecty.If(b.Ctx.Referrer != nil, elem.Div(
			elem.Button(
				vecty.Markup(event.Click(b.backButton)),
				vecty.Text("Back"),
			))),
		mainContent,
	)
}

func (b *Body) backButton(*vecty.Event) {
	dispatcher.Dispatch(&actions.Back{})
}

func (b *Body) newItem(*vecty.Event) {
	dispatcher.Dispatch(&actions.PageSelect{Page: actions.PageNewItem})
}
