package tui

import (
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
    "downloadLimited", "files",
}

func (overview *Overview) update(session *core.Session) {
    // torrent, err := core.GetTorrentByID(session, overview.id,  overviewFields)
    // if err != nil {
    //     currentWidget = "torrents"
    //     redraw(session)
    //     tui.pages.RemovePage("details")
    // }
}

func initOverview() *Overview {
    return &Overview{
        widget: tview.NewTextView().SetScrollable(true),
    }
}
