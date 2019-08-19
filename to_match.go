package mtga

import (
	"github.com/di-wu/mtga/thread"
)

type ToMatch struct {
	// TODO: split/method?
	onToMatchThreadLog func(method thread.LogMethod, payload string)
}

func (parser *Parser) parseToMatchThreadLog(l thread.Log) {
	if parser.onToMatchThreadLog != nil {
		parser.onToMatchThreadLog(l.Method, string(l.Raw))
	}
}

func (to *ToMatch) OnToMatchThreadLog(callback func(method thread.LogMethod, payload string)) {
	to.onToMatchThreadLog = callback
}
