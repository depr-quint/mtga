package mtga

import (
	"bufio"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

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

			if line := strings.TrimSpace(string(s)); line == "" {
				t.logs <- l
				l = RawLog{}
			} else {
				l.body = append(l.body, line)
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
