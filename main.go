package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/murtaza-u/trt/cli"
	"github.com/murtaza-u/trt/core"
	"github.com/murtaza-u/trt/tui"
)

const version = "1.0"

func main() {
	f := new(cli.Flags)
	f.Parse()

	if f.Version {
		fmt.Println(version)
		return
	}

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

	tui, err := tui.InitTUI(s)
	if err != nil {
		log.Fatal(err)
	}

	err = tui.Run(s)
	if err != nil {
		log.Fatal(err)
	}
}
