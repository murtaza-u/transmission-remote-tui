package tui

import "github.com/rivo/tview"

type Navigation struct {
    widget *tview.Table
}

func (nav *Navigation) setHeaders() {
    var headers []string = []string { "Overview", "Files", "Peers", "Trackers" }
    for col, header := range headers {
        nav.widget.SetCell(0, col, tview.NewTableCell(header).SetExpansion(1).SetAlign(tview.AlignCenter))
    }
}

func initNavigation() *Navigation {
    return &Navigation{
        widget: tview.NewTable().SetSelectable(false, true).SetFixed(1, 1),
    }
}
