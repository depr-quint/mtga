package mtga

import (
	"os"
	"time"
)

type watchEvent struct {
	size int64
}

type watcher struct {
	filePath string
	lastSize int64
	ticker   *time.Ticker
	events   chan watchEvent
	errors   chan error
}

func newWatcher(pathToFile string, tick time.Duration) *watcher {
	return &watcher{
		filePath: pathToFile,
		lastSize: int64(-1),
		ticker:   time.NewTicker(tick),
		events:   make(chan watchEvent),
		errors:   make(chan error),
	}
}

func (w *watcher) start() {
	go func() {
		w.tick()
		for range w.ticker.C {
			w.tick()
		}
	}()
}

func (w *watcher) tick() {
	size, err := w.size()
	if err != nil {
		w.errors <- err
	} else if size != w.lastSize {
		w.lastSize = size
		w.events <- watchEvent{
			size: size,
		}
	}
}

func (w *watcher) stop() {
	w.ticker.Stop()
}

func (w *watcher) size() (int64, error) {
	fi, err := os.Stat(w.filePath)
	if err != nil {
		return -1, err
	}
	return fi.Size(), nil
}
