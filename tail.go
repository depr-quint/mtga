package mtga

import (
	"bufio"
	"bytes"
	"github.com/fsnotify/fsnotify"
	"io"
	"os"
	"strings"
	"sync"
	"time"
	"unicode"
)

const (
	bufferSize = 4 * 1024
	peekSize   = 1024
)

// Tail can monitor data streams and open files, displaying new information as it is written.
// For example, it's a useful way to monitor the newest events in a system log in real time.
type Tail struct {
	once     sync.Once
	file     *os.File
	filePath string
	logs     chan RawLog
	err      error
	reader   *bufio.Reader
	watcher  *fsnotify.Watcher
	offset   int64
	closeCh  chan struct{}
}

// NewTail creates a new tail that monitors the file located at the given file path.
func NewTail(filePath string) (*Tail, error) {
	t := &Tail{
		filePath: filePath,
		logs:     make(chan RawLog),
		closeCh:  make(chan struct{}),
	}

	err := t.open()
	if err != nil {
		return nil, err
	}

	go t.once.Do(t.run)

	return t, nil
}

// Logs returns a channel of logs read from the monitored file.
func (t *Tail) Logs() chan RawLog {
	return t.logs
}

// Err returns a channel of errors read from the monitored file.
func (t *Tail) Err() error {
	return t.err
}

// Close stops the monitoring.
func (t *Tail) Close() {
	t.closeCh <- struct{}{}
}

func (t *Tail) run() {
	t.close(t.tail())
}

func (t *Tail) close(err error) {
	t.err = err

	if t.file != nil {
		t.file.Close()
	}

	close(t.logs)
}

func (t *Tail) tail() error {
	var (
		err     error
		eventCh = make(chan fsnotify.Event)
		errCh   = make(chan error, 1)
	)

	t.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	defer t.watcher.Close()
	go t.watch(eventCh, errCh)
	t.watcher.Add(t.filePath)

	for {
		var l RawLog
		for {
			// remove null bytes.
			for {
				b, _ := t.reader.Peek(peekSize)
				// index of the last instance of \x00 in b.
				i := bytes.LastIndexByte(b, '\x00')
				if i > 0 {
					// skips the next i + 1 bytes, n is the number of bytes discarded.
					t.reader.Discard(i + 1)
				}
				if i+1 < peekSize {
					break
				}
			}

			s, err := t.reader.ReadBytes('\n')
			if err != nil && err != io.EOF {
				return err
			}
			// end of file.
			if err == io.EOF {
				length := len(s)
				// sets the offset for the next read on file to -length
				t.offset, err = t.file.Seek(-int64(length), io.SeekCurrent)
				if err != nil {
					return err
				}
				// discards any buffered data, resets all state.
				t.reader.Reset(t.file)
				break
			}

			// bundle lines into logs
			line := string(s)
			trim := strings.TrimSpace(line)
			// empty line
			if trim == "" {
				continue
			}

			// first line
			if l == nil {
				l = append(l, trim)
				continue
			}

			// TODO: clean this...
			if (unicode.IsLetter(rune(line[0])) && strings.ToLower(trim) != "true") ||
				(unicode.IsNumber(rune(line[0]))) ||
				(strings.HasPrefix(trim, "[") && len(trim) > 2) {
				t.logs <- l
				l = []string{trim}
			} else if strings.HasPrefix(trim, "(") && strings.HasSuffix(trim, ")") {
				t.logs <- l
				t.logs <- []string{trim[1 : len(trim)-1]}
				l = RawLog{}
			} else if strings.HasPrefix(trim, "<<<<<<<<<<") {
				t.logs <- l
				t.logs <- []string{trim}
				l = RawLog{}
			} else {
				l = append(l, trim)
			}
		}

		select {
		case event := <-eventCh:
			switch event.Op {
			case fsnotify.Chmod:
				fallthrough
			case fsnotify.Write:
				info, err := t.file.Stat()
				if err != nil {
					if !os.IsNotExist(err) {
						return err
					}
					// reset if file is missing
					if err := t.reset(); err != nil {
						return err
					}
					continue
				}
				// file became shorter...
				if t.offset > info.Size() {
					t.offset, err = t.file.Seek(0, io.SeekStart)
					if err != nil {
						return err
					}
					t.reader.Reset(t.file)
				}
				continue
			default:
				if err := t.reset(); err != nil {
					return err
				}
				continue
			}
		case err := <-errCh:
			return err
		case <-t.closeCh:
			t.watcher.Remove(t.filePath)
			return nil
		case <-time.After(10 * time.Second):
			info1, err := t.file.Stat()
			if err != nil && !os.IsNotExist(err) {
				return err
			}
			info2, err := os.Stat(t.filePath)
			if err != nil && !os.IsNotExist(err) {
				return err
			}

			if os.SameFile(info1, info2) {
				continue
			}

			if err := t.reset(); err != nil {
				return err
			}
			continue
		}
	}
}

func (t *Tail) open() error {
	if t.file != nil {
		t.file.Close()
		t.file = nil
	}

	file, err := os.Open(t.filePath)
	if err != nil {
		return err
	}

	t.file = file
	t.reader = bufio.NewReaderSize(t.file, bufferSize)

	return nil
}

func (t *Tail) reset() error {
	t.watcher.Remove(t.filePath)
	if err := t.open(); err != nil {
		return err
	}
	t.watcher.Add(t.filePath)
	return nil
}

func (t *Tail) watch(eventCh chan fsnotify.Event, errChan chan error) {
	for {
		select {
		case event, ok := <-t.watcher.Events:
			if !ok {
				return
			}
			// ignore writes
			if event.Op == fsnotify.Write {
				select {
				case eventCh <- event:
				default:
				}
			} else {
				select {
				case eventCh <- event:
				case err := <-t.watcher.Errors:
					errChan <- err
					return
				}
			}

		case err := <-t.watcher.Errors:
			errChan <- err
			return
		}
	}
}
