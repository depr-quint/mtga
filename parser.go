package mtga

import (
	"fmt"
	"strings"

	"github.com/di-wu/mtga/thread"
)

// Parser is a structure that holds the parser's callbacks.
type Parser struct {
	onZoneChange func(change ZoneChange)
	// thread
	onSingleTreadLog func(first string)
	Single
	onThreadLog func(log thread.Log)
	Outgoing
	Incoming
	// unknown
	onUnknownLog func(message string)
}

// Parse parses a raw log (returned by the tails logs channel).
// It calls the callback that matches that parsed log.
func (parser *Parser) Parse(l RawLog) {
	if len(l.Body) == 0 {
		return
	}

	first := strings.TrimSpace(l.Body[0])
	if len(l.Body) == 1 {
		switch {
		case strings.HasPrefix(first, "[UnityCrossThreadLogger]"):
			log := strings.TrimPrefix(first, "[UnityCrossThreadLogger]")
			if parser.onSingleTreadLog != nil {
				parser.onSingleTreadLog(log)
			}
			parser.parseSingleTreadLog(log)
		case strings.HasPrefix(first, "<<<<<<<<<<"):
			parser.parserZoneChange(first)
		default:
			if parser.onUnknownLog != nil {
				parser.onUnknownLog(fmt.Sprintf("Unparsed log: %s", first))
			}
		}
		return
	}

	switch remaining := l.Body[1:]; {
	case strings.HasPrefix(first, "[UnityCrossThreadLogger]"):
		threadLog := thread.NewLog(strings.TrimPrefix(first, "[UnityCrossThreadLogger]"), remaining)
		if threadLog.Type == "" {
			if parser.onUnknownLog != nil {
				parser.onUnknownLog(fmt.Sprintf("Unparsed thread log: %s\n%s", first, remaining))
			}
		}
		if parser.onThreadLog != nil {
			parser.onThreadLog(threadLog)
		}
		parser.parseTreadLog(threadLog)
	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed log: %s\n%s", first, remaining))
		}
	}
}

func (parser *Parser) OnSingleTreadLog(callback func(log string)) {
	parser.onSingleTreadLog = callback
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
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed thread log: %s.\n", l.Type))
		}
	}
}

func (parser *Parser) OnUnknownLog(callback func(message string)) {
	parser.onUnknownLog = callback
}
