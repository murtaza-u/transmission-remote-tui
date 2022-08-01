package tui

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/murtaza-u/trt/core"
	"github.com/rivo/tview"
	"golang.design/x/clipboard"
)

// 30% of terminal width
const nameMaxWidth = 30

type torrentWid struct {
	widget   *tview.Table
	fields   []string
	kv       map[int]int
	torrents *core.Torrents
}

func initTorrentWid(s *core.Session) *torrentWid {
	tw := new(torrentWid)
	tw.kv = make(map[int]int)
	tw.widget = tview.NewTable().SetSelectable(true, false).SetFixed(1, 1)
	tw.setKeys(s)
	tw.fields = []string{
		"id", "name", "status", "eta", "uploadRatio", "peersConnected",
		"totalSize", "rateUpload", "rateDownload", "leftUntilDone",
		"queuePosition", "bandwidthPriority", "trackerStats", "magnetLink",
	}
	tw.setHeaders()

	return tw
}

func (tw *torrentWid) currID() int {
	row, _ := tw.widget.GetSelection()
	return tw.kv[row]
}

func (tw *torrentWid) currTorrent() (*core.Torrent, error) {
	id := tw.currID()
	t, err := tw.torrents.GetTorrentByID(id)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (tw *torrentWid) startStopTorrent(s *core.Session) {
	t, err := tw.currTorrent()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		tui.force <- struct{}{}
	}()

	if t.IsPaused() {
		err := s.StartTorrent(t.ID)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	err = s.StopTorrent(t.ID)
	if err != nil {
		log.Fatal(err)
	}
}

func (tw *torrentWid) queueMove(s *core.Session, dir string) {
	id := tw.currID()
	err := s.QueueMove(id, dir)
	if err != nil {
		log.Fatal(err)
	}

	tui.force <- struct{}{}

	var row int

	switch dir {
	case "up":
		row, _ = tw.widget.GetSelection()
		row--
	case "down":
		row, _ = tw.widget.GetSelection()
		row++
	case "top":
		row = 1
	case "bottom":
		row = tw.widget.GetRowCount() - 1
	}

	tw.widget.Select(row, 0)
}

func (tw *torrentWid) remove(s *core.Session, purge bool) {
	row, _ := tw.widget.GetSelection()
	tw.widget.RemoveRow(row)

	id := tw.currID()
	err := s.RemoveTorrent(id, purge)
	if err != nil {
		log.Fatal(err)
	}

	tui.force <- struct{}{}
}

func (tw *torrentWid) reannounce(s *core.Session) {
	id := tw.currID()
	err := s.Reannounce(id)
	if err != nil {
		log.Fatal(err)
	}
}

func (tw *torrentWid) copyMangnetLink() {
	t, err := tw.currTorrent()
	if err != nil {
		log.Fatal(err)
	}

	clipboard.Write(clipboard.FmtText, []byte(t.MagnetLink))
}

func (tw *torrentWid) verify(s *core.Session) {
	id := tw.currID()
	err := s.VerifyTorrent(id)
	if err != nil {
		log.Fatal(err)
	}

	tui.force <- struct{}{}
}

func (tw *torrentWid) setKeys(s *core.Session) {
	tw.widget.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		key := e.Rune()
		row, _ := tw.widget.GetSelection()
		count := tw.widget.GetRowCount()

		if key != 'q' && (count == 1 || count == row) {
			return e
		}

		switch key {
		case 'g':
			tw.widget.Select(1, 0)
			tw.widget.ScrollToBeginning()
			return nil

		case 'G':
			tw.widget.Select(tw.widget.GetRowCount()-1, 0)
			tw.widget.ScrollToEnd()
			return nil

		case 'K':
			tw.queueMove(s, "up")
			return nil

		case 'J':
			tw.queueMove(s, "down")
			return nil

		case 'U':
			tw.queueMove(s, "top")
			return nil

		case 'D':
			tw.queueMove(s, "bottom")
			return nil

		case 'p':
			tw.startStopTorrent(s)
			return nil

		case 'v':
			tw.verify(s)
			return nil

		case 't':
			tw.reannounce(s)
			return nil

		case 'm':
			tw.copyMangnetLink()
			return nil

		case 'r':
			tw.remove(s, false)
			return nil

		case 'R':
			tw.remove(s, true)
			return nil

		case 'l':
			tui.pages.SwitchToPage(DetailsPage)
			tui.force <- struct{}{}
			return nil

		case 'q':
			tui.app.Stop()
			return nil
		}

		return e
	})
}

func (tw *torrentWid) setHeaders() {
	var headers []string = []string{
		"Status", "ETA", "Upload Rate", "Download Rate", "Ratio", "Peers",
		"Seeders", "Leechers", "Size", "Left", "Name",
	}

	for col, h := range headers {
		tw.widget.SetCell(
			0, col, tview.NewTableCell(h).SetSelectable(false).SetExpansion(1),
		)
	}
}

func (tw *torrentWid) redraw(s *core.Session) error {
	torrents, err := s.GetTorrents(tw.fields)
	if err != nil {
		return err
	}
	torrents.SortByQueuePosition()

	for row, t := range torrents.Ts {
		attrs := make([]string, 0, 11)
		attrs = append(attrs, core.TorrentStatus[t.Status])
		attrs = append(attrs, unixTRel(t.ETA))
		attrs = append(attrs, byteCountSI(t.RateUpload)+"/s")
		attrs = append(attrs, byteCountSI(t.RateDownload)+"/s")

		var ratio string
		if t.UploadRatio >= 0 {
			ratio = fmt.Sprint(t.UploadRatio)
		}
		attrs = append(attrs, ratio)

		attrs = append(attrs, fmt.Sprint(t.PeersConnected))

		s, l := t.GetSeederLeecher()
		attrs = append(attrs, s)
		attrs = append(attrs, l)

		attrs = append(attrs, byteCountSI(t.TotalSize))
		attrs = append(attrs, byteCountSI(t.LeftUntilDone))
		attrs = append(attrs, t.Name)

		for col, attr := range attrs {
			cell := tview.NewTableCell(attr).SetSelectable(true)
			if col == len(attrs)-1 {
				cell.SetMaxWidth(tui.width * nameMaxWidth / 100)
			}

			tw.widget.SetCell(row+1, col, cell)
		}

		tw.kv[row+1] = t.ID
	}

	for i := len(torrents.Ts) + 1; i < tw.widget.GetRowCount(); i++ {
		tw.widget.RemoveRow(i)
	}

	tw.torrents = torrents
	return nil
}
