package mtga

import (
	"github.com/di-wu/mtga/thread"
	"testing"
)

func TestToMatchThreadLog(t *testing.T) {
	l := RawLog{
		Body: []string{
			`[UnityCrossThreadLogger]8/11/2019 9:28:23 AM: XXX to Match: ClientToMatchServiceMessageType_ClientToGREMessage`,
			`{ "requestId": 3, "clientToMatchServiceMessageType": "ClientToMatchServiceMessageType_ClientToGREMessage", "payload": "CBSCAQQKAkgB" }`,
		},
	}
	var callback bool
	parser := Parser{}
	parser.OnToMatchThreadLog(func(method thread.LogMethod, payload string) {
		callback = true
		if method != "ClientToMatchServiceMessageType_ClientToGREMessage" {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}
