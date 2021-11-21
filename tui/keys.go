package tui

import (
	"github.com/Murtaza-Udaipurwala/trt/core"
	"github.com/gdamore/tcell/v2"
)

func setKeys(session *core.Session) {
    tui.torrentList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'q':
            tui.app.Stop()
            return nil

        case 'p':
            core.PauseStartTorrent(tui.torrentList, name, session)
            updateTable(session)
            return nil

        case 'r':
            core.RemoveTorrent(tui.torrentList, name, session, false)
            return nil

        case 'R':
            core.RemoveTorrent(tui.torrentList, name, session, true)
            return nil

        case 'v':
            core.VerifyTorrent(tui.torrentList, name, session)
            updateTable(session)
            return nil

        case 'g':
            tui.torrentList.Select(1, 0)
            return nil

        case 'G':
            rows := tui.torrentList.GetRowCount()
            tui.torrentList.Select(rows - 1, 0)
            return nil

        case 'K':
            newRow := core.QueueMove("up", tui.torrentList, name, session)
            updateTable(session)
            tui.torrentList.Select(newRow, 0)
            return nil

        case 'J':
            newRow := core.QueueMove("down", tui.torrentList, name, session)
            updateTable(session)
            tui.torrentList.Select(newRow, 0)
            return nil

        case 'U':
            newRow := core.QueueMove("top", tui.torrentList, name, session)
            updateTable(session)
            tui.torrentList.Select(newRow, 0)
            return nil

        case 'D':
            newRow := core.QueueMove("bottom", tui.torrentList, name, session)
            updateTable(session)
            tui.torrentList.Select(newRow, 0)
            return nil

        case 'l', rune(tcell.KeyEnter):
            tui.pages.AddAndSwitchToPage("torrentDetails", tui.torrentDetails, true)
            return nil
        }

        return event
    })

    tui.torrentDetails.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'q':
            tui.pages.RemovePage("torrentDetails")
            return nil
        }

        return event
    })
}
