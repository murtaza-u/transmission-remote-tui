package tui

import (
	"fmt"
	"math"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
    _ = iota * 3
    KB
    MB
    GB
    TB
)

func parseBytes(b float64) string {
    if b >= math.Pow(10, TB) {
        return fmt.Sprintf("%.2f TB", b * math.Pow(10, -TB))
    } else if b >= math.Pow(10, GB) {
        return fmt.Sprintf("%.2f GB", b * math.Pow(10, -GB))
    } else if b >= math.Pow(10, MB) {
        return fmt.Sprintf("%.2f MB", b * math.Pow(10, -MB))
    } else {
        return fmt.Sprintf("%.2f KB", b * math.Pow(10, -KB))
    }
}

func parseTime(s float64) string {
    t := time.Unix(int64(s), 0)
    beggining := time.Unix(0, 0)
    diff := t.Sub(beggining)
    if diff < 0 {
        return ""
    }

    // day = 24h
    // week = 168h
    // month = 720h
    // year = 8760h

    hours := diff.Hours()
    if hours >= 8760 {
        return fmt.Sprintf("%d year(s)", int(hours / 8760))
    } else if hours >= 720 {
        return fmt.Sprintf("%d month(s)", int(hours / 720))
    } else if hours >= 168 {
        return fmt.Sprintf("%d week(s)", int(hours / 168))
    } else if hours >= 24 {
        return fmt.Sprintf("%d day(s)", int(hours / 24))
    } else {
        return diff.String()
    }
}

func convertUnixTime(t int64) (string, string) {
    if t == 0 {
        return "", ""
    }
    local := time.Unix(t, 0)
    currentTime := time.Now()
    diff := currentTime.Sub(local)
    if diff < 0 {
        diff = local.Sub(currentTime)
        return local.String(), fmt.Sprintf("[in %s]", parseTime(diff.Seconds()))
    }

    return local.String(), fmt.Sprintf("[%s ago]", parseTime(diff.Seconds()))
}

func setSelectedCellStyle(table *tview.Table, style tcell.Style) {
    table.SetSelectedStyle(style)
}
