package tui

import (
	"github.com/Murtaza-Udaipurwala/trt/core"
	"github.com/gdamore/tcell/v2"
)

func setKeys(session *core.Session) {
    tui.torrents.setKeys(session)
    tui.navigation.widget.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'q':
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

            case "peers":
                tui.app.SetFocus(tui.peers.widget)
                tui.navigation.widget.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorBlack))
                tui.peers.widget.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack))
                return nil

            case "files":
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


    tui.peers.widget.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'k':
            row, _ := tui.peers.widget.GetSelection()
            if row == 1 {
                tui.app.SetFocus(tui.layout)
                tui.peers.widget.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorBlack))
                tui.navigation.widget.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack))
                return nil
            }

        case 'q':
            tui.app.SetFocus(tui.layout)
            tui.peers.widget.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorBlack))
            tui.navigation.widget.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorWhite).  Foreground(tcell.ColorBlack))
            return nil
        }
        return event
    })
}
