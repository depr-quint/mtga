package main

import (
	"bufio"
	"io"
	"os"
	"sync"
	"time"
)

type tail struct {
	once     sync.Once
	file     *os.File
	filePath string
	lines    chan []byte
	reader   *bufio.Reader
	watcher  *watcher
	offset   int64
	done     chan bool
	err      error
}

func NewTail(filePath string) (*tail, error) {
	t := &tail{
		filePath: filePath,
		lines:    make(chan []byte),
		done:     make(chan bool),
	}

	if err := t.open(); err != nil {
		return nil, err
	}

	go t.once.Do(t.run)

	return t, nil
}

func (t *tail) open() error {
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

func (t *tail) run() {
	t.close(t.tail())
}

func (t *tail) close(err error) {
	t.err = err

	if t.file != nil {
		if err := t.file.Close(); err != nil {
			t.err = err
		}
	}

	close(t.lines)
}

func (t *tail) tail() error {
	var (
		events = make(chan event)
		errors = make(chan error, 1)
	)

	t.watcher = NewWatcher(t.filePath, time.Duration(5)*time.Second)
	defer t.watcher.stop()
	go t.watch(events, errors)

	for {
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

			t.lines <- s
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

func (t *tail) watch(events chan event, errors chan error) {
	for {
		select {
		case event := <-t.watcher.events:
			events <- event
		case err := <-t.watcher.errors:
			errors <- err
		}
	}
}
