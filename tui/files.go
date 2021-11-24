package tui

import (
    "strings"

    "github.com/Murtaza-Udaipurwala/trt/core"
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type Files struct {
    widget *tview.Table
    num int
    torrentID int
}

func initFiles() *Files {
    return &Files{
        widget: tview.NewTable().
                    SetSelectable(true, false).SetFixed(1, 1).
                    SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorBlack)).
                    SetSelectionChangedFunc(func(row, column int) {
                        tui.files.num = row - 1
                    }),
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
    if len(files) == 0 {
        return
    }

    f.setHeaders()
    f.torrentID = torrent.ID
    priorities := torrent.Priorities
    wanted := torrent.Wanted

    for row, file := range files {
        size := parseBytes(float64(file.Length))
        downloaded := parseBytes(float64(file.BytesCompleted))

        splits := strings.Split(file.Name, "/")
        if splits[0] == torrent.Name && len(files) != 1 {
            splits = append(splits[:0], splits[0 + 1:]...)
        }
        name := strings.Join(splits, "/")

        var priority string
        var textColor tcell.Color
        switch priorities[row] {
        case -1:
            priority = "low"
            textColor = tcell.ColorGreen
        case 0:
            priority = "normal"
            textColor = tcell.ColorWhite
        case 1:
            priority = "high"
            textColor = tcell.ColorRed
        }

        if wanted[row] == 0 {
            priority = "off"
            textColor = tcell.ColorBlue
        }

        f.widget.SetCell(row + 1, 0, tview.NewTableCell(size))
        f.widget.SetCell(row + 1, 1, tview.NewTableCell(downloaded))
        f.widget.SetCell(row + 1, 2, tview.NewTableCell(priority).SetTextColor(textColor))
        f.widget.SetCell(row + 1, 3, tview.NewTableCell(name))
    }
}

func (f *Files) setKeys(session *core.Session) {
    tui.files.widget.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'k':
            row, _ := tui.files.widget.GetSelection()
            if row == 1 {
                tui.app.SetFocus(tui.layout)
                setSelectedCellStyle(tui.files.widget,
                                     tcell.StyleDefault.Background(tcell.ColorBlack))

                setSelectedCellStyle(tui.navigation.widget,
                                     tcell.StyleDefault.
                                            Background(tcell.ColorWhite).
                                            Foreground(tcell.ColorBlack))
                return nil
            }

        case 'q':
            tui.app.SetFocus(tui.layout)
            setSelectedCellStyle(tui.files.widget,
                                 tcell.StyleDefault.
                                       Background(tcell.ColorBlack))

            setSelectedCellStyle(tui.navigation.widget,
                                 tcell.StyleDefault.
                                        Background(tcell.ColorWhite).
                                        Foreground(tcell.ColorBlack))
            return nil

        case 'g':
            f.widget.Select(1, 0)
            return nil

        case 'G':
            f.widget.Select(f.widget.GetRowCount() - 1, 0)
            return nil

        case 'h':
            currentPriority := f.widget.GetCell(f.num + 1, 2).Text
            switch currentPriority {
            case "low":
                core.ChangeFilePriority(f.num, f.torrentID, "low", false, session)
                f.update(session)
            case "normal":
                core.ChangeFilePriority(f.num, f.torrentID, "low", true, session)
                f.update(session)
            case "high":
                core.ChangeFilePriority(f.num, f.torrentID, "normal", true, session)
                f.update(session)
            }
            return nil

        case 'l':
            currentPriority := f.widget.GetCell(f.num + 1, 2).Text
            switch currentPriority {
            case "off":
                core.ChangeFilePriority(f.num, f.torrentID, "low", true, session)
                f.update(session)
            case "low":
                core.ChangeFilePriority(f.num, f.torrentID, "normal", true, session)
                f.update(session)
            case "normal":
                core.ChangeFilePriority(f.num, f.torrentID, "high", true, session)
                f.update(session)
            }
            return nil
        }
        return event
    })
}
