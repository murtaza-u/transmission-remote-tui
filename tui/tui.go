package tui

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/murtaza-u/trt/core"
	"github.com/rivo/tview"
)

var force = make(chan struct{})

const (
	TorrentPage = "torrent"
	DetailsPage = "details"
)

type TUI struct {
	app      *tview.Application
	pages    *tview.Pages
	torrent  *torrentWid
	overview *overviewWid
	nav      *navWid
	layout   *layoutWid
	files    *files
	peers    *peers
	trackers *trackers
}

var tui *TUI

func InitTUI(s *core.Session) *TUI {
	tui = &TUI{
		app:      tview.NewApplication(),
		pages:    tview.NewPages(),
		torrent:  initTorrentWid(s),
		overview: initOverviewWid(s),
		nav:      initNav(s),
		layout:   initLayout(),
		files:    initFiles(s),
		peers:    initPeers(),
		trackers: initTrackers(),
	}

	tui.setKeys(s)
	return tui
}

func (t *TUI) setKeys(s *core.Session) {
	tui.app.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		switch e.Rune() {
		case 'Q':
			err := s.Close()
			if err != nil {
				log.Fatal(err)
			}
			tui.app.Stop()
			return nil
		}

		return e
	})
}

func (t *TUI) redraw(s *core.Session) error {
	p, _ := t.pages.GetFrontPage()
	var err error

	switch p {
	case TorrentPage:
		err = t.torrent.redraw(s)
	case DetailsPage:
		err = t.layout.redraw(s)
	}

	if err != nil {
		return err
	}

	t.app.Draw()
	return nil
}

func (t *TUI) Run(s *core.Session) error {
	t.pages.AddPage(TorrentPage, t.torrent.widget, true, true)
	t.pages.AddPage(DetailsPage, t.layout.widget, true, false)

	t.layout.widget.AddItem(t.nav.widget, 1, 1, true)
	t.layout.widget.AddItem(t.overview.widget, 0, 1, false)

	t.nav.curr = tui.overview.widget

	t.files.style()
	t.peers.style()

	go func() {
		for {
			err := t.redraw(s)
			if err != nil {
				log.Fatal(err)
			}

			select {
			case <-force:
			case <-time.After(time.Second):
			}
		}
	}()

	err := t.app.SetRoot(t.pages, true).SetFocus(t.pages).Run()
	return err
}
