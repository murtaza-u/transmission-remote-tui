package main

import (
    "fmt"

    "github.com/Murtaza-Udaipurwala/trt/cli"
    "github.com/Murtaza-Udaipurwala/trt/core"
    "github.com/Murtaza-Udaipurwala/trt/tui"
)

func main() {
    username, password, port := cli.ParseArgs()

    session := core.Session{}
    session.URL = fmt.Sprintf("http://127.0.0.1:%d/transmission/rpc", port)
    session.Username = username
    session.Password = password
    session.CompileRegex()
    session.NewSessionID()

    tui.Run(&session)
}
