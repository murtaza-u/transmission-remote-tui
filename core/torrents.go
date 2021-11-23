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
    Port               int    `json:"port"`
    Progress           int    `json:"progress"`
    RateToClient       int    `json:"rateToClient"`
    RateToPeer         int    `json:"rateToPeer"`
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

type File struct {
    BytesCompleted int64  `json:"bytesCompleted"`
    Length         int64  `json:"length"`
    Name           string `json:"name"`
}

type Torrent struct {
    ActivityDate            int64          `json:"activityDate"`
    AddedDate               int64          `json:"addedDate"`
    BandwidthPriority       int            `json:"bandwidthPriority"`
    Comment                 string         `json:"comment"`
    CorruptEver             int            `json:"corruptEver"`
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
    ID                      int            `json:"id"`
    IsPrivate               bool           `json:"isPrivate"`
    LeftUntilDone           int64          `json:"leftUntilDone"`
    MagnetLink              string         `json:"magnetLink"`
    MetadataPercentComplete int            `json:"metadataPercentComplete"`
    Name                    string         `json:"name"`
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
    return SendRequest("torrent-get", fmt.Sprint(TagTorrentList), Arguments {"fields": fields}, session).Arguments.Torrents
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
    for i := 0; i < len(torrents) - 1; i ++ {
        for j := 0; j < len(torrents) - 1 - i; j ++ {
            if torrents[j].QueuePosition > torrents[j + 1].QueuePosition {
                temp := torrents[j]
                torrents[j] = torrents[j + 1]
                torrents[j + 1] = temp
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
    SendRequest("torrent-remove", "1", Arguments{"id": id, "delete-local-data": deleteLocalData}, session)
}

func VerifyTorrent(id int, session *Session) {
    SendRequest("torrent-verify", "1", Arguments{"id": id}, session)
}

func QueueMove(direction string, id int, session *Session) {
    SendRequest("queue-move-" + direction, "1", Arguments{"ids": id}, session)
}

func GetSeedersLeechers(trackerStats []TrackerStats) (string, string) {
    if len(trackerStats) == 0 {
        return "", ""
    }

    var seeders, leechers int

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
