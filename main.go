package main

import (
	"fmt"
	"strings"

	"github.com/Murtaza-Udaipurwala/trt/cli"
	"github.com/Murtaza-Udaipurwala/trt/core"
	"github.com/Murtaza-Udaipurwala/trt/tui"
)

func main() {
	username, password, url, port := cli.ParseArgs()

	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	session := core.Session{}
	session.URL = fmt.Sprintf("%s:%d/transmission/rpc", url, port)
	session.Username = username
	session.Password = password
	session.CompileRegex()
	session.NewSessionID()

	tui.Run(&session)
}
