package mtga

import (
	"github.com/di-wu/mtga/thread/unhandled"
	"testing"
)

func TestUnhandled(t *testing.T) {
	l := RawLog{
		Body: []string{
			`[UnityCrossThreadLogger]Received unhandled GREMessageType: GREMessageType_DieRollResultsResp`,
			`{ "type": "GREMessageType_DieRollResultsResp", "systemSeatIds": [ 1, 2 ], "msgId": 3, "dieRollResultsResp": { "playerDieRolls": [ { "systemSeatId": 1, "rollValue": 5 }, { "systemSeatId": 2, "rollValue": 4 } ] } }`,
		},
	}
	var callback bool
	parser := Parser{}
	parser.OnDieRollResults(func(results unhandled.DieRollResults) {
		callback = true
		if len(results.PlayerDieRolls) != 2 {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}
