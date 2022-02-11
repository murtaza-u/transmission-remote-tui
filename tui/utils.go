package tui

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	Kib = 1024
	Mib = 1048576
	Gib = 1073741824
	Tib = 1099511627776
)

func parseBytes(b float64) string {
	if b >= Tib {
		return fmt.Sprintf("%.2f TiB", b/Tib)
	} else if b >= Gib {
		return fmt.Sprintf("%.2f GiB", b/Gib)
	} else if b >= Mib {
		return fmt.Sprintf("%.2f MiB", b/Mib)
	} else {
		return fmt.Sprintf("%.2f KiB", b/Kib)
	}
}

func parseTime(s float64) string {
	t := time.Unix(int64(s), 0)
	beggining := time.Unix(0, 0)
	diff := t.Sub(beggining)
	if diff <= 0 {
		return ""
	}

	return fmt.Sprintf("%s", diff)
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

	parsedTime := parseTime(diff.Seconds())
	if parsedTime == "" || parsedTime == "1s" {
		return local.String(), "[now]"
	}

	return local.String(), fmt.Sprintf("[%s ago]", parseTime(diff.Seconds()))
}

func setSelectedCellStyle(table *tview.Table, style tcell.Style) {
	table.SetSelectedStyle(style)
}
