package tui

import (
    "strings"

    "github.com/rivo/tview"
    "github.com/Murtaza-Udaipurwala/trt/core"
)

type Navigation struct {
    widget *tview.Table
}

func (nav *Navigation) setHeaders() {
    var headers []string = []string { "Overview", "Files", "Trackers", "Peers" }
    for col, header := range headers {
        nav.widget.SetCell(0, col, tview.NewTableCell(header).SetExpansion(1).SetAlign(tview.AlignCenter))
    }
}

func initNavigation(session *core.Session) *Navigation {
    return &Navigation{
        widget: tview.NewTable().SetSelectable(false, true).SetFixed(1, 1).SetSelectionChangedFunc(func(row, column int) {
            switch currentWidget {
            case "overview":
                tui.layout.RemoveItem(tui.overview.widget)
            case "files":
                tui.layout.RemoveItem(tui.files.widget)
            case "trackers":
                tui.layout.RemoveItem(tui.trackers.widget)
            case "peers":
                tui.layout.RemoveItem(tui.peers.widget)
            }

            currentWidget = strings.ToLower(tui.navigation.widget.GetCell(row, column).Text)

            switch currentWidget {
            case "overview":
                tui.layout.AddItem(tui.overview.widget, 0, 1, false)
            case "files":
                tui.layout.AddItem(tui.files.widget, 0, 1, false)
            case "trackers":
                tui.layout.AddItem(tui.trackers.widget, 0, 1, false)
            case "peers":
                tui.layout.AddItem(tui.peers.widget, 0, 1, false)
            }

            redraw(session)
        }),
    }
}
