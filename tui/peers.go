package tui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/murtaza-u/trt/core"
	"github.com/rivo/tview"
)

type peers struct {
	widget *tview.Table
	fields []string
}

func initPeers() *peers {
	p := new(peers)
	p.widget = tview.NewTable().SetSelectable(true, false).SetFixed(1, 1)

	p.fields = []string{"peers", "id"}

	p.setHeaders()
	p.setKeys()
	return p
}

func (p *peers) style() {
	p.widget.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorWhite))
}

func (p *peers) setHeaders() {
	var heads = []string{"Flag", "Progress", "Client", "Address", "Port"}

	for col, h := range heads {
		p.widget.SetCell(
			0, col, tview.NewTableCell(h).SetSelectable(false).SetExpansion(1),
		)
	}
}

func (p *peers) redraw(s *core.Session) error {
	id := tui.torrent.currID()
	t, err := s.GetTorrentByID(id, p.fields)
	if err != nil {
		return err
	}

	for i, peer := range t.Peers {
		attrs := make([]string, 0, 5)
		attrs = append(attrs, peer.FlagStr)
		attrs = append(attrs, fmt.Sprintf("%.2f%%", peer.Progress*100))
		attrs = append(attrs, peer.ClientName)
		attrs = append(attrs, peer.Address)
		attrs = append(attrs, fmt.Sprint(peer.Port))

		for col := 0; col <= 4; col++ {
			p.widget.SetCell(
				i+1, col, tview.NewTableCell(attrs[col]).SetExpansion(1),
			)
		}
	}

	for i := len(t.Peers) + 1; i < p.widget.GetRowCount(); i++ {
		p.widget.RemoveRow(i)
	}

	return nil
}

func (p *peers) setKeys() {
	p.widget.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		switch e.Rune() {
		case 'q':
			tui.layout.focus(p.widget)
			return nil

		case 'k':
			r, _ := p.widget.GetSelection()
			if r == 1 {
				tui.layout.focus(p.widget)
				return nil
			}

		case 'g':
			p.widget.Select(1, 0)
			p.widget.ScrollToBeginning()
			return nil

		case 'G':
			p.widget.Select(p.widget.GetRowCount()-1, 0)
			p.widget.ScrollToEnd()
			return nil
		}

		return e
	})
}
