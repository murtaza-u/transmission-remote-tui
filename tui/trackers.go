package tui

import "github.com/rivo/tview"

type Trackers struct {
    widget *tview.TextView
}

func initTrackers() *Trackers {
    return &Trackers{
        widget: tview.NewTextView(),
    }
}
