package tui

import (
    "github.com/Murtaza-Udaipurwala/trt/core"
    "github.com/gdamore/tcell/v2"
)

func setKeys(session *core.Session) {
    tui.torrents.setKeys(session)
    tui.layout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'q', rune(tcell.KeyESC):
            currentWidget = "torrents"
            tui.pages.RemovePage("details")
            return nil

        case 'j':
            switch currentWidget {
            case "overview":
                row, col := tui.overview.widget.GetScrollOffset()
                tui.overview.widget.ScrollTo(row + 1, col)

            case "trackers":
                row, col := tui.trackers.widget.GetScrollOffset()
                tui.trackers.widget.ScrollTo(row + 1, col)

            case "peers":
                row, col := tui.peers.widget.GetScrollOffset()
                tui.peers.widget.ScrollTo(row + 1, col)
            }
            return nil

        case 'k':
            switch currentWidget {
            case "overview":
                row, col := tui.overview.widget.GetScrollOffset()
                tui.overview.widget.ScrollTo(row - 1, col)

            case "trackers":
                row, col := tui.trackers.widget.GetScrollOffset()
                tui.trackers.widget.ScrollTo(row - 1, col)

            case "peers":
                row, col := tui.peers.widget.GetScrollOffset()
                tui.peers.widget.ScrollTo(row - 1, col)
            }
            return nil
        }

        return event
    })

    tui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'Q':
            core.SendRequest("session-close", "1", core.Arguments{}, session)
            tui.app.Stop()
            return nil
        }

        return event
    })
}
