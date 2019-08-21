package mtga

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/match_to"
	"github.com/di-wu/mtga/thread/unhandled"
)

// Unhandled is a structure that holds the parser's unhandled callbacks.
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
				MsgID              int                      `json:"msgId"`
				DieRollResultsResp unhandled.DieRollResults `json:"dieRollResultsResp"`
			}{}
			err := json.Unmarshal(l.Raw, &results)
			if err != nil {
				log.Fatalln(err)
			}
			parser.onDieRollResults(results.DieRollResultsResp)
		}
	case unhandled.SubmitTargetsRespMethod:
		if parser.onSubmitTargetsResp != nil {
			var resp match_to.Response
			err := json.Unmarshal(l.Raw, &resp)
			if err != nil {
				log.Fatalln(err)
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
				log.Fatalln(err)
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

// OnDieRollResults attaches the given callback, which will be called on getting the die roll results.
func (unhandled *Unhandled) OnDieRollResults(callback func(results unhandled.DieRollResults)) {
	unhandled.onDieRollResults = callback
}

// OnSubmitTargetsResponse attaches the given callback, which will be called on submitting an targets response.
func (unhandled *Unhandled) OnSubmitTargetsResponse(callback func(resp match_to.Submit)) {
	unhandled.onSubmitTargetsResp = callback
}

// OnSubmitAttackersResponse attaches the given callback, which will be called on submitting an attackers response.
func (unhandled *Unhandled) OnSubmitAttackersResponse(callback func(prompt, nonDecision match_to.Prompt, submit match_to.Submit)) {
	unhandled.onSubmitAttackersResp = callback
}
