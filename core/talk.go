package core

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	ErrConnectionFailed = errors.New("Could not connect to Transmission Daemon")
	ErrAuthFailed       = errors.New("Authentication Failed")
)

type ReqArgs map[string]any

type Req struct {
	Method string  `json:"method"`
	Tag    string  `json:"tag"`
	Args   ReqArgs `json:"arguments"`
}

type RespArgs struct {
	Torrents []Torrent `json:"torrents"`
}

type Resp struct {
	Args   RespArgs `json:"arguments"`
	Result string   `json:"result"`
}

func (s *Session) Talk(r *Req) (*Resp, error) {
	buff, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	b := bufio.NewReader(bytes.NewReader(buff))
	c := new(http.Client)
	var body []byte

	for {
		req, err := http.NewRequest(http.MethodPost, s.URL, b)
		if err != nil {
			return nil, err
		}

		req.SetBasicAuth(s.Username, s.Password)
		req.Header.Add("X-Transmission-Session-Id", s.ID)

		resp, err := c.Do(req)
		if err != nil {
			return nil, err
		}

		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, err
		}

		if !s.IsExpired(string(body)) {
			break
		}

		err = s.NewID()
		if err != nil {
			return nil, err
		}
	}

	v := new(Resp)
	err = json.Unmarshal(body, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
