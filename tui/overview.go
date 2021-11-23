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
    location := torrent.DownloadDir
    isPrivate := torrent.IsPrivate

    var privacy string

    if isPrivate {
        privacy = "Private torrent"
    } else {
        privacy = "Public torrent"
    }

    addedDate, addedDateAgo := convertUnixTime(torrent.AddedDate)
    activityDate, activityDateAgo := convertUnixTime(torrent.ActivityDate)
    creationDate, creationDateAgo := convertUnixTime(torrent.DateCreated)
    startDate, startDateAgo := convertUnixTime(torrent.StartDate)
    completionDate, completionDateAgo := convertUnixTime(torrent.DoneDate)
    comment := torrent.Comment
    creator := torrent.Creator
    hash := torrent.HashString
    size := parseBytes(float64(torrent.TotalSize))
    left := parseBytes(float64(torrent.LeftUntilDone))
    pieceCount := torrent.PieceCount
    pieceSize := parseBytes(float64(torrent.PieceSize))
    id := torrent.Id
    filesCount := len(torrent.Files)

    var downloadLimit, uploadLimit, seedRatioLimit string

    if torrent.DownloadLimited {
        downloadLimit = fmt.Sprint(torrent.DownloadLimit)
    } else {
        downloadLimit = "No limit"
    }

    if torrent.UploadLimited {
        uploadLimit = fmt.Sprint(torrent.UploadLimit)
    } else {
        uploadLimit = "No limit"
    }

    switch torrent.SeedRatioMode {
    case 0:
        seedRatioLimit = ""
    case 1:
        seedRatioLimit = fmt.Sprint(torrent.SeedRatioLimit)
    case 2:
        seedRatioLimit = "No limit"
    }

    var content string
    content += fmt.Sprintf("\n\tName:              %v", name)
    content += fmt.Sprintf("\n\tID:                %v", id)
    content += fmt.Sprintf("\n\tHash:              %v", hash)
    content += fmt.Sprintf("\n\tLocation:          %v", location)
    content += fmt.Sprintf("\n\tTotal Size:        %v", size)
    content += fmt.Sprintf("\n\tLeft until done:   %v", left)
    content += fmt.Sprintf("\n\tChunks:            %v (around %v each)", pieceCount, pieceSize)
    content += fmt.Sprintf("\n\tPrivacy:           %v", privacy)
    content += fmt.Sprintf("\n\tNo. of Files:      %v", filesCount)

    content += "\n\n=============================================================================\n"

    content += fmt.Sprintf("\n\tDownload limit:    %v", downloadLimit)
    content += fmt.Sprintf("\n\tUpload limit:      %v", uploadLimit)
    content += fmt.Sprintf("\n\tSeed ratio limit:  %v", seedRatioLimit)

    content += "\n\n=============================================================================\n"

    content += fmt.Sprintf("\n\tComment:           %v", comment)
    content += fmt.Sprintf("\n\tCreator:           %v", creator)

    content += "\n\n=============================================================================\n"

    content += fmt.Sprintf("\n\tCreated at:        %v %v", creationDate, creationDateAgo)
    content += fmt.Sprintf("\n\tAdded at:          %v %v", addedDate, addedDateAgo)
    content += fmt.Sprintf("\n\tstarted at:        %v %v", startDate, startDateAgo)
    content += fmt.Sprintf("\n\tactivity at:       %v %v", activityDate, activityDateAgo)
    content += fmt.Sprintf("\n\tcompleted at:      %v %v", completionDate, completionDateAgo)

    overview.widget.SetText(content)
}

func initOverview() *Overview {
    return &Overview{
        widget: tview.NewTextView().SetScrollable(true).SetWrap(true),
    }
}
