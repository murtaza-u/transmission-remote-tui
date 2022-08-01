package cli

import "flag"

type Flags struct {
	Username string
	Password string
	URL      string
	Version  bool
}

const defaultURL = "http://localhost:9091/transmission/rpc"

func (f *Flags) Parse() {
	flag.StringVar(&f.URL, "url", defaultURL, "RPC url")
	flag.StringVar(&f.Username, "username", "", "username")
	flag.StringVar(&f.Password, "password", "", "password")
	flag.BoolVar(&f.Version, "version", false, "version")
	flag.Parse()
}
