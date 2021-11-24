package cli

import "flag"

func ParseArgs() (username, password string, port int)  {
    flag.IntVar(&port, "port", 9091, "rpc port")
    flag.StringVar(&username, "username", "", "username")
    flag.StringVar(&password, "password", "", "password")
    flag.Parse()
    return
}
