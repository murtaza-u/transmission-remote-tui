package tui

import (
	"fmt"

	"github.com/Murtaza-Udaipurwala/trt/core"
	"github.com/rivo/tview"
)

type Trackers struct {
	widget *tview.TextView
}

func initTrackers() *Trackers {
	return &Trackers{
		widget: tview.NewTextView().SetScrollable(true).SetWrap(true),
	}
}

var trackerFields []string = []string{"trackerStats", "id"}

func (trackers *Trackers) update(session *core.Session) {
	torrent, err := core.GetTorrentByID(session, tui.id, trackerFields)
	if err != nil {
		currentWidget = "torrents"
		redraw(session)
		tui.pages.RemovePage("details")
		return
	}

	var content string
	trackerStats := torrent.TrackerStats
	for tier, stat := range trackerStats {
		announce := stat.Announce
		lastAnnounceTime, lastAnnounceTimeAgo := convertUnixTime(stat.LastAnnounceTime)

		nextAnnounceTime, nextAnnounceTimeAgo := convertUnixTime(stat.NextAnnounceTime)
		if nextAnnounceTime == "" {
			nextAnnounceTime = "never"
		}

		lastScrapeTime, lastScrapeTimeAgo := convertUnixTime(stat.LastScrapeTime)
		nextScrapeTime, nextScrapeTimeAgo := convertUnixTime(stat.NextScrapeTime)
		seeders := stat.SeederCount
		leechers := stat.LeecherCount
		peers := stat.LastAnnouncePeerCount

		content += fmt.Sprintf("\n\nTier %d", tier)
		content += fmt.Sprintf("\n\t%v", announce)
		content += fmt.Sprintf("\n\tLast announced:  %v %v", lastAnnounceTime, lastAnnounceTimeAgo)
		content += fmt.Sprintf("\n\tNext announce:   %v %v", nextAnnounceTime, nextAnnounceTimeAgo)
		content += fmt.Sprintf("\n\tLast Scraped:    %v %v", lastScrapeTime, lastScrapeTimeAgo)
		content += fmt.Sprintf("\n\tNext Scrape:     %v %v", nextScrapeTime, nextScrapeTimeAgo)
		content += fmt.Sprintf("\n\tTracker knows:   %d seeders, %d leechers", seeders, leechers)
		content += fmt.Sprintf("\n\tResult:          %d peers received", peers)
	}

	trackers.widget.SetText(content)
}
