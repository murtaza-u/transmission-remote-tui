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

var TorrentStatus map[int]string = map[int]string{
    0: "Stopped",
    1: "Queued to check files",
    2: "Checking files",
    3: "Queued to download",
    4: "Downloading",
    5: "Queued to seed",
    6: "Seeding",
}

func GetTorrentID(name string, col int, torrents []interface{}) int {
    for _, torrent := range torrents {
        t := torrent.(map[string]interface{})
        if t["name"].(string) == name {
            return int(t["id"].(float64)) 
        }
    }
    return -1
}

func IsTorrentPause(id int, torrents []interface{}) (bool, error) {
    for _, torrent := range torrents {
        t := torrent.(map[string]interface{})
        if t["id"].(float64) == float64(id) {
            statusCode  := int((t["status"].(float64)))
            return statusCode == 0, nil
        }
    }

    return false, errors.New("Torrent not found")
}

var torrentFields []string = []string {
    "id", "name", "status", "eta", "uploadRatio", "peersConnected",
    "totalSize", "rateUpload", "rateDownload", "leftUntilDone","queuePosition",
    "bandwidthPriority", "trackerStats",
}

func GetTorrents(session *Session) []interface{} {
    return SendRequest("torrent-get", fmt.Sprint(TagTorrentList), Arguments {"fields": torrentFields}, session)["torrents"].([]interface{})
}

func SortTorrentsByQueuePosition(torrents []interface{}) []interface{} {
    for i := 0; i < len(torrents) - 1; i ++ {
        for j := 0; j < len(torrents) - 1 - i; j ++ {
            t1 := torrents[j].(map[string]interface{})
            t2 := torrents[j + 1].(map[string]interface{})
            if t1["queuePosition"].(float64) > t2["queuePosition"].(float64) {
                temp := torrents[j]
                torrents[j] = torrents[j + 1]
                torrents[j + 1] = temp
            }
        }
    }

    return torrents
}

func PauseStartTorrent(id int, session *Session, torrents []interface{}) {
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

func GetSeedersLeechers(trackerStats []interface{}) (string, string) {
    if len(trackerStats) == 0 {
        return "", ""
    }

    var seeders, leechers int

    for _, stat := range trackerStats {
        s := stat.(map[string]interface{})
        seeders += int(s["seederCount"].(float64))
        leechers += int(s["leecherCount"].(float64))
    }

    return fmt.Sprint(seeders), fmt.Sprint(leechers)
}
