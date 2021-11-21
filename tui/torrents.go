package tui

import (
    "fmt"

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
    seeders
    leechers
    size
    left
    name
)

func updateTable(session *core.Session) {
    core.GetTorrents(session)
    core.SortTorrentsByQueuePosition()

    for row, torrent := range core.Torrents {
        t := torrent.(map[string]interface{})

        statusCode  := int((t["status"].(float64)))
        seeders, leechers := core.GetSeedersLeechers(t["trackerStats"].([]interface{}))

        tui.torrentList.SetCell(row + 1, 0, tview.NewTableCell(core.TorrentStatus[statusCode]))
        tui.torrentList.SetCell(row + 1, 1, tview.NewTableCell(fmt.Sprintf("%s", convertSecondsTo(t["eta"].(float64)))))
        tui.torrentList.SetCell(row + 1, 2, tview.NewTableCell(fmt.Sprintf("%s/s", convertBytesTo(t["rateUpload"].(float64)))))
        tui.torrentList.SetCell(row + 1, 3, tview.NewTableCell(fmt.Sprintf("%s/s", convertBytesTo(t["rateDownload"].(float64)))))
        tui.torrentList.SetCell(row + 1, 4, tview.NewTableCell(fmt.Sprintf("%v", t["uploadRatio"])))
        tui.torrentList.SetCell(row + 1, 5, tview.NewTableCell(fmt.Sprintf("%v", t["peersConnected"])))
        tui.torrentList.SetCell(row + 1, 6, tview.NewTableCell(seeders))
        tui.torrentList.SetCell(row + 1, 7, tview.NewTableCell(leechers))
        tui.torrentList.SetCell(row + 1, 8, tview.NewTableCell(fmt.Sprintf("%s", convertBytesTo(t["totalSize"].(float64)))))
        tui.torrentList.SetCell(row + 1, 9, tview.NewTableCell(fmt.Sprintf("%s", convertBytesTo(t["leftUntilDone"].(float64)))))
        tui.torrentList.SetCell(row + 1, 10, tview.NewTableCell(fmt.Sprintf("%v", t["name"])))
    }
}

func setTableHeaders(table *tview.Table) {
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
