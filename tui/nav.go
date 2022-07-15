package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/murtaza-u/trt/core"
	"github.com/rivo/tview"
)

type navWid struct {
	widget *tview.Table
	curr   tview.Primitive
}

const (
	NavOverview = iota
	NavFiles
	NavTrackers
	NavPeers
)

func (nav *navWid) setHeaders() {
	var headers = []string{"Overview", "Files", "Trackers", "Peers"}
	for col, h := range headers {
		nav.widget.SetCell(
			0, col,
			tview.NewTableCell(h).SetExpansion(1).SetAlign(tview.AlignCenter),
		)
	}
}

func initNav(s *core.Session) *navWid {
	nav := new(navWid)
	nav.widget = tview.NewTable().
		SetSelectable(false, true).
		SetFixed(1, 1).
		SetSelectionChangedFunc(func(row, col int) {
			tui.layout.widget.RemoveItem(nav.curr)

			switch col {
			case NavOverview:
				nav.curr = tui.overview.widget
			case NavFiles:
				nav.curr = tui.files.widget
			case NavTrackers:
				nav.curr = tui.trackers.widget
			case NavPeers:
				nav.curr = tui.peers.widget
			}

			tui.layout.widget.AddItem(nav.curr, 0, 1, false)
			tui.force <- struct{}{}
		})
	nav.setHeaders()
	nav.setKeys()
	return nav
}

func (nav *navWid) keyj() {
	_, col := nav.widget.GetSelection()

	switch col {
	case NavOverview:
		r, c := tui.overview.widget.GetScrollOffset()
		tui.overview.widget.ScrollTo(r+1, c)

	case NavTrackers:
		r, c := tui.trackers.widget.GetScrollOffset()
		tui.trackers.widget.ScrollTo(r+1, c)

	case NavPeers:
		if tui.peers.widget.GetRowCount() == 1 {
			break
		}
		tui.layout.unfocus(tui.peers.widget)

	case NavFiles:
		if tui.files.widget.GetRowCount() == 1 {
			break
		}
		tui.layout.unfocus(tui.files.widget)
	}
}

func (nav *navWid) keyk() {
	_, col := nav.widget.GetSelection()

	switch col {
	case NavOverview:
		r, c := tui.overview.widget.GetScrollOffset()
		tui.overview.widget.ScrollTo(r-1, c)

	case NavTrackers:
		r, c := tui.trackers.widget.GetScrollOffset()
		tui.trackers.widget.ScrollTo(r-1, c)
	}
}

func (nav *navWid) keyg() {
	_, col := nav.widget.GetSelection()

	switch col {
	case NavOverview:
		tui.overview.widget.ScrollToBeginning()

	case NavTrackers:
		tui.trackers.widget.ScrollToBeginning()
	}
}

func (nav *navWid) keyG() {
	_, col := nav.widget.GetSelection()

	switch col {
	case NavOverview:
		tui.overview.widget.ScrollToEnd()

	case NavTrackers:
		tui.trackers.widget.ScrollToEnd()
	}
}

func (nav *navWid) back() {
	tui.pages.SwitchToPage(TorrentPage)
	tui.force <- struct{}{}
}

func (nav *navWid) setKeys() {
	nav.widget.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		switch e.Rune() {
		case 'q':
			nav.back()
			return nil

		case 'h':
			_, col := nav.widget.GetSelection()
			if col == NavOverview {
				nav.back()
				return nil
			}

		case 'j':
			nav.keyj()
			return nil

		case 'k':
			nav.keyk()
			return nil

		case 'g':
			nav.keyg()
			return nil

		case 'G':
			nav.keyG()
			return nil
		}

		return e
	})
}
