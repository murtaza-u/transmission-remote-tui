package tui

import (
    "github.com/Murtaza-Udaipurwala/trt/core"
    "github.com/gdamore/tcell/v2"
)

func setKeys(session *core.Session) {
    tui.torrents.setKeys(session)
    tui.layout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'q':
            currentWidget = "torrents"
            tui.pages.RemovePage("details")
            return nil
        }

        return event
    })
}
