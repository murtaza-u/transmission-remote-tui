package core

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	TagTorrentList    = 7
	TagTorrentDetails = 77
	TagSessionStats   = 21
	TagSessionGet     = 22
	TagSessionClose   = 23
)

type Arguments map[string]interface{}

type RequestBody struct {
	Method    string    `json:"method"`
	Tag       string    `json:"tag"`
	Arguments Arguments `json:"arguments"`
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func SendRequest(method, tag string, arguments Arguments, session *Session) Response {
	encoded, err := json.Marshal(RequestBody{method, tag, arguments})
	HandleError(err)

	requestBody := bufio.NewReader(bytes.NewReader(encoded))
	client := &http.Client{}
	var body []byte

	for {
		request, err := http.NewRequest("POST", session.URL, requestBody)
		request.SetBasicAuth(session.Username, session.Password)
		HandleError(err)

		request.Header.Add("X-Transmission-Session-Id", session.ID)

		resp, err := client.Do(request)
		HandleError(err)

		body, err = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		HandleError(err)

		if !session.IsExpired(string(body)) {
			break
		}

		session.NewSessionID()
	}

	args := Response{}
	json.Unmarshal(body, &args)
	return args
}
