package mtga

import (
	"github.com/di-wu/mtga/thread"
)

// ToMatch is a structure that holds the parser's to match callbacks.
type ToMatch struct {
	// TODO: split/method?
	onToMatchThreadLog func(method thread.LogMethod, payload string)
}

func (parser *Parser) parseToMatchThreadLog(l thread.Log) {
	if parser.onToMatchThreadLog != nil {
		parser.onToMatchThreadLog(l.Method, string(l.Raw))
	}
}

// OnToMatchThreadLog attaches the given callback, which will be called on a match to thread log.
func (to *ToMatch) OnToMatchThreadLog(callback func(method thread.LogMethod, payload string)) {
	to.onToMatchThreadLog = callback
}
