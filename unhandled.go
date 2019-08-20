package mtga

import (
	"encoding/json"
	"fmt"
	"github.com/di-wu/mtga/thread/match_to"
	panic "log"

	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/unhandled"
)

type Unhandled struct {
	onDieRollResults      func(results unhandled.DieRollResults)
	onSubmitTargetsResp   func(resp match_to.Submit)
	onSubmitAttackersResp func(prompt, nonDecision match_to.Prompt, submit match_to.Submit)
}

func (parser *Parser) parseUnhandledThreadLog(l thread.Log) {
	switch l.Method {
	case unhandled.DieRollResultsMethod:
		if parser.onDieRollResults != nil {
			results := struct {
				Type               string                   `json:"type"`
				SystemSeatIds      []int                    `json:"systemSeatIds"`
				MsgId              int                      `json:"msgId"`
				DieRollResultsResp unhandled.DieRollResults `json:"dieRollResultsResp"`
			}{}
			err := json.Unmarshal(l.Raw, &results)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onDieRollResults(results.DieRollResultsResp)
		}
	case unhandled.SubmitTargetsRespMethod:
		if parser.onSubmitTargetsResp != nil {
			var resp match_to.Response
			err := json.Unmarshal(l.Raw, &resp)
			if err != nil {
				panic.Fatalln(err)
			}

			if resp.SubmitTargetsResp != nil {
				parser.onSubmitTargetsResp(*resp.SubmitTargetsResp)
			}
		}
	case unhandled.SubmitAttackersRespMethod:
		if parser.onSubmitAttackersResp != nil {
			var resp match_to.Response
			err := json.Unmarshal(l.Raw, &resp)
			if err != nil {
				panic.Fatalln(err)
			}

			if resp.Prompt != nil && resp.SubmitAttackersResp != nil && resp.NonDecisionPlayerPrompt != nil {
				parser.onSubmitAttackersResp(*resp.Prompt, *resp.NonDecisionPlayerPrompt, *resp.SubmitAttackersResp)
			}
		}
	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed unhandled log: %s.\n%s", l.Method, l.Raw))
		}
	}
}

func (unhandled *Unhandled) OnDieRollResults(callback func(results unhandled.DieRollResults)) {
	unhandled.onDieRollResults = callback
}

func (unhandled *Unhandled) OnSubmitTargetsResp(callback func(resp match_to.Submit)) {
	unhandled.onSubmitTargetsResp = callback
}

func (unhandled *Unhandled) OnSubmitAttackersResp(callback func(prompt, nonDecision match_to.Prompt, submit match_to.Submit)) {
	unhandled.onSubmitAttackersResp = callback
}
