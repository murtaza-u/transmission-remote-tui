package tui

import (
	"errors"
	"fmt"

	"github.com/Murtaza-Udaipurwala/trt/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Torrents struct {
    widget *tview.Table
    torrents []interface{}
}

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

func initTorrents() *Torrents {
    return &Torrents{
        widget: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1),
    }
}

func (torrents *Torrents) update(session *core.Session) {
    torrents.torrents = core.SortTorrentsByQueuePosition(core.GetTorrents(session))

    for row, torrent := range torrents.torrents {
        t := torrent.(map[string]interface{})

        statusCode  := int((t["status"].(float64)))
        seeders, leechers := core.GetSeedersLeechers(t["trackerStats"].([]interface{}))

        tui.torrents.widget.SetCell(row + 1, 0, tview.NewTableCell(core.TorrentStatus[statusCode]))
        tui.torrents.widget.SetCell(row + 1, 1, tview.NewTableCell(fmt.Sprintf("%s", convertSecondsTo(t["eta"].(float64)))))
        tui.torrents.widget.SetCell(row + 1, 2, tview.NewTableCell(fmt.Sprintf("%s/s", convertBytesTo(t["rateUpload"].(float64)))))
        tui.torrents.widget.SetCell(row + 1, 3, tview.NewTableCell(fmt.Sprintf("%s/s", convertBytesTo(t["rateDownload"].(float64)))))
        tui.torrents.widget.SetCell(row + 1, 4, tview.NewTableCell(fmt.Sprintf("%v", t["uploadRatio"])))
        tui.torrents.widget.SetCell(row + 1, 5, tview.NewTableCell(fmt.Sprintf("%v", t["peersConnected"])))
        tui.torrents.widget.SetCell(row + 1, 6, tview.NewTableCell(seeders))
        tui.torrents.widget.SetCell(row + 1, 7, tview.NewTableCell(leechers))
        tui.torrents.widget.SetCell(row + 1, 8, tview.NewTableCell(fmt.Sprintf("%s", convertBytesTo(t["totalSize"].(float64)))))
        tui.torrents.widget.SetCell(row + 1, 9, tview.NewTableCell(fmt.Sprintf("%s", convertBytesTo(t["leftUntilDone"].(float64)))))
        tui.torrents.widget.SetCell(row + 1, 10, tview.NewTableCell(fmt.Sprintf("%v", t["name"])))
    }
}

func (torrents *Torrents) setHeaders() {
    var headers []string = []string {
        "Status", "ETA", "Upload Rate", "Download Rate", "Ratio", "Peers",
        "Seeders", "Leechers", "Size", "Left", "Name",
    }

    for col, header := range headers {
        torrents.widget.
            SetCell(0, col, tview.NewTableCell(header).
            SetSelectable(false).
            SetTextColor(tcell.ColorYellow).
            SetExpansion(1))
    }
}

func (torrents *Torrents) currentSelected() (map[string]interface{}, error) {
    row, _ := torrents.widget.GetSelection()
    name := torrents.widget.GetCell(row, name).Text
    for _, torrent := range torrents.torrents {
        t := torrent.(map[string]interface{})
        if t["name"].(string) == name {
            return t, nil
        }
    }

    return nil, errors.New("Torrent not found")
}

func (torrents *Torrents) currentSelectedID() (int, error) {
    torrent, err := torrents.currentSelected()
    if err != nil {
        return -1, err
    }

    return int(torrent["id"].(float64)), nil
}

func (torrents *Torrents) setKeys(session *core.Session) {
    torrents.widget.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'q':
            tui.app.Stop()
            return nil

        case 'p':
            id, err := torrents.currentSelectedID()
            if err != nil {
                return nil
            }
            core.PauseStartTorrent(id, session,  torrents.torrents)
            torrents.update(session)
            return nil

        case 'r':
            id, err := torrents.currentSelectedID()
            if err != nil {
                return nil
            }
            row, _ := torrents.widget.GetSelection()
            torrents.widget.RemoveRow(row)
            core.RemoveTorrent(id, session, false)
            return nil

        case 'R':
            id, err := torrents.currentSelectedID()
            if err != nil {
                return nil
            }
            row, _ := torrents.widget.GetSelection()
            torrents.widget.RemoveRow(row)
            core.RemoveTorrent(id, session, true)
            return nil

        case 'v':
            id, err := torrents.currentSelectedID()
            if err != nil {
                return nil
            }
            core.VerifyTorrent(id, session)
            torrents.update(session)
            return nil

        case 'g':
            tui.torrents.widget.Select(1, 0)
            return nil

        case 'G':
            tui.torrents.widget.Select(torrents.widget.GetRowCount() - 1, 0)
            return nil

        case 'K':
            id, err := torrents.currentSelectedID()
            if err != nil {
                return nil
            }
            core.QueueMove("up", id, session)
            torrents.update(session)
            row, _ := torrents.widget.GetSelection()
            torrents.widget.Select(row - 1, 0)
            return nil

        case 'J':
            id, err := torrents.currentSelectedID()
            if err != nil {
                return nil
            }
            core.QueueMove("down", id, session)
            torrents.update(session)
            row, _ := torrents.widget.GetSelection()
            torrents.widget.Select(row + 1, 0)
            return nil

        case 'U':
            id, err := torrents.currentSelectedID()
            if err != nil {
                return nil
            }
            core.QueueMove("top", id, session)
            torrents.update(session)
            torrents.widget.Select(1, 0)
            return nil

        case 'D':
            id, err := torrents.currentSelectedID()
            if err != nil {
                return nil
            }
            core.QueueMove("bottom", id, session)
            torrents.update(session)
            torrents.widget.Select(torrents.widget.GetRowCount() - 1, 0)
            return nil

        case 'l', rune(tcell.KeyEnter):
            tui.pages.AddAndSwitchToPage("torrentDetails", tui.layout, true)
            return nil
        }

        return event
    })
}
