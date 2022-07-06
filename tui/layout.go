package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/murtaza-u/trt/core"
	"github.com/rivo/tview"
)

type layoutWid struct {
	widget *tview.Flex
}

func initLayout() *layoutWid {
	l := new(layoutWid)
	l.widget = tview.NewFlex().SetDirection(tview.FlexRow)
	return l
}

func (l *layoutWid) redraw(s *core.Session) error {
	var err error
	_, c := tui.nav.widget.GetSelection()

	switch c {
	case NavOverview:
		err = tui.overview.redraw(s)
	case NavFiles:
		err = tui.files.redraw(s)
	case NavTrackers:
		err = tui.trackers.redraw(s)
	case NavPeers:
		err = tui.peers.redraw(s)
	}

	return err
}

func (l *layoutWid) focus(t *tview.Table) {
	tui.app.SetFocus(l.widget)
	t.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorBlack))
	tui.nav.widget.SetSelectedStyle(
		tcell.StyleDefault.
			Background(tcell.ColorWhite).
			Foreground(tcell.ColorBlack),
	)
}

func (l *layoutWid) unfocus(t *tview.Table) {
	tui.app.SetFocus(t)

	t.SetSelectedStyle(
		tcell.StyleDefault.
			Background(tcell.ColorWhite).
			Foreground(tcell.ColorBlack),
	)

	tui.nav.widget.SetSelectedStyle(
		tcell.StyleDefault.Background(tcell.ColorBlack),
	)
}
