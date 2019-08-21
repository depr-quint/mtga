package mtga

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/connect_resp"
)

// Parser is a structure that holds all the parser's callbacks.
type Parser struct {
	onZoneChange  func(change ZoneChange)
	onConnectResp func(resp connect_resp.Response)
	// thread
	onSingleTreadLog func(first string)
	Single
	onThreadLog func(log thread.Log)
	Outgoing
	Incoming
	Unhandled
	MinusOne
	ToMatch
	MatchTo
	// unknown
	onUnknownLog func(message string)
}

// Parse parses a raw log (returned by the tails logs channel).
// It calls the callback that matches that parsed log.
func (parser *Parser) Parse(l RawLog) {
	// empty log
	if len(l) == 0 {
		return
	}

	first := strings.TrimSpace(l[0])
	if len(l) == 1 {
		switch {
		case strings.HasPrefix(first, "[UnityCrossThreadLogger]"):
			line := strings.TrimPrefix(first, "[UnityCrossThreadLogger]")
			if parser.onSingleTreadLog != nil {
				parser.onSingleTreadLog(line)
			}
			parser.parseSingleTreadLog(line)
		case strings.HasPrefix(first, "<<<<<<<<<<"):
			parser.parserZoneChange(first)
		default:
			if parser.onUnknownLog != nil {
				parser.onUnknownLog(fmt.Sprintf("Unparsed log: %s", first))
			}
		}
		return
	}

	switch remaining := l[1:]; {
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
		parser.parseMultilineTreadLog(threadLog)
	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed log: %s\n%s", first, remaining))
		}
	}
}

// OnSingleLineTreadLog attaches the given callback, which will be called on single line thread log.
func (parser *Parser) OnSingleLineTreadLog(callback func(log string)) {
	parser.onSingleTreadLog = callback
}

// OnTreadLog attaches the given callback, which will be called on every thread log.
func (parser *Parser) OnTreadLog(callback func(log thread.Log)) {
	parser.onThreadLog = callback
}

func (parser *Parser) parseMultilineTreadLog(l thread.Log) {
	if len(l.Raw) <= 2 {
		return
	}

	switch l.Type {
	case thread.Outgoing:
		parser.parseOutgoingThreadLog(l)
	case thread.Incoming:
		parser.parseIncomingThreadLog(l)
	case thread.ConnectResp:
		if parser.onConnectResp != nil {
			var resp connect_resp.Response
			err := json.Unmarshal(l.Raw, &resp)
			if err != nil {
				log.Fatalln(err)
			}
			parser.onConnectResp(resp)
		}
	case thread.Unhandled:
		parser.parseUnhandledThreadLog(l)
	case thread.MinusOne:
		parser.parseMinusOneThreadLog(l)
	case thread.ToMatch:
		parser.parseToMatchThreadLog(l)
	case thread.MatchTo:
		parser.parseMatchToThreadLog(l)
	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed thread log: %s\n%s", l.Type, string(l.Raw)))
		}
	}
}

// OnConnectResponse attaches the given callback, which will be called on getting the connection response.
func (parser *Parser) OnConnectResponse(callback func(resp connect_resp.Response)) {
	parser.onConnectResp = callback
}

// OnUnknownLog attaches the given callback, which will be called on unparsed logs.
func (parser *Parser) OnUnknownLog(callback func(message string)) {
	parser.onUnknownLog = callback
}
