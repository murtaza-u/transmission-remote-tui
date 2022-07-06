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
	BandwidthPriority       uint           `json:"bandwidthPriority"`
	Comment                 string         `json:"comment"`
	CorruptEver             uint64         `json:"corruptEver"`
	Creator                 string         `json:"creator"`
	DateCreated             int64          `json:"dateCreated"`
	DesiredAvailable        uint           `json:"desiredAvailable"`
	DoneDate                int64          `json:"doneDate"`
	DownloadDir             string         `json:"downloadDir"`
	DownloadLimit           uint           `json:"downloadLimit"`
	DownloadLimited         bool           `json:"downloadLimited"`
	DownloadedEver          uint64         `json:"downloadedEver"`
	ErrorString             string         `json:"errorString"`
	ETA                     int64          `json:"eta"`
	HashString              string         `json:"hashString"`
	HaveUnchecked           uint64         `json:"haveUnchecked"`
	HaveValid               uint64         `json:"haveValid"`
	HonorsSessionLimits     bool           `json:"honorsSessionLimits"`
	IsPrivate               bool           `json:"isPrivate"`
	LeftUntilDone           uint64         `json:"leftUntilDone"`
	MagnetLink              string         `json:"magnetLink"`
	MetadataPercentComplete uint           `json:"metadataPercentComplete"`
	UploadLimit             uint           `json:"uploadLimit"`
	UploadLimited           bool           `json:"uploadLimited"`
	UploadRatio             float64        `json:"uploadRatio"`
	UploadedEver            uint64         `json:"uploadedEver"`
	Peers                   []Peer         `json:"peers"`
	PeersConnected          uint           `json:"peersConnected"`
	QueuePosition           uint           `json:"queuePosition"`
	RateDownload            uint64         `json:"rateDownload"`
	RateUpload              uint64         `json:"rateUpload"`
	RecheckProgress         uint           `json:"recheckProgress"`
	SeedRatioLimit          uint           `json:"seedRatioLimit"`
	SeedRatioMode           uint           `json:"seedRatioMode"`
	SizeWhenDone            uint64         `json:"sizeWhenDone"`
	StartDate               int64          `json:"startDate"`
	Status                  int            `json:"status"`
	TotalSize               uint64         `json:"totalSize"`
	TrackerStats            []TrackerStats `json:"trackerStats"`
	Files                   []File         `json:"files"`
	PieceCount              uint64         `json:"pieceCount"`
	PieceSize               uint64         `json:"pieceSize"`
	Priorities              []int          `json:"priorities"`
	Wanted                  []int          `json:"wanted"`
}

type File struct {
	BytesCompleted uint64 `json:"bytesCompleted"`
	Length         uint64 `json:"length"`
	Name           string `json:"name"`
}

type TrackerStats struct {
	Announce              string `json:"announce"`
	AnnounceState         uint   `json:"announceState"`
	DownloadCount         int    `json:"downloadCount"`
	HasAnnounced          bool   `json:"hasAnnounced"`
	HasScraped            bool   `json:"hasScraped"`
	Host                  string `json:"host"`
	Id                    uint   `json:"id"`
	IsBackup              bool   `json:"isBackup"`
	LastAnnouncePeerCount uint   `json:"lastAnnouncePeerCount"`
	LastAnnounceResult    string `json:"lastAnnounceResult"`
	LastAnnounceStartTime uint64 `json:"lastAnnounceStartTime"`
	LastAnnounceSucceeded bool   `json:"lastAnnounceSucceeded"`
	LastAnnounceTime      int64  `json:"lastAnnounceTime"`
	LastAnnounceTimedOut  bool   `json:"lastAnnounceTimedOut"`
	LastScrapeResult      string `json:"lastScrapeResult"`
	LastScrapeStartTime   int64  `json:"lastScrapeStartTime"`
	LastScrapeSucceeded   bool   `json:"lastScrapeSucceeded"`
	LastScrapeTime        int64  `json:"lastScrapeTime"`
	LastScrapeTimedOut    bool   `json:"lastScrapeTimedOut"`
	LeecherCount          uint   `json:"leecherCount"`
	NextAnnounceTime      int64  `json:"nextAnnounceTime"`
	NextScrapeTime        int64  `json:"nextScrapeTime"`
	Scrape                string `json:"scrape"`
	ScrapeState           uint   `json:"scrapeState"`
	SeederCount           uint   `json:"seederCount"`
	Tier                  uint   `json:"tier"`
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
	Port               uint    `json:"port"`
	Progress           float64 `json:"progress"`
	RateToClient       uint    `json:"rateToClient"`
	RateToPeer         uint    `json:"rateToPeer"`
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
)

var ErrTorrentNotFound = errors.New("Torrent not found")

func (t *Torrent) GetSeederLeecher() (string, string) {
	if len(t.TrackerStats) == 0 {
		return "", ""
	}

	var s, l uint
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
