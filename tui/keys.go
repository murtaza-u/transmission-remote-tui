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
                return nil

            case "trackers":
                row, col := tui.trackers.widget.GetScrollOffset()
                tui.trackers.widget.ScrollTo(row + 1, col)
                return nil
            }

        case 'k':
            switch currentWidget {
            case "overview":
                row, col := tui.overview.widget.GetScrollOffset()
                tui.overview.widget.ScrollTo(row - 1, col)
                return nil

            case "trackers":
                row, col := tui.trackers.widget.GetScrollOffset()
                tui.trackers.widget.ScrollTo(row - 1, col)
                return nil
            }

        case 'h':
            _, col := tui.navigation.widget.GetSelection()
            if col == 0 {
                return nil
            }
            tui.navigation.widget.Select(0, col - 1)
            return nil

        case 'l':
            _, col := tui.navigation.widget.GetSelection()
            if col == tui.navigation.widget.GetColumnCount() - 1 {
                return nil
            }
            tui.navigation.widget.Select(0, col + 1)
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

    // case "peers":
    //     row, col := tui.peers.widget.GetSelection()
    //     if row == tui.peers.widget.GetRowCount() - 1 {
    //         return nil
    //     }
    //     tui.peers.widget.Select(row + 1, col)

    // case "peers":
    //     row, col := tui.peers.widget.GetSelection()
    //     if row == 1 {
    //         return nil
    //     }
    //     tui.peers.widget.Select(row - 1, col)

    // case 'g':
    //     if currentWidget != "peers" {
    //         return nil
    //     }
    //     tui.peers.widget.Select(1, 0)
    //     return nil

    // case 'G':
    //     if currentWidget != "peers" {
    //         return nil
    //     }
    //     tui.peers.widget.Select(tui.peers.widget.GetRowCount() - 1, 0)
    //     return nil
}
