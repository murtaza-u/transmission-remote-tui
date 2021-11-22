package tui

import (
    "log"
    "time"

    "github.com/Murtaza-Udaipurwala/trt/core"
    "github.com/rivo/tview"
)

type TUI struct {
    app *tview.Application
    pages *tview.Pages
    torrents *Torrents
    layout *tview.Flex
    navigation *Navigation
    overview *tview.TextView
    files *tview.Table
    trackers *tview.TextView
    peers *tview.TextView
}

var tui *TUI

func initTUI() *TUI {
    return &TUI{
        app: tview.NewApplication(),
        pages: tview.NewPages(),
        torrents: initTorrents(),
        layout: tview.NewFlex().SetDirection(tview.FlexRow),
        navigation: initNavigation(),
    }
}

func Run(session *core.Session) {
    tui = initTUI()
    tui.pages.AddPage("torrents", tui.torrents.widget, true, true)

    tui.navigation.setHeaders()
    tui.layout.AddItem(tui.navigation.widget, 1, 1, true)

    tui.torrents.setHeaders()

    go func() {
        for {
            tui.torrents.update(session)
            tui.app.Draw()
            time.Sleep(time.Second)
        }
    }()

    setKeys(session)

    if err := tui.app.SetRoot(tui.pages, true).SetFocus(tui.pages).Run(); err != nil {
        log.Panic(err)
    }
}
