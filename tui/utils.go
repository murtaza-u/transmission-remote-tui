package tui

import (
    "fmt"
    "math"
)

const (
    _ = iota * 3
    KB
    MB
    GB
    TB
)

func convertBytesTo(b float64) string {
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

func convertSecondsTo(s float64) string {
    if s < 0 {
        return ""
    } else if s < 60 {
        return fmt.Sprintf("%ds", int(s))
    } else if s < 3600 {
        return fmt.Sprintf("%dm", int(s/60.0))
    } else if s < 86400 {
        return fmt.Sprintf("%dday(s)", int(s/3600.0))
    } else if s < 2592000 {
        return fmt.Sprintf("%dmonth(s)", int(s/85400.0))
    } else {
        return fmt.Sprintf("%dyear(s)", int(s/2592000.0))
    }
}
