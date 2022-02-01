package tui

import (
	"log"
	"time"

	"github.com/Murtaza-Udaipurwala/trt/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	app        *tview.Application
	pages      *tview.Pages
	layout     *tview.Flex
	torrents   *List
	navigation *Navigation
	overview   *Overview
	files      *Files
	trackers   *Trackers
	peers      *Peers
	id         int
}

var tui *TUI

func initTUI(session *core.Session) *TUI {
	return &TUI{
		app:        tview.NewApplication(),
		pages:      tview.NewPages(),
		torrents:   initTorrents(),
		layout:     tview.NewFlex().SetDirection(tview.FlexRow),
		navigation: initNavigation(session),
		overview:   initOverview(),
		files:      initFiles(),
		peers:      initPeers(),
		trackers:   initTrackers(),
		id:         -1,
	}
}

var currentWidget string

func redraw(session *core.Session) {
	switch currentWidget {
	case "torrents":
		tui.torrents.update(session)
	case "overview":
		tui.overview.update(session)
	case "files":
		tui.files.update(session)
	case "peers":
		tui.peers.update(session)
	case "trackers":
		tui.trackers.update(session)
	}
}

func setKeys(session *core.Session) {
	tui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'Q':
			core.SendRequest("session-close", "1", core.Arguments{}, session)
			tui.app.Stop()
			return nil
		}

		return event
	})

	tui.torrents.setKeys(session)
	tui.navigation.setKeys()
	tui.peers.setKeys()
	tui.files.setKeys(session)
}

func Run(session *core.Session) {
	currentWidget = "torrents"
	tui = initTUI(session)
	tui.pages.AddPage("torrents", tui.torrents.widget, true, true)

	tui.navigation.setHeaders()

	tui.layout.AddItem(tui.navigation.widget, 1, 1, true)
	tui.layout.AddItem(tui.overview.widget, 0, 1, false)

	tui.torrents.setHeaders()

	go func() {
		for {
			redraw(session)
			tui.app.Draw()
			time.Sleep(time.Second)
		}
	}()

	setKeys(session)

	if err := tui.app.SetRoot(tui.pages, true).SetFocus(tui.pages).Run(); err != nil {
		log.Panic(err)
	}
}
