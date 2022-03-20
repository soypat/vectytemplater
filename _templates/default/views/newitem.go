package views

import (
	"vecty-templater-project/store/actions"
	"vecty-templater-project/store/dispatcher"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
)

type NewItem struct {
	vecty.Core
	input *vecty.HTML
}

func (l *NewItem) Render() vecty.ComponentOrHTML {
	l.input = elem.Input()

	return elem.Form(
		vecty.Markup(
			// PreventDefault prevents the Form from navigating away from page.
			event.Submit(l.addItem).PreventDefault(),
		),
		elem.Label(vecty.Text("Press enter to add item.")),
		l.input,
	)
}

func (l *NewItem) addItem(e *vecty.Event) {
	val := l.input.Node().Get("value").String()
	if val == "" {
		return // do not add empty items.
	}
	dispatcher.Dispatch(&actions.NewItem{Item: val}) // add new item
	dispatcher.Dispatch(&actions.Back{})             // Back to previous page.
}
