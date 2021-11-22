package tui

import (
    "fmt"

    "github.com/Murtaza-Udaipurwala/trt/core"
    "github.com/rivo/tview"
)

type Overview struct {
    widget *tview.TextView
    id int
}

var overviewFields []string = []string {
    "name", "downloadDir", "isPrivate", "addedDate", "activityDate",
    "dateCreated", "startDate", "doneDate", "comment", "creator", "hashString",
    "totalSize", "leftUntilDone", "pieceCount", "pieceSize", "seedRatioLimit",
    "seedRatioMode", "uploadLimit", "downloadLimit", "uploadLimited",
    "downloadLimited", "files", "id",
}

func (overview *Overview) update(session *core.Session) {
    torrent, err := core.GetTorrentByID(session, overview.id,  overviewFields)
    if err != nil {
        currentWidget = "torrents"
        redraw(session)
        tui.pages.RemovePage("details")
    }

    name := torrent.Name
    hash := torrent.HashString
    location := torrent.DownloadDir

    content := fmt.Sprintf("Name: %s\nHash: %s\nLocation: %s", name, hash, location)
    overview.widget.SetText(content)
}

func initOverview() *Overview {
    return &Overview{
        widget: tview.NewTextView().SetScrollable(true),
    }
}
