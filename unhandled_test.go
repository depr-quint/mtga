package mtga

import (
	"github.com/di-wu/mtga/thread/match_to"
	"github.com/di-wu/mtga/thread/unhandled"
	"testing"
)

func TestUnhandledDieRollResults(t *testing.T) {
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

func TestUnhandledSubmitTargetsResp(t *testing.T) {
	l := RawLog{
		Body: []string{
			`[UnityCrossThreadLogger]Received unhandled GREMessageType: GREMessageType_SubmitTargetsResp`,
			`{ "type": "GREMessageType_SubmitTargetsResp", "systemSeatIds": [ 2 ], "msgId": 398, "gameStateId": 254, "submitTargetsResp": { "result": "ResultCode_Success" } }`,
		},
	}
	var callback bool
	parser := Parser{}
	parser.OnSubmitTargetsResp(func(resp match_to.Submit) {
		callback = true
		if resp.Result != "ResultCode_Success" {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}

func TestUnhandledSubmitAttackersResp(t *testing.T) {
	l := RawLog{
		Body: []string{
			`[UnityCrossThreadLogger]Received unhandled GREMessageType: GREMessageType_SubmitAttackersResp`,
			`{ "type": "GREMessageType_SubmitAttackersResp", "systemSeatIds": [ 2 ], "msgId": 414, "gameStateId": 263, "prompt": { "promptId": 6 }, "submitAttackersResp": { "result": "ResultCode_Success" }, "nonDecisionPlayerPrompt": { "promptId": 16, "parameters": [ { "parameterName": "PlayerId", "type": "ParameterType_Number", "numberValue": 2 } ] } }`,
		},
	}
	var callback bool
	parser := Parser{}
	parser.OnSubmitAttackersResp(func(prompt, nonDecision match_to.Prompt, submit match_to.Submit) {
		callback = true
		if prompt.PromptId != 6 || submit.Result != "ResultCode_Success" || len(nonDecision.Parameters) != 1 {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}
