package tui

import (
	"fmt"

	"github.com/Murtaza-Udaipurwala/trt/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Peers struct {
    widget *tview.Table
}

func initPeers() *Peers {
    return &Peers{
        widget: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1).
                                 SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorBlack)),
    }
}

func (p *Peers) setHeaders() {
    var headers []string = []string {
        "Flag", "Progress", "Client", "Address", "Port",
    }

    for col, header := range headers {
        p.widget.SetCell(0, col, tview.NewTableCell(header).
                                       SetSelectable(false).
                                       SetExpansion(1))
    }
}

var peersFields []string = []string { "peers", "id" }

func (p *Peers) update(session *core.Session) {
    torrent, err := core.GetTorrentByID(session, tui.id,  peersFields)
    if err != nil {
        currentWidget = "torrents"
        redraw(session)
        tui.pages.RemovePage("details")
    }

    if len(torrent.Peers) == 0 {
        p.widget.Clear()
        return
    }

    p.setHeaders()

    for row, peer := range torrent.Peers {
        flag := peer.FlagStr
        address := peer.Address
        client := peer.ClientName
        port := fmt.Sprint(peer.Port)
        progress := fmt.Sprintf("%d%%", peer.Progress * 100)

        p.widget.SetCell(row + 1, 0, tview.NewTableCell(flag).SetExpansion(1))
        p.widget.SetCell(row + 1, 1, tview.NewTableCell(progress).SetExpansion(1))
        p.widget.SetCell(row + 1, 2, tview.NewTableCell(client).SetExpansion(1))
        p.widget.SetCell(row + 1, 3, tview.NewTableCell(address).SetExpansion(1))
        p.widget.SetCell(row + 1, 4, tview.NewTableCell(port).SetExpansion(1))
    }
}

func (p *Peers) setKeys() {
    tui.peers.widget.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'k':
            row, _ := tui.peers.widget.GetSelection()
            if row == 1 {
                tui.app.SetFocus(tui.layout)
                setSelectedCellStyle(tui.peers.widget,
                                     tcell.StyleDefault.Background(tcell.ColorBlack))

                setSelectedCellStyle(tui.navigation.widget,
                                     tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack))
                return nil
            }

        case 'q':
            tui.app.SetFocus(tui.layout)
            setSelectedCellStyle(tui.peers.widget,
                                 tcell.StyleDefault.Background(tcell.ColorBlack))

            setSelectedCellStyle(tui.navigation.widget,
                                 tcell.StyleDefault.Background(tcell.ColorWhite).  Foreground(tcell.ColorBlack))
            return nil
        }
        return event
    })
}
