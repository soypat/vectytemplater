package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type LandingView struct {
	vecty.Core
	Items []string `vecty:"prop"`
}

func (l *LandingView) Render() vecty.ComponentOrHTML {
	var items vecty.List
	for _, item := range l.Items {
		items = append(items, elem.ListItem(vecty.Text(item)))
	}
	return elem.UnorderedList(items)
}
