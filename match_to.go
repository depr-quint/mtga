package mtga

import (
	"encoding/json"
	"fmt"
	panic "log"

	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/match_to"
)

type MatchTo struct {
	onAuthenticateResponse           func(response match_to.AuthenticateResponse)
	onGreToClientEvent               func(gre match_to.GreToClientEvent)
	onMatchGameRoomStateChangedEvent func(change match_to.RoomStateChange)

	onGreConnectResp            func(resp match_to.ConnectResp)
	onGreDieRollResultsResp     func(resp match_to.DieRollResultsResp)
	onGreGameStateMessage       func(msg match_to.GameStateMessage)
	onGreQueuedGameStateMessage func(msg match_to.GameStateMessage)
	onGreGetSettingsResp        func(resp match_to.Settings)
	onGreSetSettingsResp        func(resp match_to.Settings)
	onGrePromptReq              func(req match_to.Prompt)
	onGreMulliganReq            func(prompt, nonDecision match_to.Prompt, req match_to.MulliganReq)
	onGreTimerStateMessage      func(msg match_to.TimerStateMessage)
	onGreUIMessage              func(msg match_to.UiMessage)
	onGreActionsAvailableReq    func(prompt match_to.Prompt, req match_to.ActionsAvailableReq)
	onGreDeclareAttackersReq    func(prompt match_to.Prompt, req match_to.DeclareAttackersReq)
	onGreSubmitTargetsResp      func(submit match_to.Submit)
	onGreSubmitAttackersResp    func(prompt, nonDecision match_to.Prompt, submit match_to.Submit)
	onGreSelectTargetsReq       func(prompt, nonDecision match_to.Prompt, targets match_to.Select, allowCancel string, allowUndo bool)
	onGreIntermissionReq        func(req match_to.IntermissionReq)
}

func (parser *Parser) parseMatchToThreadLog(l thread.Log) {
	switch l.Method {
	case match_to.AuthenticateResponseMethod:
		if parser.onAuthenticateResponse != nil {
			var resp match_to.AuthenticateResponse
			err := json.Unmarshal(l.Raw, &resp)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onAuthenticateResponse(resp)
		}
	case match_to.GreToClientEventMethod:
		var gre match_to.GreToClientEvent
		err := json.Unmarshal(l.Raw, &gre)
		if err != nil {
			panic.Fatalln(err)
		}
		if parser.onGreToClientEvent != nil {
			parser.onGreToClientEvent(gre)
		}
		for _, resp := range gre.GreToClientMessages {
			parser.parseGreResponse(resp)
		}
	case match_to.MatchGameRoomStateChangedEventMethod:
		var change match_to.RoomStateChange
		err := json.Unmarshal(l.Raw, &change)
		if err != nil {
			panic.Fatalln(err)
		}
		if parser.onMatchGameRoomStateChangedEvent != nil {
			parser.onMatchGameRoomStateChangedEvent(change)
		}
	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed match to log: %s.\n%s", l.Method, l.Raw))
		}
	}
}

func (to *MatchTo) OnAuthenticateResponse(callback func(response match_to.AuthenticateResponse)) {
	to.onAuthenticateResponse = callback
}

func (to *MatchTo) OnGreToClientEvent(callback func(gre match_to.GreToClientEvent)) {
	to.onGreToClientEvent = callback
}

