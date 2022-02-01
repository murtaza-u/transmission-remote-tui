// "id", "name", "downloadDir", "status", "trackerStats", "desiredAvailable",
// "rateDownload", "rateUpload", "eta", "uploadRatio", "sizeWhenDone",
// "haveValid", "haveUnchecked", "addedDate", "uploadedEver", "errorString",
// "recheckProgress", "peersConnected", "uploadLimit", "downloadLimit",
// "uploadLimited", "downloadLimited", "bandwidthPriority", "peersSendingToUs",
// "peersGettingFromUs", "seedRatioLimit", "seedRatioMode", "isPrivate",
// "magnetLink", "honorsSessionLimits", "metadataPercentComplete", "files",
// "priorities", "wanted", "peers", "trackers", "activityDate", "dateCreated",
// "startDate", "doneDate", "totalSize", "leftUntilDone", "comment", "creator",
// "hashString", "pieceCount", "pieceSize", "pieces", "downloadedEver",
// "corruptEver", "peersFrom", "queuePosition",

package core

import (
	"errors"
	"fmt"
)

type Response struct {
	Arguments RespArguments `json:"arguments"`
	Result    string        `json:"result"`
}

type RespArguments struct {
	Torrents []Torrent `json:"torrents"`
}

type Peer struct {
	Address            string `json:"address"`
	ClientIsChoked     bool   `json:"clientIsChoked"`
	ClientIsInterested bool   `json:"clientIsInterested"`
	ClientName         string `json:"clientName"`
	FlagStr            string `json:"flagStr"`
	IsDownloadingFrom  bool   `json:"isDownloadingFrom"`
	IsEncrypted        bool   `json:"isEncrypted"`
	IsIncoming         bool   `json:"isIncoming"`
	IsUTP              bool   `json:"isUTP"`
	IsUploadingTo      bool   `json:"isUploadingTo"`
	PeerIsChoked       bool   `json:"peerIsChoked"`
	PeerIsInterested   bool   `json:"peerIsInterested"`
	Port               uint   `json:"port"`
	Progress           uint   `json:"progress"`
	RateToClient       uint   `json:"rateToClient"`
	RateToPeer         uint   `json:"rateToPeer"`
}

type TrackerStats struct {
	Announce              string `json:"announce"`
	AnnounceState         uint   `json:"announceState"`
	DownloadCount         uint   `json:"downloadCount"`
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

type File struct {
	BytesCompleted uint64 `json:"bytesCompleted"`
	Length         uint64 `json:"length"`
	Name           string `json:"name"`
}

type Torrent struct {
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
	ETA                     uint64         `json:"eta"`
	HashString              string         `json:"hashString"`
	HaveUnchecked           uint64         `json:"haveUnchecked"`
	HaveValid               uint64         `json:"haveValid"`
	HonorsSessionLimits     bool           `json:"honorsSessionLimits"`
	ID                      int            `json:"id"`
	IsPrivate               bool           `json:"isPrivate"`
	LeftUntilDone           uint64         `json:"leftUntilDone"`
	MagnetLink              string         `json:"magnetLink"`
	MetadataPercentComplete uint           `json:"metadataPercentComplete"`
	Name                    string         `json:"name"`
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

var TorrentStatus map[int]string = map[int]string{
	0: "Stopped",
	1: "Queued to check files",
	2: "Checking files",
	3: "Queued to download",
	4: "Downloading",
	5: "Queued to seed",
	6: "Seeding",
}

func GetTorrentID(name string, col int, torrents []Torrent) int {
	for _, torrent := range torrents {
		if torrent.Name == name {
			return torrent.ID
		}
	}
	return -1
}

func IsTorrentPause(id int, torrents []Torrent) (bool, error) {
	for _, torrent := range torrents {
		if torrent.ID == id {
			return torrent.Status == 0, nil
		}
	}

	return false, errors.New("Torrent not found")
}

func GetTorrents(session *Session, fields []string) []Torrent {
	return SendRequest("torrent-get", fmt.Sprint(TagTorrentList),
		Arguments{"fields": fields}, session).Arguments.Torrents
}

func GetTorrentByID(session *Session, id int, fields []string) (Torrent, error) {
	torrents := GetTorrents(session, fields)
	for _, torrent := range torrents {
		if torrent.ID == id {
			return torrent, nil
		}
	}

	return Torrent{}, errors.New("Torrent not found")
}

func SortTorrentsByQueuePosition(torrents []Torrent) []Torrent {
	for i := 0; i < len(torrents)-1; i++ {
		for j := 0; j < len(torrents)-1-i; j++ {
			if torrents[j].QueuePosition > torrents[j+1].QueuePosition {
				temp := torrents[j]
				torrents[j] = torrents[j+1]
				torrents[j+1] = temp
			}
		}
	}
	return torrents
}

func PauseStartTorrent(id int, session *Session, torrents []Torrent) {
	isPaused, err := IsTorrentPause(id, torrents)
	HandleError(err)

	if isPaused {
		SendRequest("torrent-start", "1", Arguments{"id": id}, session)
	} else {
		SendRequest("torrent-stop", "1", Arguments{"id": id}, session)
	}
}

func RemoveTorrent(id int, session *Session, deleteLocalData bool) {
	SendRequest("torrent-remove", "1",
		Arguments{"id": id, "delete-local-data": deleteLocalData},
		session)
}

func VerifyTorrent(id int, session *Session) {
	SendRequest("torrent-verify", "1", Arguments{"id": id}, session)
}

func QueueMove(direction string, id int, session *Session) {
	SendRequest("queue-move-"+direction, "1", Arguments{"ids": id}, session)
}

func GetSeedersLeechers(trackerStats []TrackerStats) (string, string) {
	if len(trackerStats) == 0 {
		return "", ""
	}

	var seeders, leechers uint

	for _, stat := range trackerStats {
		seeders += stat.SeederCount
		leechers += stat.LeecherCount
	}

	if seeders < 0 && leechers < 0 {
		return "", ""
	} else if seeders < 0 {
		return "", fmt.Sprint(leechers)
	} else if leechers < 0 {
		return fmt.Sprint(seeders), ""
	}

	return fmt.Sprint(seeders), fmt.Sprint(leechers)
}

func ChangeFilePriority(fileNums []int, torrentID int, priority string, wanted bool, session *Session) {
	var args Arguments

	if wanted {
		args = Arguments{"ids": torrentID, "priority-" + priority: fileNums,
			"files-wanted": fileNums}
	} else {
		args = Arguments{"ids": torrentID, "priority-" + priority: fileNums,
			"files-unwanted": fileNums}
	}

	SendRequest("torrent-set", "1", args, session)
}

func AskTrackersForMorePeers(torrentID int, session *Session) {
	SendRequest("torrent-reannounce", "1", Arguments{"id": torrentID}, session)
}
