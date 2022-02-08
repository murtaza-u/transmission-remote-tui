package tui

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Murtaza-Udaipurwala/trt/core"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.design/x/clipboard"
)

type List struct {
	widget   *tview.Table
	torrents []core.Torrent
}

const (
	statusCol = iota
	etaCol
	uploadRateCol
	downloadRateCol
	ratioCol
	peersCol
	seedersCol
	leechersCol
	sizeCol
	leftCol
	nameCol
)

func initTorrents() *List {
	return &List{
		widget: tview.NewTable().SetSelectable(true, false).SetFixed(1, 1),
	}
}

var torrentFields []string = []string{
	"id", "name", "status", "eta", "uploadRatio", "peersConnected",
	"totalSize", "rateUpload", "rateDownload", "leftUntilDone", "queuePosition",
	"bandwidthPriority", "trackerStats", "magnetLink",
}

func (torrents *List) update(session *core.Session) {
	torrents.torrents = core.SortTorrentsByQueuePosition(core.GetTorrents(session, torrentFields))

	for row, torrent := range torrents.torrents {
		status := core.TorrentStatus[torrent.Status]
		eta := fmt.Sprintf("%s", parseTime(float64(torrent.ETA)))
		uploadRate := fmt.Sprintf("%s/s", parseBytes(float64(torrent.RateUpload)))
		downloadRate := fmt.Sprintf("%s/s", parseBytes(float64(torrent.RateDownload)))
		seeders, leechers := core.GetSeedersLeechers(torrent.TrackerStats)
		size := parseBytes(float64(torrent.TotalSize))
		left := parseBytes(float64(torrent.LeftUntilDone))
		name := torrent.Name

		var ratio string
		if torrent.UploadRatio >= 0 {
			ratio = fmt.Sprintf("%.3f", torrent.UploadRatio)
		}

		var peers string
		if torrent.PeersConnected >= 0 {
			peers = fmt.Sprint(torrent.PeersConnected)
		}

		tui.torrents.widget.SetCell(row+1, statusCol, tview.NewTableCell(status))
		tui.torrents.widget.SetCell(row+1, etaCol, tview.NewTableCell(eta))
		tui.torrents.widget.SetCell(row+1, uploadRateCol, tview.NewTableCell(uploadRate))
		tui.torrents.widget.SetCell(row+1, downloadRateCol, tview.NewTableCell(downloadRate))
		tui.torrents.widget.SetCell(row+1, ratioCol, tview.NewTableCell(ratio))
		tui.torrents.widget.SetCell(row+1, peersCol, tview.NewTableCell(peers))
		tui.torrents.widget.SetCell(row+1, seedersCol, tview.NewTableCell(seeders))
		tui.torrents.widget.SetCell(row+1, leechersCol, tview.NewTableCell(leechers))
		tui.torrents.widget.SetCell(row+1, sizeCol, tview.NewTableCell(size))
		tui.torrents.widget.SetCell(row+1, leftCol, tview.NewTableCell(left))
		tui.torrents.widget.SetCell(row+1, nameCol, tview.NewTableCell(name).SetMaxWidth(50))
	}
}

func (torrents *List) setHeaders() {
	var headers []string = []string{
		"Status", "ETA", "Upload Rate", "Download Rate", "Ratio", "Peers",
		"Seeders", "Leechers", "Size", "Left", "Name",
	}

	for col, header := range headers {
		torrents.widget.
			SetCell(0, col, tview.NewTableCell(header).
				SetSelectable(false).
				SetTextColor(tcell.ColorYellow).
				SetExpansion(1))
	}
}

func (torrents *List) currentSelected() (*core.Torrent, error) {
	row, _ := torrents.widget.GetSelection()
	name := torrents.widget.GetCell(row, nameCol).Text
	for _, torrent := range torrents.torrents {
		if torrent.Name == name {
			return &torrent, nil
		}
	}

	return &core.Torrent{}, errors.New("Torrent not found")
}

func (torrents *List) currentSelectedID() (int, error) {
	torrent, err := torrents.currentSelected()
	if err != nil {
		return -1, err
	}

	return torrent.ID, nil
}

func (torrents *List) setKeys(session *core.Session) {
	torrents.widget.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			tui.app.Stop()
			return nil

		case 'p':
			id, err := torrents.currentSelectedID()
			if err != nil {
				return nil
			}
			core.PauseStartTorrent(id, session, torrents.torrents)
			torrents.update(session)
			return nil

		case 'r':
			id, err := torrents.currentSelectedID()
			if err != nil {
				return nil
			}
			row, _ := torrents.widget.GetSelection()
			torrents.widget.RemoveRow(row)
			core.RemoveTorrent(id, session, false)
			return nil

		case 'R':
			id, err := torrents.currentSelectedID()
			if err != nil {
				return nil
			}
			row, _ := torrents.widget.GetSelection()
			torrents.widget.RemoveRow(row)
			core.RemoveTorrent(id, session, true)
			return nil

		case 'v':
			id, err := torrents.currentSelectedID()
			if err != nil {
				return nil
			}
			core.VerifyTorrent(id, session)
			torrents.update(session)
			return nil

		case 'g':
			tui.torrents.widget.Select(1, 0)
			tui.torrents.widget.ScrollToBeginning()
			return nil

		case 'G':
			tui.torrents.widget.Select(torrents.widget.GetRowCount()-1, 0)
			tui.torrents.widget.ScrollToEnd()
			return nil

		case 'K':
			id, err := torrents.currentSelectedID()
			if err != nil {
				return nil
			}
			core.QueueMove("up", id, session)
			torrents.update(session)
			row, _ := torrents.widget.GetSelection()
			torrents.widget.Select(row-1, 0)
			return nil

		case 'J':
			id, err := torrents.currentSelectedID()
			if err != nil {
				return nil
			}
			core.QueueMove("down", id, session)
			torrents.update(session)
			row, _ := torrents.widget.GetSelection()
			torrents.widget.Select(row+1, 0)
			return nil

		case 'U':
			id, err := torrents.currentSelectedID()
			if err != nil {
				return nil
			}
			core.QueueMove("top", id, session)
			torrents.update(session)
			torrents.widget.Select(1, 0)
			return nil

		case 'D':
			id, err := torrents.currentSelectedID()
			if err != nil {
				return nil
			}
			core.QueueMove("bottom", id, session)
			torrents.update(session)
			torrents.widget.Select(torrents.widget.GetRowCount()-1, 0)
			return nil

		case 't':
			id, err := torrents.currentSelectedID()
			if err != nil {
				return nil
			}
			core.AskTrackersForMorePeers(id, session)
			return nil

		case 'm':
			torrent, err := torrents.currentSelected()
			if err != nil {
				return nil
			}
			magnetLink := torrent.MagnetLink
			clipboard.Write(clipboard.FmtText, []byte(magnetLink))
			return nil

		case 'l', rune(tcell.KeyEnter):
			id, err := torrents.currentSelectedID()
			if err != nil {
				return nil
			}
			tui.id = id
			_, col := tui.navigation.widget.GetSelection()
			currentWidget = strings.ToLower(tui.navigation.widget.GetCell(0, col).Text)
			redraw(session)
			tui.pages.AddAndSwitchToPage("details", tui.layout, true)
			return nil
		}

		return event
	})
}
