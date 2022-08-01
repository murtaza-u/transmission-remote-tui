package tui

import (
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/murtaza-u/trt/core"
	"github.com/rivo/tview"
	"golang.org/x/term"
)

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
	force    chan struct{}
	width    int
}

var tui *TUI

func InitTUI(s *core.Session) (*TUI, error) {
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
		force:    make(chan struct{}, 1000),
	}

	err := tui.setWidth()
	if err != nil {
		return nil, err
	}

	tui.pages.AddPage(TorrentPage, tui.torrent.widget, true, true)
	tui.pages.AddPage(DetailsPage, tui.layout.widget, true, false)

	tui.layout.widget.AddItem(tui.nav.widget, 1, 1, true)
	tui.layout.widget.AddItem(tui.overview.widget, 0, 1, false)

	tui.nav.curr = tui.overview.widget

	tui.files.style()
	tui.peers.style()

	tui.setKeys(s)
	return tui, nil
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

func (t *TUI) drainForce() {
	select {
	case <-t.force:
		t.drainForce()
	default:
		return
	}
}

func (t *TUI) setWidth() error {
	// on non-Unix systems fd may not be 0
	fd := int(os.Stdin.Fd())
	w, _, err := term.GetSize(fd)
	if err != nil {
		return err
	}

	t.width = w
	return nil
}

func (t *TUI) Run(s *core.Session) error {
	go func() {
		for {
			t.drainForce()
			err := t.redraw(s)
			if err != nil {
				log.Fatal(err)
			}

			select {
			case <-tui.force:
			case <-time.After(time.Second):
			}
		}
	}()

	err := t.app.SetRoot(t.pages, true).SetFocus(t.pages).Run()
	return err
}
