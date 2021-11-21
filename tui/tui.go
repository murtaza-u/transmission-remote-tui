package tui

import (
    "fmt"
    "log"
    "time"

    "github.com/Murtaza-Udaipurwala/trt/core"
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

const (
    status = iota
    eta
    uploadRate
    downloadRate
    ratio
    peers
    size
    left
    name
)

type TUI struct {
    app *tview.Application
    table *tview.Table
}

var tui *TUI

func initTUI() *TUI {
    return &TUI{
        app: tview.NewApplication(),
        table: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1),
    }
}

func updateTable(session *core.Session) {
    core.GetTorrents(session)
    core.SortTorrentsByQueuePosition()

    for row, torrent := range core.Torrents {
        t := torrent.(map[string]interface{})

        statusCode  := int((t["status"].(float64)))
        seeders, leechers := core.GetSeedersLeechers(t["trackerStats"].([]interface{}))

        tui.table.SetCell(row + 1, 0, tview.NewTableCell(core.TorrentStatus[statusCode]))
        tui.table.SetCell(row + 1, 1, tview.NewTableCell(fmt.Sprintf("%s", convertSecondsTo(t["eta"].(float64)))))
        tui.table.SetCell(row + 1, 2, tview.NewTableCell(fmt.Sprintf("%s/s", convertBytesTo(t["rateUpload"].(float64)))))
        tui.table.SetCell(row + 1, 3, tview.NewTableCell(fmt.Sprintf("%s/s", convertBytesTo(t["rateDownload"].(float64)))))
        tui.table.SetCell(row + 1, 4, tview.NewTableCell(fmt.Sprintf("%v", t["uploadRatio"])))
        tui.table.SetCell(row + 1, 5, tview.NewTableCell(fmt.Sprintf("%v", t["peersConnected"])))
        tui.table.SetCell(row + 1, 6, tview.NewTableCell(seeders))
        tui.table.SetCell(row + 1, 7, tview.NewTableCell(leechers))
        tui.table.SetCell(row + 1, 8, tview.NewTableCell(fmt.Sprintf("%s", convertBytesTo(t["totalSize"].(float64)))))
        tui.table.SetCell(row + 1, 9, tview.NewTableCell(fmt.Sprintf("%s", convertBytesTo(t["leftUntilDone"].(float64)))))
        tui.table.SetCell(row + 1, 10, tview.NewTableCell(fmt.Sprintf("%v", t["name"])))
    }
}

func SetTableHeaders(table *tview.Table) {
    var headers []string = []string {
        "Status", "ETA", "Upload Rate", "Download Rate", "Ratio", "Peers",
        "Seeders", "Leechers", "Size", "Left", "Name",
    }

    for col, header := range headers {
        table.
            SetCell(0, col, tview.NewTableCell(header).
            SetSelectable(false).
            SetTextColor(tcell.ColorYellow).
            SetExpansion(1))
    }
}

func Run(session *core.Session) {
    tui = initTUI()
    SetTableHeaders(tui.table)

    tui.table.SetFocusFunc(func() {
    })

    go func() {
        for {
            updateTable(session)
            tui.app.Draw()
            time.Sleep(time.Second)
        }
    }()

    setKeys(session)

    if err := tui.app.SetRoot(tui.table, true).SetFocus(tui.table).Run(); err != nil {
        log.Panic(err)
    }
}
