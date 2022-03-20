package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type Landing struct {
	vecty.Core
	Items []string `vecty:"prop"`
}

func (l *Landing) Render() vecty.ComponentOrHTML {
	var items vecty.List
	for _, item := range l.Items {
		items = append(items, elem.ListItem(vecty.Text(item)))
	}
	return elem.UnorderedList(items)
}
