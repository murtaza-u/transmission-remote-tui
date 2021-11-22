package tui

import "github.com/rivo/tview"

type Peers struct {
    widget *tview.TextView
}

func initPeers() *Peers {
    return &Peers{
        widget: tview.NewTextView(),
    }
}
