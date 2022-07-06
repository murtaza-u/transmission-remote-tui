package main

import (
	"log"
	"os"
	"strings"

	"github.com/murtaza-u/trt/cli"
	"github.com/murtaza-u/trt/core"
	"github.com/murtaza-u/trt/tui"
)

func main() {
	f := new(cli.Flags)
	f.Parse()

	if !strings.HasPrefix(f.URL, "http") {
		f.URL = "http://" + f.URL
	}

	s := new(core.Session)
	s.URL = f.URL
	s.Username = f.Username
	s.Password = f.Password

	s.CompileRegex()
	err := s.NewID()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	tui := tui.InitTUI(s)
	err = tui.Run(s)
	if err != nil {
		log.Fatal(err)
	}
}
