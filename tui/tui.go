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
    torrentList *tview.Table
    torrentDetails *tview.Flex
    navigation *tview.Table
    overview *tview.TextView
    files *tview.Table
    trackers *tview.TextView
}

var tui *TUI

func initTUI() *TUI {
    return &TUI{
        app: tview.NewApplication(),
        pages: tview.NewPages(),
        torrentList: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1),
        torrentDetails: tview.NewFlex().SetDirection(tview.FlexRow),
        navigation: tview.NewTable().SetSelectable(false, true).SetFixed(1, 1),
    }
}

func setNavigation(navigation *tview.Table) {
    var headers []string = []string { "Overview", "Files", "Trackers" }
    for col, header := range headers {
        navigation.SetCell(0, col, tview.NewTableCell(header).SetExpansion(1))
    }
}

func setTorrentDetails(tui *TUI) {
    setNavigation(tui.navigation)
    tui.torrentDetails.AddItem(tui.navigation, 1, 1, true)
}

func Run(session *core.Session) {
    tui = initTUI()
    tui.pages.AddPage("torrentList", tui.torrentList, true, true)

    setTorrentDetails(tui)
    setTableHeaders(tui.torrentList)

    go func() {
        for {
            updateTable(session)
            tui.app.Draw()
            time.Sleep(time.Second)
        }
    }()

    setKeys(session)

    if err := tui.app.SetRoot(tui.pages, true).SetFocus(tui.pages).Run(); err != nil {
        log.Panic(err)
    }
}
