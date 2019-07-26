package mtga

import (
	"bufio"
	"io"
	"os"
	"strings"
	"sync"
	"time"
	"unicode"
)

// Tail can monitor data streams and open files, displaying new information as it is written.
// For example, it's a useful way to monitor the newest events in a system log in real time.
type Tail struct {
	once     sync.Once
	file     *os.File
	filePath string
	logs     chan RawLog
	reader   *bufio.Reader
	watcher  *watcher
	offset   int64
	done     chan bool
	err      error
}

// NewTail creates a new tail that monitors the file located at the given file path.
func NewTail(filePath string) (*Tail, error) {
	t := &Tail{
		filePath: filePath,
		logs:     make(chan RawLog),
		done:     make(chan bool),
	}

	if err := t.open(); err != nil {
		return nil, err
	}

	go t.once.Do(t.run)

	return t, nil
}

// Logs returns a channel of (incoming) logs read from the monitored file.
func (t Tail) Logs() chan RawLog {
	return t.logs
}

func (t *Tail) open() error {
	if t.file != nil {
		if err := t.file.Close(); err != nil {
			return err
		}
		t.file = nil
	}

	if file, err := os.Open(t.filePath); err != nil {
		return err
	} else {
		t.file = file
		t.reader = bufio.NewReader(t.file)
		return nil
	}
}

func (t *Tail) run() {
	t.close(t.tail())
}

func (t *Tail) close(err error) {
	t.err = err

	if t.file != nil {
		if err := t.file.Close(); err != nil {
			t.err = err
		}
	}

	close(t.logs)
}

func (t *Tail) tail() error {
	var (
		events = make(chan watchEvent)
		errors = make(chan error, 1)
	)

	t.watcher = newWatcher(t.filePath, time.Duration(5)*time.Second)
	defer t.watcher.stop()
	go t.watch(events, errors)

	for {
		var l RawLog
		for {
			s, err := t.reader.ReadBytes('\n')
			if err != nil && err != io.EOF {
				return err
			}

			if err == io.EOF {
				l := len(s)

				t.offset, err = t.file.Seek(-int64(l), io.SeekCurrent)
				if err != nil {
					return err
				}

				t.reader.Reset(t.file)
				break
			}

			line := string(s)
			trim := strings.TrimSpace(line)
			// empty line
			if strings.TrimSpace(line) == "" {
				continue
			}

			// raw log is empty
			if l.body == nil {
				l.body = append(l.body, trim)
				continue
			}

			if (unicode.IsLetter(rune(line[0])) && strings.ToLower(trim) != "true") ||
				(unicode.IsNumber(rune(line[0]))) ||
				(strings.HasPrefix(trim, "[") && len(trim) > 2) {
				t.logs <- l
				l = RawLog{
					body: []string{trim},
				}
			} else if strings.HasPrefix(trim, "(") && strings.HasSuffix(trim, ")") {
				t.logs <- l
				t.logs <- RawLog{body: []string{trim[1 : len(trim)-1]}}
				l = RawLog{}
			} else if strings.HasPrefix(trim, "<<<<<<<<<<") {
				t.logs <- l
				t.logs <- RawLog{body: []string{trim}}
				l = RawLog{}
			} else {
				l.body = append(l.body, trim)
			}
		}

		select {
		case <-events:
			fi, err := t.file.Stat()
			if err != nil {
				if !os.IsNotExist(err) {
					return err
				}

				if err := t.open(); err != nil {
					return err
				}
				continue
			}

			if t.offset > fi.Size() {
				t.offset, err = t.file.Seek(0, io.SeekStart)
				if err != nil {
					return err
				}

				t.reader.Reset(t.file)
			}
			continue

		case err := <-errors:
			return err

		case <-t.done:
			t.watcher.stop()
			return nil

		case <-time.After(10 * time.Second):
			fi1, err := t.file.Stat()
			if err != nil && !os.IsNotExist(err) {
				return err
			}

			fi2, err := os.Stat(t.filePath)
			if err != nil && !os.IsNotExist(err) {
				return err
			}

			if os.SameFile(fi1, fi2) {
				continue
			}

			if err := t.open(); err != nil {
				return err
			}
			continue
		}
	}
}

func (t *Tail) watch(events chan watchEvent, errors chan error) {
	for {
		select {
		case event := <-events:
			events <- event
		case err := <-errors:
			errors <- err
		}
	}
}
