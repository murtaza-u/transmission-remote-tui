package core

import (
	"errors"
	"fmt"
)

type Torrent struct {
	ID                      int            `json:"id"`
	Name                    string         `json:"name"`
	ActivityDate            int64          `json:"activityDate"`
	AddedDate               int64          `json:"addedDate"`
	BandwidthPriority       int            `json:"bandwidthPriority"`
	Comment                 string         `json:"comment"`
	CorruptEver             int64          `json:"corruptEver"`
	Creator                 string         `json:"creator"`
	DateCreated             int64          `json:"dateCreated"`
	DesiredAvailable        int            `json:"desiredAvailable"`
	DoneDate                int64          `json:"doneDate"`
	DownloadDir             string         `json:"downloadDir"`
	DownloadLimit           int            `json:"downloadLimit"`
	DownloadLimited         bool           `json:"downloadLimited"`
	DownloadedEver          int64          `json:"downloadedEver"`
	ErrorString             string         `json:"errorString"`
	ETA                     int64          `json:"eta"`
	HashString              string         `json:"hashString"`
	HaveUnchecked           int64          `json:"haveUnchecked"`
	HaveValid               int64          `json:"haveValid"`
	HonorsSessionLimits     bool           `json:"honorsSessionLimits"`
	IsPrivate               bool           `json:"isPrivate"`
	LeftUntilDone           int64          `json:"leftUntilDone"`
	MagnetLink              string         `json:"magnetLink"`
	MetadataPercentComplete int            `json:"metadataPercentComplete"`
	UploadLimit             int            `json:"uploadLimit"`
	UploadLimited           bool           `json:"uploadLimited"`
	UploadRatio             float64        `json:"uploadRatio"`
	UploadedEver            int64          `json:"uploadedEver"`
	Peers                   []Peer         `json:"peers"`
	PeersConnected          int            `json:"peersConnected"`
	QueuePosition           int            `json:"queuePosition"`
	RateDownload            int64          `json:"rateDownload"`
	RateUpload              int64          `json:"rateUpload"`
	RecheckProgress         int            `json:"recheckProgress"`
	SeedRatioLimit          int            `json:"seedRatioLimit"`
	SeedRatioMode           int            `json:"seedRatioMode"`
	SizeWhenDone            int64          `json:"sizeWhenDone"`
	StartDate               int64          `json:"startDate"`
	Status                  int            `json:"status"`
	TotalSize               int64          `json:"totalSize"`
	TrackerStats            []TrackerStats `json:"trackerStats"`
	Files                   []File         `json:"files"`
	PieceCount              int64          `json:"pieceCount"`
	PieceSize               int64          `json:"pieceSize"`
	Priorities              []int          `json:"priorities"`
	Wanted                  []int          `json:"wanted"`
}

type File struct {
	BytesCompleted int64  `json:"bytesCompleted"`
	Length         int64  `json:"length"`
	Name           string `json:"name"`
}

type TrackerStats struct {
	Announce              string `json:"announce"`
	AnnounceState         int    `json:"announceState"`
	DownloadCount         int    `json:"downloadCount"`
	HasAnnounced          bool   `json:"hasAnnounced"`
	HasScraped            bool   `json:"hasScraped"`
	Host                  string `json:"host"`
	Id                    int    `json:"id"`
	IsBackup              bool   `json:"isBackup"`
	LastAnnouncePeerCount int    `json:"lastAnnouncePeerCount"`
	LastAnnounceResult    string `json:"lastAnnounceResult"`
	LastAnnounceStartTime int64  `json:"lastAnnounceStartTime"`
	LastAnnounceSucceeded bool   `json:"lastAnnounceSucceeded"`
	LastAnnounceTime      int64  `json:"lastAnnounceTime"`
	LastAnnounceTimedOut  bool   `json:"lastAnnounceTimedOut"`
	LastScrapeResult      string `json:"lastScrapeResult"`
	LastScrapeStartTime   int64  `json:"lastScrapeStartTime"`
	LastScrapeSucceeded   bool   `json:"lastScrapeSucceeded"`
	LastScrapeTime        int64  `json:"lastScrapeTime"`
	LastScrapeTimedOut    bool   `json:"lastScrapeTimedOut"`
	LeecherCount          int    `json:"leecherCount"`
	NextAnnounceTime      int64  `json:"nextAnnounceTime"`
	NextScrapeTime        int64  `json:"nextScrapeTime"`
	Scrape                string `json:"scrape"`
	ScrapeState           int    `json:"scrapeState"`
	SeederCount           int    `json:"seederCount"`
	Tier                  int    `json:"tier"`
}

type Peer struct {
	Address            string  `json:"address"`
	ClientIsChoked     bool    `json:"clientIsChoked"`
	ClientIsInterested bool    `json:"clientIsInterested"`
	ClientName         string  `json:"clientName"`
	FlagStr            string  `json:"flagStr"`
	IsDownloadingFrom  bool    `json:"isDownloadingFrom"`
	IsEncrypted        bool    `json:"isEncrypted"`
	IsIncoming         bool    `json:"isIncoming"`
	IsUTP              bool    `json:"isUTP"`
	IsUploadingTo      bool    `json:"isUploadingTo"`
	PeerIsChoked       bool    `json:"peerIsChoked"`
	PeerIsInterested   bool    `json:"peerIsInterested"`
	Port               int     `json:"port"`
	Progress           float64 `json:"progress"`
	RateToClient       int     `json:"rateToClient"`
	RateToPeer         int     `json:"rateToPeer"`
}

var TorrentStatus map[int]string = map[int]string{
	0: "Stopped",
	1: "Queued to check files",
	2: "Checking files",
	3: "Queued to download",
	4: "Downloading",
	5: "Queued to seed",
	6: "Seeding",
}

const (
	TagDefault        = "1"
	TagTorrentList    = "7"
	TagTorrentDetails = "77"
	TagSessionStats   = "21"
	TagSessionGet     = "22"
	TagSessionClose   = "23"

	MethodTorrentGet        = "torrent-get"
	MethodTorrentSet        = "torrent-set"
	MethodTorrentStart      = "torrent-start"
	MethodTorrentStop       = "torrent-stop"
	MethodTorrentRemove     = "torrent-remove"
	MethodTorrentVerify     = "torrent-verify"
	MethodTorrentReannounce = "torrent-reannounce"
	MethodSessionClose      = "session-close"
)

var ErrTorrentNotFound = errors.New("Torrent not found")

func (t *Torrent) GetSeederLeecher() (string, string) {
	if len(t.TrackerStats) == 0 {
		return "", ""
	}

	var s, l int
	for _, stat := range t.TrackerStats {
		s += stat.SeederCount
		l += stat.LeecherCount
	}

	var sstr, lstr string
	if s != 0 {
		sstr = fmt.Sprint(s)
	}

	if l != 0 {
		lstr = fmt.Sprint(l)
	}

	return sstr, lstr
}

func (t *Torrent) IsPaused() bool {
	return t.Status == 0
}

type Torrents struct {
	Ts []Torrent
}

func (ts *Torrents) GetTorrentByID(id int) (*Torrent, error) {
	for _, t := range ts.Ts {
		if t.ID == id {
			return &t, nil
		}
	}

	return nil, ErrTorrentNotFound
}

func (ts *Torrents) SortByQueuePosition() {
	n := len(ts.Ts)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if ts.Ts[j].QueuePosition > ts.Ts[j+1].QueuePosition {
				temp := ts.Ts[j]
				ts.Ts[j] = ts.Ts[j+1]
				ts.Ts[j+1] = temp
			}
		}
	}
}
