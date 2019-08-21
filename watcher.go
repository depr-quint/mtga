package mtga

import (
	"time"
)

type watchEvent struct {}

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

func (w *watcher) stop() {
	w.ticker.Stop()
}
