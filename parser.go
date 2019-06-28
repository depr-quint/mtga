package mtga

import (
	"strings"

	"github.com/di-wu/mtga/thread"
)

// Parser is a structure that holds the parser's callbacks.
type Parser struct {
	// thread
	onThreadLog func(log thread.Log)
	Outgoing
	Incoming
}

// Parse parses a raw log (returned by the tails logs channel).
// It calls the callback that matches that parsed log.
func (parser *Parser) Parse(l RawLog) {
	if len(l.body) <= 1 {
		return
	}

	switch first, remaining := l.body[0], l.body[1:]; {
	case strings.HasPrefix(first, "[UnityCrossThreadLogger]"):
		threadLog := thread.NewLog(strings.TrimPrefix(first, "[UnityCrossThreadLogger]"), remaining)
		if parser.onThreadLog != nil {
			parser.onThreadLog(threadLog)
		}
		parser.parseTreadLog(threadLog)
	default:
		// log.Fatalf("Unparsed log: %s.\n%s\n", first, remaining)
	}
}

func (parser *Parser) OnTreadLog(callback func(log thread.Log)) {
	parser.onThreadLog = callback
}

func (parser *Parser) parseTreadLog(l thread.Log) {
	if len(l.Json) <= 2 {
		return
	}

	switch l.Type {
	case thread.Outgoing:
		parser.parseOutgoingThreadLog(l)
	case thread.Incoming:
		parser.parseIncomingThreadLog(l)
	default:
		// log.Fatalf("Unparsed log: %s.\n", l.Type)
	}
}