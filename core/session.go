package core

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Session struct {
	ID       string
	Regex    *regexp.Regexp
	Username string
	Password string
	URL      string
}

func (session *Session) NewSessionID() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", session.URL, nil)
	request.SetBasicAuth(session.Username, session.Password)
	resp, err := client.Do(request)

	if err != nil {
		log.Println("Transmission daemon not running")
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	HandleError(err)

	match := session.Regex.FindString(string(body))
	if match == "" {
		log.Println("Authentication failed")
		os.Exit(1)
	}
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
