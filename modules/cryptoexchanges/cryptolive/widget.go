package cryptolive

import (
	"fmt"
	"sync"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/modules/cryptoexchanges/cryptolive/price"
	"github.com/wtfutil/wtf/modules/cryptoexchanges/cryptolive/toplist"
	"github.com/wtfutil/wtf/view"
)

// Widget define wtf widget to register widget later
type Widget struct {
	view.TextWidget

	priceWidget   *price.Widget
	toplistWidget *toplist.Widget
	settings      *Settings
}

// NewWidget Make new instance of widget
func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		priceWidget:   price.NewWidget(settings.priceSettings),
		toplistWidget: toplist.NewWidget(settings.toplistSettings),
		settings:      settings,
	}

	widget.priceWidget.RefreshInterval = widget.RefreshInterval()
	widget.toplistWidget.RefreshInterval = widget.RefreshInterval()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh & update after interval time
func (widget *Widget) Refresh() {
	var wg sync.WaitGroup

	wg.Add(2)
	widget.priceWidget.Refresh(&wg)
	widget.toplistWidget.Refresh(&wg)
	wg.Wait()

	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	str := ""
	str += widget.priceWidget.Result
	str += widget.toplistWidget.Result

	return widget.CommonSettings().Title, fmt.Sprintf("\n%s", str), false
}
