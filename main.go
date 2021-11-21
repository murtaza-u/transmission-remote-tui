package main

import (
    "github.com/Murtaza-Udaipurwala/trt/core"
    "github.com/Murtaza-Udaipurwala/trt/tui"
)

func main() {
    session := core.Session{}
    session.CompileRegex()
    session.NewSessionID()

    tui.Run(&session)
}
