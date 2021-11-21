package tui

import (
    "fmt"
    "log"
    "time"

    "github.com/rivo/tview"
    "github.com/Murtaza-Udaipurwala/trt/core"
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

func initTUI() *TUI {
    return &TUI{
        app: tview.NewApplication(),
        table: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1),
    }
}

var tui *TUI

func update(session *core.Session) {
    core.GetTorrents(session)
    core.SortTorrentsByQueuePosition()

    for row, torrent := range core.Torrents {
        t := torrent.(map[string]interface{})

        statusCode  := int((t["status"].(float64)))

        tui.table.SetCell(row + 1, 0, tview.NewTableCell(core.TorrentStatus[statusCode]))
        tui.table.SetCell(row + 1, 1, tview.NewTableCell(fmt.Sprintf("%s", convertSecondsTo(t["eta"].(float64)))))
        tui.table.SetCell(row + 1, 2, tview.NewTableCell(fmt.Sprintf("%s/s🔼", convertBytesTo(t["rateUpload"].(float64)))))
        tui.table.SetCell(row + 1, 3, tview.NewTableCell(fmt.Sprintf("%s/s🔽", convertBytesTo(t["rateDownload"].(float64)))))
        tui.table.SetCell(row + 1, 4, tview.NewTableCell(fmt.Sprintf("%v", t["uploadRatio"])))
        tui.table.SetCell(row + 1, 5, tview.NewTableCell(fmt.Sprintf("%v", t["peersConnected"])))
        tui.table.SetCell(row + 1, 6, tview.NewTableCell(fmt.Sprintf("%s", convertBytesTo(t["totalSize"].(float64)))))
        tui.table.SetCell(row + 1, 7, tview.NewTableCell(fmt.Sprintf("%s", convertBytesTo(t["leftUntilDone"].(float64)))))
        tui.table.SetCell(row + 1, 8, tview.NewTableCell(fmt.Sprintf("%v", t["name"])))
    }
}

func Run(session *core.Session) {
    tui = initTUI()

    tui.table.SetCell(0, status, tview.NewTableCell("Status").SetSelectable(false).SetExpansion(1))
    tui.table.SetCell(0, eta, tview.NewTableCell("ETA").SetSelectable(false).SetExpansion(1))
    tui.table.SetCell(0, uploadRate, tview.NewTableCell("Upload Rate").SetSelectable(false).SetExpansion(1))
    tui.table.SetCell(0, downloadRate, tview.NewTableCell("Download Rate").SetSelectable(false).SetExpansion(1))
    tui.table.SetCell(0, ratio, tview.NewTableCell("Ratio").SetSelectable(false).SetExpansion(1))
    tui.table.SetCell(0, peers, tview.NewTableCell("Peers").SetSelectable(false).SetExpansion(1))
    tui.table.SetCell(0, size, tview.NewTableCell("Size").SetSelectable(false).SetExpansion(1))
    tui.table.SetCell(0, left, tview.NewTableCell("Left").SetSelectable(false).SetExpansion(1))
    tui.table.SetCell(0, name, tview.NewTableCell("Name").SetSelectable(false).SetExpansion(2))

    go func() {
        for {
            update(session)
            tui.app.Draw()
            time.Sleep(time.Second)
        }
    }()

    setKeys(session)

    if err := tui.app.SetRoot(tui.table, true).SetFocus(tui.table).Run(); err != nil {
        log.Panic(err)
    }
}
