package main

import (
	"log"
	"strings"
	"time"
)

type threadLog struct {
	Method outgoingThreadLogType  `json:"method"`
	Params map[string]interface{} `json:"params"`
	Id     string                 `json:"id"`
}

func parseTreadLogger(t time.Time, body []string) {
	switch first, remaining := body[0], body[1:]; {

	// outgoing thread logs
	case strings.HasPrefix(first, "==>"):
		parseOutgoing(remaining)

	// incoming thread logs
	case strings.HasPrefix(first, "<=="):
		// TODO

	case strings.HasPrefix(first, "(-1)"):
		// TODO

	default:
		log.Fatalf("Unparsed outgoing thread log: %s\n%s\n", first, remaining)
	}
}
