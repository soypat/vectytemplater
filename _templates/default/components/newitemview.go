package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/user/vecty-project/store/actions"
	"github.com/user/vecty-project/store/dispatcher"
)

type NewItemView struct {
	vecty.Core
	input *vecty.HTML
}

func (l *NewItemView) Render() vecty.ComponentOrHTML {
	l.input = elem.Input()

	return elem.Form(
		vecty.Markup(
			event.Submit(l.addItem).PreventDefault(),
		),
		l.input,
	)
}

func (l *NewItemView) addItem(e *vecty.Event) {
	val := l.input.Node().Get("value").String()
	dispatcher.Dispatch(&actions.NewItem{Item: val}) // add new item
	dispatcher.Dispatch(&actions.Back{})             // Back to previous page.
}
