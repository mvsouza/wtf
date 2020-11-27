package view

import (
	"strings"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/wtf"
)

// TextWidget defines the data necessary to make a text widget
type TextWidget struct {
	*Base
	*KeyboardWidget

	View *tview.TextView
}

// NewTextWidget creates and returns an instance of TextWidget
func NewTextWidget(tviewApp *tview.Application, pages *tview.Pages, commonSettings *cfg.Common) TextWidget {
	widget := TextWidget{
		Base:           NewBase(commonSettings),
		KeyboardWidget: NewKeyboardWidget(tviewApp, pages, commonSettings),
	}

	widget.View = widget.createView(widget.bordered)
	widget.View.SetInputCapture(widget.KeyboardWidget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	return widget
}

/* -------------------- Exported Functions -------------------- */

// TextView returns the tview.TextView instance
func (widget *TextWidget) TextView() *tview.TextView {
	return widget.View
}

func (widget *TextWidget) Redraw(data func() (string, string, bool)) {
	widget.tviewApp.QueueUpdateDraw(func() {
		title, content, wrap := data()

		widget.View.Clear()
		widget.View.SetWrap(wrap)
		widget.View.SetTitle(widget.ContextualTitle(title))
		widget.View.SetText(strings.TrimRight(content, "\n"))
	})
}

/* -------------------- Unexported Functions -------------------- */

func (widget *TextWidget) createView(bordered bool) *tview.TextView {
	view := tview.NewTextView()

	view.SetBackgroundColor(wtf.ColorFor(widget.commonSettings.Colors.WidgetTheme.Background))
	view.SetBorder(bordered)
	view.SetBorderColor(wtf.ColorFor(widget.BorderColor()))
	view.SetDynamicColors(true)
	view.SetTextColor(wtf.ColorFor(widget.commonSettings.Colors.TextTheme.Text))
	view.SetTitleColor(wtf.ColorFor(widget.commonSettings.Colors.TextTheme.Title))
	view.SetWrap(false)

	return view
}
