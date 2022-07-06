package tui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/murtaza-u/trt/core"
	"github.com/rivo/tview"
)

type overviewWid struct {
	widget *tview.TextView
	fields []string
}

func initOverviewWid(s *core.Session) *overviewWid {
	o := new(overviewWid)
	o.widget = tview.NewTextView().SetScrollable(true).SetWrap(true)
	o.fields = []string{
		"name", "downloadDir", "isPrivate", "addedDate", "activityDate",
		"dateCreated", "startDate", "doneDate", "comment", "creator",
		"hashString", "totalSize", "leftUntilDone", "pieceCount", "pieceSize",
		"uploadLimit", "downloadLimit", "uploadLimited", "downloadLimited",
		"files", "id", "uploadedEver", "downloadedEver", "haveValid",
		"corruptEver",
	}

	o.setKeys()

	return o
}

func (o *overviewWid) setKeys() {
	o.widget.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		switch e.Rune() {
		case 'h', 'q':
			tui.pages.SwitchToPage(TorrentPage)
			return nil
		}

		return e
	})
}

func (o *overviewWid) redraw(s *core.Session) error {
	id := tui.torrent.currID()
	t, err := s.GetTorrentByID(id, o.fields)
	if err != nil {
		return err
	}

	var privacy string
	if t.IsPrivate {
		privacy = "Private torrent"
	} else {
		privacy = "Public torrent"
	}

	var downLim, upLim string
	if t.DownloadLimited {
		downLim = fmt.Sprint(t.DownloadLimit)
	} else {
		downLim = "No Limit"
	}

	if t.UploadLimited {
		upLim = fmt.Sprint(t.UploadLimit)
	} else {
		upLim = "No Limit"
	}

	plate := `
    Name:         %v
    ID:           %v
    Hash:         %v
    Location:     %v
    Chunks:       %v (around %v each)
    Privacy:      %v
    No. of files: %v

    ===========================================================================

    Size:            %v
    Left until done: %v
    Downloaded:      %v
    Verified:        %v
    Corrupt:         %v
    Uploaded:        %v

    ===========================================================================

    Download limit: %v
    Upload limit:   %v

    ===========================================================================

    Comment: %v
    Creator: %v

    ===========================================================================

    Created at:       %s
    Added at:         %s
    Started at:       %s
    Last activity at: %s
    Completed at:     %s
    `

	txt := fmt.Sprintf(
		plate, t.Name, t.ID, t.HashString, t.DownloadDir, t.PieceCount,
		byteCountSI(t.PieceSize), privacy, len(t.Files),
		byteCountSI(t.TotalSize), byteCountSI(t.LeftUntilDone),
		byteCountSI(t.DownloadedEver), byteCountSI(t.HaveValid),
		byteCountSI(t.CorruptEver), byteCountSI(t.UploadedEver), downLim,
		upLim, t.Comment, t.Creator, unixTAbs(t.DateCreated),
		unixTAbs(t.AddedDate), unixTAbs(t.StartDate), unixTAbs(t.ActivityDate),
		unixTAbs(t.DoneDate),
	)

	o.widget.SetText(txt)
	return nil
}
