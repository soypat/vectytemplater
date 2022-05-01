package views

import (
	"vecty-templater-project/app/store/actions"
	"vecty-templater-project/model"

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
	actions.Dispatch(&actions.NewItem{Item: model.Item{Title: val}}) // add new item
	actions.Dispatch(&actions.Back{})                                // Back to previous page.
}

func itemToHTML(e model.Item) *vecty.HTML {
	return elem.Span(
		elem.Strong(vecty.Text(e.Title)),
		elem.Paragraph(vecty.Text(e.Description)),
	)
}
