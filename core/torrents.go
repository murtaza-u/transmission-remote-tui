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

	"github.com/rivo/tview"
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

var Torrents []interface{}

func GetTorrentID(table *tview.Table, col int) (int, int) {
    row, _ := table.GetSelection()
    name := table.GetCell(row, col).Text

    for _, torrent := range Torrents {
        t := torrent.(map[string]interface{})
        if t["name"].(string) == name {
            return int(t["id"].(float64)), row
        }
    }
    return -1, row
}

func IsTorrentPause(id int) (bool, error) {
    for _, torrent := range Torrents {
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
    "bandwidthPriority",
}

func GetTorrents(session *Session) {
    Torrents = SendRequest("torrent-get", fmt.Sprint(TagTorrentList), Arguments {"fields": torrentFields}, session)["torrents"].([]interface{})
}

func SortTorrentsByQueuePosition() {
    for i := 0; i < len(Torrents) - 1; i ++ {
        for j := 0; j < len(Torrents) - 1 - i; j ++ {
            t1 := Torrents[j].(map[string]interface{})
            t2 := Torrents[j + 1].(map[string]interface{})
            if t1["queuePosition"].(float64) > t2["queuePosition"].(float64) {
                temp := Torrents[j]
                Torrents[j] = Torrents[j + 1]
                Torrents[j + 1] = temp
            }
        }
    }
}

func PauseStartTorrent(table *tview.Table, col int, session *Session) {
    id, _ := GetTorrentID(table, col)
    isPaused, err := IsTorrentPause(id)
    HandleError(err)

    if isPaused {
        SendRequest("torrent-start", "1", Arguments{"id": id}, session)
    } else {
        SendRequest("torrent-stop", "1", Arguments{"id": id}, session)
    }
}

func RemoveTorrent(table *tview.Table, col int, session *Session, deleteLocalData bool) {
    id, row := GetTorrentID(table, col)
    SendRequest("torrent-remove", "1", Arguments{"id": id, "delete-local-data": deleteLocalData}, session)
    table.RemoveRow(row)
    if row == 1 {
        table.Select(row + 1, 0)
    } else {
        table.Select(row - 1, 0)
    }
}

func VerifyTorrent(table *tview.Table, col int, session *Session) {
    id, _ := GetTorrentID(table, col)
    SendRequest("torrent-verify", "1", Arguments{"id": id}, session)
}

func QueueMove(direction string, table *tview.Table, col int, session *Session) int {
    id, row := GetTorrentID(table, col)
    SendRequest("queue-move-" + direction, "1", Arguments{"ids": id}, session)

    switch direction {
    case "up":
        row --
    case "down":
        row ++
    case "top":
        row = 1
    case "bottom":
        row = table.GetRowCount() - 1
    }

    return row
}