func (parser *Parser) parseGreResponse(resp match_to.Response) {
	switch resp.Type {
	case match_to.GreConnectRespMethod:
		if parser.onGreConnectResp != nil && resp.ConnectResp != nil {
			parser.onGreConnectResp(*resp.ConnectResp)
		}
	case match_to.GreDieRollResultsRespMethod:
		if parser.onGreDieRollResultsResp != nil && resp.DieRollResultsResp != nil {
			parser.onGreDieRollResultsResp(*resp.DieRollResultsResp)
		}
	case match_to.GreGameStateMessageMethodMethod:
		if parser.onGreGameStateMessage != nil && resp.GameStateMessage != nil {
			parser.onGreGameStateMessage(*resp.GameStateMessage)
		}
	case match_to.GreQueuedGameStateMessageMethod:
		if parser.onGreQueuedGameStateMessage != nil && resp.GameStateMessage != nil {
			parser.onGreQueuedGameStateMessage(*resp.GameStateMessage)
		}
	case match_to.GreGetSettingsRespMethod:
		if parser.onGreGetSettingsResp != nil && resp.GetSettingsResp != nil {
			parser.onGreGetSettingsResp(*resp.GetSettingsResp)
		}
	case match_to.GreSetSettingsRespMethod:
		if parser.onGreSetSettingsResp != nil && resp.SetSettingsResp != nil {
			parser.onGreSetSettingsResp(*resp.SetSettingsResp)
		}
	case match_to.GrePromptReqMethod:
		if parser.onGrePromptReq != nil && resp.Prompt != nil {
			parser.onGrePromptReq(*resp.Prompt)
		}
	case match_to.GreMulliganReqMethod:
		if parser.onGreMulliganReq != nil && resp.Prompt != nil &&
			resp.MulliganReq != nil && resp.NonDecisionPlayerPrompt != nil {
			parser.onGreMulliganReq(*resp.Prompt, *resp.NonDecisionPlayerPrompt, *resp.MulliganReq)
		}
	case match_to.GreTimerStateMessageMethod:
		if parser.onGreTimerStateMessage != nil && resp.TimerStateMessage != nil {
			parser.onGreTimerStateMessage(*resp.TimerStateMessage)
		}
	case match_to.GreUIMessageMethod:
		if parser.onGreUIMessage != nil && resp.UiMessage != nil {
			parser.onGreUIMessage(*resp.UiMessage)
		}
	case match_to.GreActionsAvailableReqMethod:
		if parser.onGreActionsAvailableReq != nil && resp.Prompt != nil && resp.ActionsAvailableReq != nil {
			parser.onGreActionsAvailableReq(*resp.Prompt, *resp.ActionsAvailableReq)
		}
	case match_to.GreDeclareAttackersReMethod:
		if parser.onGreDeclareAttackersReq != nil && resp.Prompt != nil && resp.DeclareAttackersReq != nil {
			parser.onGreDeclareAttackersReq(*resp.Prompt, *resp.DeclareAttackersReq)
		}
	case match_to.GreSubmitAttackersRespMethod:
		if parser.onGreSubmitAttackersResp != nil && resp.Prompt != nil &&
			resp.SubmitAttackersResp != nil && resp.NonDecisionPlayerPrompt != nil {
			parser.onGreSubmitAttackersResp(*resp.Prompt, *resp.NonDecisionPlayerPrompt, *resp.SubmitAttackersResp)
		}
	case match_to.GreSubmitTargetsRespMethod:
		if parser.onGreSubmitTargetsResp != nil && resp.SubmitTargetsResp != nil {
			parser.onGreSubmitTargetsResp(*resp.SubmitTargetsResp)
		}
	case match_to.GreSelectTargetsReqMethod:
		if parser.onGreSelectTargetsReq != nil && resp.Prompt != nil && resp.SelectTargetsReq != nil &&
			resp.NonDecisionPlayerPrompt != nil && resp.AllowCancel != nil && resp.AllowUndo != nil {
			parser.onGreSelectTargetsReq(*resp.Prompt, *resp.NonDecisionPlayerPrompt, *resp.SelectTargetsReq, *resp.AllowCancel, *resp.AllowUndo)
		}
	case match_to.GreIntermissionReqMethod:
		if parser.onGreIntermissionReq != nil && resp.IntermissionReq != nil {
			parser.onGreIntermissionReq(*resp.IntermissionReq)
		}
	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed gre log: %s", resp.Type))
		}
	}
}

func (to *MatchTo) OnGreConnectResp(callback func(resp match_to.ConnectResp)) {
	to.onGreConnectResp = callback
}

func (to *MatchTo) OnGreDieRollResultsResp(callback func(resp match_to.DieRollResultsResp)) {
	to.onGreDieRollResultsResp = callback
}

func (to *MatchTo) OnGreGameStateMessage(callback func(msg match_to.GameStateMessage)) {
	to.onGreGameStateMessage = callback
}

func (to *MatchTo) OnGreQueuedGameStateMessage(callback func(msg match_to.GameStateMessage)) {
	to.onGreQueuedGameStateMessage = callback
}

func (to *MatchTo) OnGreGetSettingsResp(callback func(resp match_to.Settings)) {
	to.onGreGetSettingsResp = callback
}

func (to *MatchTo) OnGreSetSettingsResp(callback func(resp match_to.Settings)) {
	to.onGreSetSettingsResp = callback
}

func (to *MatchTo) OnGrePromptReq(callback func(req match_to.Prompt)) {
	to.onGrePromptReq = callback
}

func (to *MatchTo) OnGreMulliganReq(callback func(prompt, nonDecision match_to.Prompt, req match_to.MulliganReq)) {
	to.onGreMulliganReq = callback
}

func (to *MatchTo) OnGreTimerStateMessage(callback func(msg match_to.TimerStateMessage)) {
	to.onGreTimerStateMessage = callback
}

func (to *MatchTo) OnGreUIMessage(callback func(msg match_to.UiMessage)) {
	to.onGreUIMessage = callback
}

func (to *MatchTo) OnGreActionsAvailableReq(callback func(prompt match_to.Prompt, req match_to.ActionsAvailableReq)) {
	to.onGreActionsAvailableReq = callback
}

func (to *MatchTo) OnGreDeclareAttackersReq(callback func(prompt match_to.Prompt, req match_to.DeclareAttackersReq)) {
	to.onGreDeclareAttackersReq = callback
}

func (to *MatchTo) OnGreSubmitTargetsResp(callback func(submit match_to.Submit)) {
	to.onGreSubmitTargetsResp = callback
}

func (to *MatchTo) OnGreSubmitAttackersResp(callback func(prompt, nonDecision match_to.Prompt, submit match_to.Submit)) {
	to.onGreSubmitAttackersResp = callback
}

func (to *MatchTo) OnGreSelectTargetsReq(callback func(prompt, nonDecision match_to.Prompt, targets match_to.Select, allowCancel string, allowUndo bool)) {
	to.onGreSelectTargetsReq = callback
}

func (to *MatchTo) OnGreIntermissionReq(callback func(req match_to.IntermissionReq)) {
	to.onGreIntermissionReq = callback
}
