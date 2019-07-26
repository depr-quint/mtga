package mtga

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/di-wu/mtga/thread"
)

// Parser is a structure that holds the parser's callbacks.
type Parser struct {
	onZoneChange func(change ZoneChange)
	// thread
	onThreadLog func(log thread.Log)
	Outgoing
	Incoming
	// unknown
	onUnknownLog func(message string)
}

// Parse parses a raw log (returned by the tails logs channel).
// It calls the callback that matches that parsed log.
func (parser *Parser) Parse(l RawLog) {
	if len(l.body) == 0 {
		return
	}

	first := l.body[0]
	if strings.HasPrefix(first, "<<<<<<<<<<") {
		parser.parserZoneChange(first)
	}

	if len(l.body) == 1 {
		return
	}

	switch remaining := l.body[1:]; {
	case strings.HasPrefix(first, "[UnityCrossThreadLogger]"):
		threadLog := thread.NewLog(strings.TrimPrefix(first, "[UnityCrossThreadLogger]"), remaining)
		if parser.onThreadLog != nil {
			parser.onThreadLog(threadLog)
		}
		parser.parseTreadLog(threadLog)
	case strings.HasPrefix(first, "XInput"),
		strings.HasPrefix(first, "NullReferenceException"),
		strings.HasPrefix(first, "[Get SKUs]"),
		strings.HasPrefix(first, "[Client GRE]"),
		strings.HasPrefix(first, "Initialize engine version"),
		strings.HasPrefix(first, "Fallback handler"),
		strings.HasPrefix(first, "Unloading"),
		strings.HasPrefix(first, "Begin"),
		strings.HasPrefix(first, "Uploading"),
		strings.HasPrefix(first, "Setting up"),
		strings.HasPrefix(first, "WARNING"),
		strings.HasPrefix(first, "BIError"),
		strings.HasPrefix(first, "Direct3D"),
		strings.HasPrefix(first, "System.InvalidOperationException"):
		// ignore
	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed log: %s\n%s\n", first, remaining))
		}
	}
}

func (parser *Parser) OnZoneChange(callback func(change ZoneChange)) {
	parser.onZoneChange = callback
}

func (parser *Parser) parserZoneChange(first string) {
	if parser.onZoneChange != nil {
		str := regexp.MustCompile(`([a-zA-Z]*?) for \[\"([a-zA-Z ,\'-]*?)\" InstanceId:([0-9]*?), GrpId:([0-9]*?)\] \(\".*?\"\) had Instigator ([0-9]*?) \(\"([a-zA-Z ,\'-]*?)\"\)`).FindStringSubmatch(first)
		if str != nil {
			instanceId, _ := strconv.Atoi(str[3])
			grpId, _ := strconv.Atoi(str[4])
			instigator, _ := strconv.Atoi(str[5])
			parser.onZoneChange(ZoneChange{
				Type:       ZoneChangeType(str[1]),
				Target:     str[2],
				InstanceId: instanceId,
				GrpId:      grpId,
				Instigator: instigator,
				Source:     str[6],
			})
		} else {
			null := regexp.MustCompile(`([a-zA-Z]*?) for ([0-9]*?) \(\"\[NULL\]\"\) had Instigator ([0-9]*?) \(\"([a-zA-Z ,\'-]*?)\"\)`).FindStringSubmatch(first)
			instanceId, _ := strconv.Atoi(null[2])
			instigator, _ := strconv.Atoi(null[3])
			parser.onZoneChange(ZoneChange{
				Type:       ZoneChangeType(null[1]),
				Target:     "NULL",
				InstanceId: instanceId,
				Instigator: instigator,
				Source:     null[4],
			})
		}
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
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed thread log: %s.\n", l.Type))
		}
	}
}

func (parser *Parser) OnUnknownLog(callback func(message string)) {
	parser.onUnknownLog = callback
}
