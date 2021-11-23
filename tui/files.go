package tui

import (
	"strings"

	"github.com/Murtaza-Udaipurwala/trt/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Files struct {
    widget *tview.Table
}

func initFiles() *Files {
    return &Files{
        widget: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1).
                                 SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorBlack)),
    }
}

func (f *Files) setHeaders() {
    var headers []string = []string { "Total Size", "Downloaded", "Priority", "Name" }
    for col, header := range headers {
        f.widget.SetCell(0, col, tview.NewTableCell(header).
                                       SetSelectable(false).
                                       SetExpansion(1).
                                       SetTextColor(tcell.ColorYellow))
    }
}

var filesFields []string = []string { "id", "files", "wanted", "priorities", "name" }

func (f *Files) update(session *core.Session) {
    torrent, err := core.GetTorrentByID(session, tui.id, filesFields)
    if err != nil {
        currentWidget = "torrents"
        redraw(session)
        tui.pages.RemovePage("details")
    }

    files := torrent.Files
    priorities := torrent.Priorities
    wanted := torrent.Wanted

    for row, file := range files {
        size := parseBytes(float64(file.Length))
        downloaded := parseBytes(float64(file.BytesCompleted))

        splits := strings.Split(file.Name, "/")
        if splits[0] == torrent.Name {
            splits = append(splits[:0], splits[0 + 1:]...)
        }
        name := strings.Join(splits, "/")

        var priority string
        switch priorities[row] {
        case -1:
            priority = "Low"
        case 0:
            priority = "Normal"
        case 1:
            priority = "High"
        }

        if wanted[row] == 0 {
            priority = "Off"
        }

        f.widget.SetCell(row + 1, 0, tview.NewTableCell(size))
        f.widget.SetCell(row + 1, 1, tview.NewTableCell(downloaded))
        f.widget.SetCell(row + 1, 2, tview.NewTableCell(priority))
        f.widget.SetCell(row + 1, 3, tview.NewTableCell(name))
    }
}

func (f *Files) setKeys() {
    tui.files.widget.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'k':
            row, _ := tui.files.widget.GetSelection()
            if row == 1 {
                tui.app.SetFocus(tui.layout)
                setSelectedCellStyle(tui.files.widget,
                                     tcell.StyleDefault.Background(tcell.ColorBlack))

                setSelectedCellStyle(tui.navigation.widget,
                                     tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack))
                return nil
            }

        case 'q':
            tui.app.SetFocus(tui.layout)
            setSelectedCellStyle(tui.files.widget,
                                 tcell.StyleDefault.Background(tcell.ColorBlack))

            setSelectedCellStyle(tui.navigation.widget,
                                 tcell.StyleDefault.Background(tcell.ColorWhite).  Foreground(tcell.ColorBlack))
            return nil

        case 'g':
            f.widget.Select(1, 0)
            return nil

        case 'G':
            f.widget.Select(f.widget.GetRowCount() - 1, 0)
            return nil
        }
        return event
    })
}
