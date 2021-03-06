package digitalclock

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget is a text widget struct to hold info about the current widget
type Widget struct {
	view.TextWidget

	app      *tview.Application
	settings *Settings
}

// NewWidget creates a new widget using settings
func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, nil, settings.Common),

		app:      app,
		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	widget.display()
}
