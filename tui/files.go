package tui

import (
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
    var headers []string = []string { "Name", "Total Size", "Downloaded" }
    for col, header := range headers {
        f.widget.SetCell(0, col, tview.NewTableCell(header).
                                       SetSelectable(false).
                                       SetExpansion(1).
                                       SetTextColor(tcell.ColorYellow))
    }
}

var filesFields []string = []string { "id", "files" }

func (f *Files) update(session *core.Session) {
    torrent, err := core.GetTorrentByID(session, tui.id, filesFields)
    if err != nil {
        currentWidget = "torrents"
        redraw(session)
        tui.pages.RemovePage("details")
    }

    files := torrent.Files
    for row, file := range files {
        size := parseBytes(float64(file.Length))
        downloaded := parseBytes(float64(file.BytesCompleted))
        name := file.Name

        f.widget.SetCell(row + 1, 0, tview.NewTableCell(name))
        f.widget.SetCell(row + 1, 1, tview.NewTableCell(size))
        f.widget.SetCell(row + 1, 2, tview.NewTableCell(downloaded))
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
        }
        return event
    })
}
