package tui

import "github.com/rivo/tview"

type Files struct {
    widget *tview.Table
}

func initFiles() *Files {
    return &Files{
        widget: tview.NewTable(),
    }
}