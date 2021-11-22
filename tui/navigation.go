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
    var headers []string = []string { "Overview", "Files", "Peers", "Trackers" }
    for col, header := range headers {
        nav.widget.SetCell(0, col, tview.NewTableCell(header).SetExpansion(1).SetAlign(tview.AlignCenter))
    }
}

func initNavigation(session *core.Session) *Navigation {
    return &Navigation{
        widget: tview.NewTable().SetSelectable(false, true).SetFixed(1, 1).SetSelectionChangedFunc(func(row, column int) {
            currentWidget = strings.ToLower(tui.navigation.widget.GetCell(row, column).Text)
            tui.layout.Clear()
            tui.layout.AddItem(tui.navigation.widget, 1, 1, true)

            switch currentWidget {
            case "overview":
                tui.layout.AddItem(tui.overview.widget, 0, 1, true)
            case "files":
                tui.layout.AddItem(tui.files.widget, 0, 1, true)
            case "peers":
                tui.layout.AddItem(tui.peers.widget, 0, 1, true)
            case "trackers":
                tui.layout.AddItem(tui.trackers.widget, 0, 1, true)
            }

            redraw(session)
        }),
    }
}
