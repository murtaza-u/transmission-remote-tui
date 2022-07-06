package tui

import (
	"fmt"
	"time"
)

func byteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}

	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}

	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func unixTRel(t int64) string {
	t0 := time.Unix(0, 0)
	t1 := time.Unix(t, 0)
	diff := t1.Sub(t0)
	if diff < 0 {
		return ""
	}
	return diff.String()
}

func unixTAbs(t int64) string {
	if t == 0 {
		return ""
	}

	return time.Unix(t, 0).String()
}
