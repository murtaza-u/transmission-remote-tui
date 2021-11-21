package core

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type Session struct {
    ID string
    Regex *regexp.Regexp
}

func (session *Session) NewSessionID() {
    resp, err := http.Get(URL)
    HandleError(err)
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    HandleError(err)

    match := session.Regex.FindString(string(body))
    session.ID = strings.Split(match, ":")[1]
}

func (session *Session) CompileRegex() {
    r, err := regexp.Compile("X-Transmission-Session-Id:\\s*(\\w+)")
    HandleError(err)
    session.Regex = r
}

func (session *Session) IsExpired(body string) bool {
    match := session.Regex.FindString(body)
    if match == "" {
        return false
    }
    return true
}
