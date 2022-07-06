package tui

import (
	"fmt"

	"github.com/murtaza-u/trt/core"
	"github.com/rivo/tview"
)

type trackers struct {
	widget *tview.TextView
	fields []string
}

func initTrackers() *trackers {
	t := new(trackers)
	t.widget = tview.NewTextView().SetScrollable(true).SetWrap(true)
	t.fields = []string{"trackerStats", "id"}
	return t
}

func (t *trackers) redraw(s *core.Session) error {
	id := tui.torrent.currID()
	tor, err := s.GetTorrentByID(id, t.fields)
	if err != nil {
		return err
	}

	plate := `
    Tier %d
    %v
    Last announced: %v
    Next announce:  %v
    Last scraped:   %v
    Next scrape:    %v
    Tracker knows:  %d seeders, %d leechers
    %d peers received
	`

	var txt string
	for tier, s := range tor.TrackerStats {
		txt += fmt.Sprintf(
			plate, tier, s.Announce, unixTAbs(s.LastAnnounceTime),
			unixTAbs(s.NextAnnounceTime), unixTAbs(s.LastScrapeTime),
			unixTAbs(s.NextScrapeTime), s.SeederCount, s.LeecherCount,
			s.LastAnnouncePeerCount,
		)
	}

	t.widget.SetText(txt)
	return nil
}
