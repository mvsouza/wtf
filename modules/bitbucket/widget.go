package bitbucket

import (
	"fmt"
	"github.com/rivo/tview"

	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget
	settings *Settings

	prs []PullRequest
	err      error
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		TextWidget: view.NewTextWidget(app, pages, settings.common),

		settings: settings,
	}

	widget.View.SetScrollable(true)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	return widget
}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	prs, err := widget.GetPrs()

	if err != nil {
		widget.err = err
		widget.prs = nil
	} else {
		widget.prs = prs
	}

	widget.display()
}

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {

	title := fmt.Sprintf("%s/%s PRs", widget.settings.workspace, widget.settings.repo)

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	prs := widget.prs
	if len(prs) == 0 {
		return title, "No PRs to display", false
	}

	var str string
	for _, pr := range prs {
		str += fmt.Sprintf(
			`[green][%s]. %s [lightblue](%s -> %s)[white]`,
			pr.author.display_name,
			pr.title,
			pr.source.branch,
			pr.target.branch,
		)
		str += "\n"

	}

	widget.View.SetWrap(false)
	return title, str, false
}

