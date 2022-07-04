package cli

import "flag"

func ParseArgs() (username, password, url string, port int) {
	flag.IntVar(&port, "port", 9091, "rpc port")
	flag.StringVar(&username, "username", "", "username")
	flag.StringVar(&password, "password", "", "password")
	flag.StringVar(&url, "url", "http://localhost", "rpc url")
	flag.Parse()
	return
}
