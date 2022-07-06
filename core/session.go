package core

import (
	"io"
	"net/http"
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

func (s *Session) NewID() error {
	c := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, s.URL, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(s.Username, s.Password)

	resp, err := c.Do(req)
	if err != nil {
		return ErrConnectionFailed
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	match := s.Regex.FindString(string(b))
	if match == "" {
		return ErrAuthFailed
	}

	s.ID = strings.Split(match, ":")[1]
	return nil
}

func (s *Session) CompileRegex() error {
	r, err := regexp.Compile("X-Transmission-Session-Id:\\s*(\\w+)")
	if err != nil {
		return err
	}

	s.Regex = r
	return nil
}

func (s *Session) IsExpired(body string) bool {
	match := s.Regex.FindString(body)
	return match != ""
}

func (s *Session) GetTorrents(fields []string) (*Torrents, error) {
	r := &Req{
		Method: MethodTorrentGet,
		Tag:    TagTorrentList,
		Args:   ReqArgs{"fields": fields},
	}
	resp, err := s.Talk(r)
	if err != nil {
		return nil, err
	}

	return &Torrents{resp.Args.Torrents}, nil
}

func (s *Session) GetTorrentByID(id int, fields []string) (*Torrent, error) {
	r := &Req{
		Method: MethodTorrentGet,
		Tag:    TagDefault,
		Args:   ReqArgs{"id": id, "fields": fields},
	}

	resp, err := s.Talk(r)
	if err != nil {
		return nil, err
	}

	return &resp.Args.Torrents[0], nil
}

func (s *Session) StartTorrent(id int) error {
	r := &Req{
		Method: MethodTorrentStart,
		Tag:    TagDefault,
		Args:   ReqArgs{"id": id},
	}
	_, err := s.Talk(r)
	return err
}

func (s *Session) StopTorrent(id int) error {
	r := &Req{
		Method: MethodTorrentStop,
		Tag:    TagDefault,
		Args:   ReqArgs{"id": id},
	}
	_, err := s.Talk(r)
	return err
}

func (s *Session) RemoveTorrent(id int, purge bool) error {
	r := &Req{
		Method: MethodTorrentRemove,
		Tag:    TagDefault,
		Args:   ReqArgs{"id": id, "delete-local-data": purge},
	}
	_, err := s.Talk(r)
	return err
}

func (s *Session) VerifyTorrent(id int) error {
	r := &Req{
		Method: MethodTorrentVerify,
		Tag:    TagDefault,
		Args:   ReqArgs{"id": id},
	}
	_, err := s.Talk(r)
	return err
}

func (s *Session) Reannounce(id int) error {
	r := &Req{
		Method: MethodTorrentReannounce,
		Tag:    TagDefault,
		Args:   ReqArgs{"id": id},
	}
	_, err := s.Talk(r)
	return err
}

func (s *Session) QueueMove(id int, dir string) error {
	r := &Req{
		Method: "queue-move-" + dir,
		Tag:    TagDefault,
		Args:   ReqArgs{"ids": id},
	}
	_, err := s.Talk(r)
	return err
}

func (s *Session) FilePriority(id int, fn []int, pri string, want bool) error {
	args := make(ReqArgs)
	args["ids"] = id
	args["priority-"+pri] = fn

	if want {
		args["files-wanted"] = fn
	} else {
		args["files-unwanted"] = fn
	}

	r := &Req{
		Method: MethodTorrentSet,
		Tag:    TagDefault,
		Args:   args,
	}
	_, err := s.Talk(r)
	return err
}

func (s *Session) Close() error {
	r := &Req{
		Method: MethodSessionClose,
		Tag:    TagDefault,
		Args:   ReqArgs{},
	}
	_, err := s.Talk(r)
	return err
}
