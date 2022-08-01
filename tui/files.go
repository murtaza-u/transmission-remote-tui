package tui

import (
	"log"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/murtaza-u/trt/core"
	"github.com/rivo/tview"
)

type files struct {
	widget   *tview.Table
	fields   []string
	priority map[int]string
}

const (
	FileOffPriority    = "off"
	FileLowPriority    = "low"
	FileNormalPriority = "normal"
	FileHighPriority   = "high"
)

func initFiles(s *core.Session) *files {
	f := new(files)
	f.widget = tview.NewTable().SetSelectable(true, false).SetFixed(1, 1)

	f.fields = []string{"id", "files", "wanted", "priorities", "name"}

	f.priority = make(map[int]string, 3)
	f.priority[-1] = FileLowPriority
	f.priority[0] = FileNormalPriority
	f.priority[1] = FileHighPriority

	f.setHeaders()
	f.setKeys(s)
	return f
}

func (f *files) style() {
	f.widget.SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorBlack))
}

func (f *files) setHeaders() {
	var headers = []string{"Total Size", "Downloaded", "Priority", "Name"}
	for col, h := range headers {
		f.widget.SetCell(
			0, col, tview.NewTableCell(h).SetSelectable(false).SetExpansion(1),
		)
	}
}

func (f *files) redraw(s *core.Session) error {
	id := tui.torrent.currID()
	t, err := s.GetTorrentByID(id, f.fields)
	if err != nil {
		return err
	}

	for i, fi := range t.Files {
		attrs := make([]string, 0, 4)
		attrs = append(attrs, byteCountSI(fi.Length))
		attrs = append(attrs, byteCountSI(fi.BytesCompleted))

		p := f.priority[t.Priorities[i]]
		if t.Wanted[i] == 0 {
			p = FileOffPriority
		}

		attrs = append(attrs, p)
		attrs = append(attrs, filepath.Base(fi.Name))

		for col := 0; col <= 3; col++ {
			f.widget.SetCell(i+1, col, tview.NewTableCell(attrs[col]))
		}
	}

	for i := len(t.Files) + 1; i < f.widget.GetRowCount(); i++ {
		f.widget.RemoveRow(i)
	}

	return err
}

func (f *files) setKeys(s *core.Session) {
	f.widget.SetInputCapture(func(e *tcell.EventKey) *tcell.EventKey {
		switch e.Rune() {
		case 'q':
			tui.layout.focus(f.widget)
			return nil

		case 'k':
			r, _ := f.widget.GetSelection()
			if r == 1 {
				tui.layout.focus(f.widget)
				return nil
			}

		case 'g':
			f.widget.Select(1, 0)
			f.widget.ScrollToBeginning()
			return nil

		case 'G':
			f.widget.Select(f.widget.GetRowCount()-1, 0)
			f.widget.ScrollToEnd()
			return nil

		case 'i':
			f.stepPri(s, 1)
			return nil

		case 'd':
			f.stepPri(s, -1)
			return nil

		case 'o':
			f.setPri(s, FileOffPriority, true)
			return nil

		case 'l':
			f.setPri(s, FileLowPriority, true)
			return nil

		case 'n':
			f.setPri(s, FileNormalPriority, true)
			return nil

		case 'h':
			f.setPri(s, FileHighPriority, true)
			return nil

		case 'O':
			f.setPri(s, FileOffPriority, false)
			return nil

		case 'L':
			f.setPri(s, FileLowPriority, false)
			return nil

		case 'N':
			f.setPri(s, FileNormalPriority, false)
			return nil

		case 'H':
			f.setPri(s, FileHighPriority, false)
			return nil
		}

		return e
	})
}

func (f *files) stepPri(s *core.Session, dir int) {
	id := tui.torrent.currID()
	r, _ := f.widget.GetSelection()
	fn := []int{r - 1}

	curr := f.widget.GetCell(r, 2).Text
	var err error

	switch curr {
	case FileOffPriority:
		if dir < 0 {
			return
		}
		err = s.FilePriority(id, fn, FileLowPriority, true)

	case FileLowPriority:
		if dir < 0 {
			err = s.FilePriority(id, fn, FileLowPriority, false)
			break
		}
		err = s.FilePriority(id, fn, FileNormalPriority, true)

	case FileNormalPriority:
		if dir < 0 {
			err = s.FilePriority(id, fn, FileLowPriority, true)
			break
		}
		err = s.FilePriority(id, fn, FileHighPriority, true)

	case FileHighPriority:
		if dir > 0 {
			return
		}
		err = s.FilePriority(id, fn, FileNormalPriority, true)
	}

	if err != nil {
		log.Fatal(err)
	}

	tui.force <- struct{}{}
}

func (f *files) setPri(s *core.Session, pri string, one bool) {
	id := tui.torrent.currID()
	fn := make([]int, 0, 1)

	if one {
		r, _ := f.widget.GetSelection()
		fn = append(fn, r-1)
	} else {
		for i := 0; i <= f.widget.GetRowCount()-2; i++ {
			fn = append(fn, i)
		}
	}

	var want bool
	if pri != FileOffPriority {
		want = true
	}

	err := s.FilePriority(id, fn, pri, want)
	if err != nil {
		log.Fatal(err)
	}

	tui.force <- struct{}{}
}
