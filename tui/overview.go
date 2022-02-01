package tui

import (
	"fmt"

	"github.com/Murtaza-Udaipurwala/trt/core"
	"github.com/rivo/tview"
)

type Overview struct {
	widget *tview.TextView
}

var overviewFields []string = []string{
	"name", "downloadDir", "isPrivate", "addedDate", "activityDate",
	"dateCreated", "startDate", "doneDate", "comment", "creator", "hashString",
	"totalSize", "leftUntilDone", "pieceCount", "pieceSize", "uploadLimit",
	"downloadLimit", "uploadLimited", "downloadLimited", "files", "id",
	"uploadedEver", "downloadedEver", "haveValid", "corruptEver",
}

func (overview *Overview) update(session *core.Session) {
	torrent, err := core.GetTorrentByID(session, tui.id, overviewFields)
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

	downloaded := parseBytes(float64(torrent.DownloadedEver))
	downloadedPercentage := fmt.Sprintf("(%.2f%%)",
		float64(torrent.DownloadedEver)/float64(torrent.TotalSize)*100)

	var corrupted, corruptedPercentage string
	if torrent.CorruptEver == 0 {
		corrupted = "Nothing corrupt"
	} else {
		corrupted = parseBytes(float64(torrent.CorruptEver))
		corruptedPercentage = fmt.Sprintf("(%.2f%%)",
			float64(torrent.CorruptEver)/float64(torrent.TotalSize)*100)
	}

	verified := parseBytes(float64(torrent.HaveValid))
	verifiedPercentage := fmt.Sprintf("(%.2f%%)", float64(torrent.HaveValid)/
		float64(torrent.TotalSize)*100)

	uploaded := parseBytes(float64(torrent.UploadedEver))

	pieceCount := torrent.PieceCount
	pieceSize := parseBytes(float64(torrent.PieceSize))
	id := torrent.ID
	filesCount := len(torrent.Files)

	var downloadLimit, uploadLimit string

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

	var content string
	content += fmt.Sprintf("\n\tName:              %v", name)
	content += fmt.Sprintf("\n\tID:                %v", id)
	content += fmt.Sprintf("\n\tHash:              %v", hash)
	content += fmt.Sprintf("\n\tLocation:          %v", location)
	content += fmt.Sprintf("\n\tChunks:            %v (around %v each)", pieceCount, pieceSize)
	content += fmt.Sprintf("\n\tPrivacy:           %v", privacy)
	content += fmt.Sprintf("\n\tNo. of Files:      %v", filesCount)

	content += "\n\n============================================================================\n"

	content += fmt.Sprintf("\n\tTotal Size:        %v", size)
	content += fmt.Sprintf("\n\tLeft until done:   %v", left)
	content += fmt.Sprintf("\n\tDownloaded:        %v %v", downloaded, downloadedPercentage)
	content += fmt.Sprintf("\n\tVerified:          %v %v", verified, verifiedPercentage)
	content += fmt.Sprintf("\n\tCorrupt:           %v %v", corrupted, corruptedPercentage)
	content += fmt.Sprintf("\n\tUploaded:          %v", uploaded)

	content += "\n\n============================================================================\n"

	content += fmt.Sprintf("\n\tDownload limit:    %v", downloadLimit)
	content += fmt.Sprintf("\n\tUpload limit:      %v", uploadLimit)

	content += "\n\n============================================================================\n"

	content += fmt.Sprintf("\n\tComment:           %v", comment)
	content += fmt.Sprintf("\n\tCreator:           %v", creator)

	content += "\n\n============================================================================\n"

	content += fmt.Sprintf("\n\tCreated at:        %v %v", creationDate, creationDateAgo)
	content += fmt.Sprintf("\n\tAdded at:          %v %v", addedDate, addedDateAgo)
	content += fmt.Sprintf("\n\tstarted at:        %v %v", startDate, startDateAgo)
	content += fmt.Sprintf("\n\tLast activity at:  %v %v", activityDate, activityDateAgo)
	content += fmt.Sprintf("\n\tcompleted at:      %v %v", completionDate, completionDateAgo)

	overview.widget.SetText(content)
}

func initOverview() *Overview {
	return &Overview{
		widget: tview.NewTextView().SetScrollable(true).SetWrap(true),
	}
}
