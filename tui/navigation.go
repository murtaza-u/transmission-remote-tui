package tui

import (
    "strings"

    "github.com/Murtaza-Udaipurwala/trt/core"
    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

type Navigation struct {
    widget *tview.Table
}

func (nav *Navigation) setHeaders() {
    var headers []string = []string { "Overview", "Files", "Trackers", "Peers" }
    for col, header := range headers {
        nav.widget.SetCell(0, col, tview.NewTableCell(header).
                                        SetExpansion(1).
                                        SetAlign(tview.AlignCenter))
    }
}

func initNavigation(session *core.Session) *Navigation {
    return &Navigation{
        widget: tview.NewTable().
                    SetSelectable(false, true).
                    SetFixed(1, 1).
                    SetSelectionChangedFunc(func(row, column int) {
                        switch currentWidget {
                        case "overview":
                            tui.layout.RemoveItem(tui.overview.widget)
                        case "files":
                            tui.layout.RemoveItem(tui.files.widget)
                        case "trackers":
                            tui.layout.RemoveItem(tui.trackers.widget)
                        case "peers":
                            tui.layout.RemoveItem(tui.peers.widget)
                        }

                        currentWidget = strings.ToLower(tui.navigation.widget.
                                                                        GetCell(row, column).
                                                                        Text)

                        switch currentWidget {
                        case "overview":
                            tui.layout.AddItem(tui.overview.widget, 0, 1, false)
                        case "files":
                            tui.layout.AddItem(tui.files.widget, 0, 1, false)
                        case "trackers":
                            tui.layout.AddItem(tui.trackers.widget, 0, 1, false)
                        case "peers":
                            tui.layout.AddItem(tui.peers.widget, 0, 1, false)
                        }

                        redraw(session)
                    }),
    }
}

func (nav *Navigation) setKeys() {
    tui.navigation.widget.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'q':
            currentWidget = "torrents"
            tui.files.widget.Clear()
            tui.peers.widget.Clear()
            tui.pages.RemovePage("details")
            return nil

        case 'j':
            switch currentWidget {
            case "overview":
                row, col := tui.overview.widget.GetScrollOffset()
                tui.overview.widget.ScrollTo(row + 1, col)
                return nil

            case "trackers":
                row, col := tui.trackers.widget.GetScrollOffset()
                tui.trackers.widget.ScrollTo(row + 1, col)
                return nil

            case "peers":
                if tui.peers.widget.GetRowCount() > 0 {
                    tui.app.SetFocus(tui.peers.widget)
                    setSelectedCellStyle(tui.navigation.widget,
                                         tcell.StyleDefault.Background(tcell.ColorBlack))

                    setSelectedCellStyle(tui.peers.widget,
                                         tcell.StyleDefault.Background(tcell.ColorWhite).
                                                            Foreground(tcell.ColorBlack))

                    return nil
                }

            case "files":
                tui.app.SetFocus(tui.files.widget)
                setSelectedCellStyle(tui.navigation.widget,
                                     tcell.StyleDefault.Background(tcell.ColorBlack))

                setSelectedCellStyle(tui.files.widget,
                                     tcell.StyleDefault.Background(tcell.ColorWhite).
                                                        Foreground(tcell.ColorBlack))

                return nil
            }

        case 'k':
            switch currentWidget {
            case "overview":
                row, col := tui.overview.widget.GetScrollOffset()
                tui.overview.widget.ScrollTo(row - 1, col)
                return nil

            case "trackers":
                row, col := tui.trackers.widget.GetScrollOffset()
                tui.trackers.widget.ScrollTo(row - 1, col)
                return nil
            }

        case 'g':
            switch currentWidget {
            case "overview":
                tui.overview.widget.ScrollToBeginning()
                return nil

            case "trackers":
                tui.trackers.widget.ScrollToBeginning()
                return nil
            }

        case 'G':
            switch currentWidget {
            case "overview":
                tui.overview.widget.ScrollToEnd()
                return nil

            case "trackers":
                tui.trackers.widget.ScrollToEnd()
                return nil
            }
        }

        return event
    })
}
