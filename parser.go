package mtga

import (
	"encoding/json"
	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/outgoing"
	"log"
	"strings"
)

type Parser struct {
	onThreadLog    func(log thread.Log)
	onAuthenticate func(auth outgoing.Authenticate)
}

func (parser *Parser) ParseRawLog(l RawLog) {
	if len(l.body) <= 1 {
		return
	}

	switch first, remaining := l.body[0], l.body[1:]; {
	case strings.HasPrefix(first, "[UnityCrossThreadLogger]"):
		threadLog := thread.NewLog(strings.TrimPrefix(first, "[UnityCrossThreadLogger]"), remaining)
		if parser.onThreadLog != nil {
			parser.onThreadLog(threadLog)
		}
		parser.ParseTreadLog(threadLog)
	default:
		// log.Fatalf("Unparsed log: %s.\n%s\n", first, remaining)
	}
}

func (parser *Parser) OnTreadLog(callback func(log thread.Log)) {
	parser.onThreadLog = callback
}

func (parser *Parser) ParseTreadLog(l thread.Log) {
	if l.Json == nil {
		return
	}

	switch l.Type {
	case thread.Outgoing:
		parser.parseOutgoingThreadLog(l)
	}
}

func (parser *Parser) parseOutgoingThreadLog(l thread.Log) {
	switch l.Method {
	case thread.AuthenticateMethod:
		if (parser.onAuthenticate) != nil {
			var auth outgoing.Authenticate
			err := json.Unmarshal(l.Json, &auth)
			if err != nil {
				log.Fatalln(err)
			}
			parser.onAuthenticate(auth)
		}
	}
}

func (parser *Parser) OnAuthenticate(callback func(auth outgoing.Authenticate)) {
	parser.onAuthenticate = callback
}
