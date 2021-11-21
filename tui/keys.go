package tui

import (
	"github.com/Murtaza-Udaipurwala/trt/core"
	"github.com/gdamore/tcell/v2"
)

func setKeys(session *core.Session) {
    tui.table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'q':
            tui.app.Stop()
            return nil

        case 'p':
            core.PauseStartTorrent(tui.table, name, session)
            update(session)
            return nil

        case 'r':
            core.RemoveTorrent(tui.table, name, session, false)
            return nil

        case 'R':
            core.RemoveTorrent(tui.table, name, session, true)
            return nil

        case 'v':
            core.VerifyTorrent(tui.table, name, session)
            update(session)
            return nil

        case 'g':
            tui.table.Select(1, 0)
            return nil

        case 'G':
            rows := tui.table.GetRowCount()
            tui.table.Select(rows - 1, 0)
            return nil

        case 'K':
            newRow := core.QueueMove("up", tui.table, name, session)
            update(session)
            tui.table.Select(newRow, 0)
            return nil

        case 'J':
            newRow := core.QueueMove("down", tui.table, name, session)
            update(session)
            tui.table.Select(newRow, 0)
            return nil
        }

        return event
    })
}
