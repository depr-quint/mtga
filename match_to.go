package mtga

import (
	"encoding/json"
	"fmt"
	panic "log"

	"github.com/di-wu/mtga/thread"
	"github.com/di-wu/mtga/thread/match_to"
)

type MatchTo struct {
	onAuthenticateResponse func(response match_to.AuthenticateResponse)
	onGreToClientEvent     func(gre match_to.GreToClientEvent)

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
	default:
		// fmt.Println(l.Method, string(l.Raw))
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
	case "GREMessageType_ConnectResp":
		if parser.onGreConnectResp != nil && resp.ConnectResp != nil {
			parser.onGreConnectResp(*resp.ConnectResp)
		}
	case "GREMessageType_DieRollResultsResp":
		if parser.onGreDieRollResultsResp != nil && resp.DieRollResultsResp != nil {
			parser.onGreDieRollResultsResp(*resp.DieRollResultsResp)
		}
	case "GREMessageType_GameStateMessage":
		if parser.onGreGameStateMessage != nil && resp.GameStateMessage != nil {
			parser.onGreGameStateMessage(*resp.GameStateMessage)
		}
	case "GREMessageType_QueuedGameStateMessage":
		if parser.onGreQueuedGameStateMessage != nil && resp.GameStateMessage != nil {
			parser.onGreQueuedGameStateMessage(*resp.GameStateMessage)
		}
	case "GREMessageType_GetSettingsResp":
		if parser.onGreGetSettingsResp != nil && resp.GetSettingsResp != nil {
			parser.onGreGetSettingsResp(*resp.GetSettingsResp)
		}
	case "GREMessageType_SetSettingsResp":
		if parser.onGreSetSettingsResp != nil && resp.SetSettingsResp != nil {
			parser.onGreSetSettingsResp(*resp.SetSettingsResp)
		}
	case "GREMessageType_PromptReq":
		if parser.onGrePromptReq != nil && resp.Prompt != nil {
			parser.onGrePromptReq(*resp.Prompt)
		}
	case "GREMessageType_MulliganReq":
		if parser.onGreMulliganReq != nil && resp.Prompt != nil &&
			resp.MulliganReq != nil && resp.NonDecisionPlayerPrompt != nil {
			parser.onGreMulliganReq(*resp.Prompt, *resp.NonDecisionPlayerPrompt, *resp.MulliganReq)
		}
	case "GREMessageType_TimerStateMessage":
		if parser.onGreTimerStateMessage != nil && resp.TimerStateMessage != nil {
			parser.onGreTimerStateMessage(*resp.TimerStateMessage)
		}
	case "GREMessageType_UIMessage":
		if parser.onGreUIMessage != nil && resp.UiMessage != nil {
			parser.onGreUIMessage(*resp.UiMessage)
		}
	case "GREMessageType_ActionsAvailableReq":
		if parser.onGreActionsAvailableReq != nil && resp.Prompt != nil && resp.ActionsAvailableReq != nil {
			parser.onGreActionsAvailableReq(*resp.Prompt, *resp.ActionsAvailableReq)
		}
	case "GREMessageType_DeclareAttackersReq":
		if parser.onGreDeclareAttackersReq != nil && resp.Prompt != nil && resp.DeclareAttackersReq != nil {
			parser.onGreDeclareAttackersReq(*resp.Prompt, *resp.DeclareAttackersReq)
		}
	case "GREMessageType_SubmitAttackersResp":
		if parser.onGreSubmitAttackersResp != nil && resp.Prompt != nil &&
			resp.SubmitAttackersResp != nil && resp.NonDecisionPlayerPrompt != nil {
			parser.onGreSubmitAttackersResp(*resp.Prompt, *resp.NonDecisionPlayerPrompt, *resp.SubmitAttackersResp)
		}
	case "GREMessageType_SubmitTargetsResp":
		if parser.onGreSubmitTargetsResp != nil && resp.SubmitTargetsResp != nil {
			parser.onGreSubmitTargetsResp(*resp.SubmitTargetsResp)
		}
	case "GREMessageType_SelectTargetsReq":
		if parser.onGreSelectTargetsReq != nil && resp.Prompt != nil && resp.SelectTargetsReq != nil &&
			resp.NonDecisionPlayerPrompt != nil && resp.AllowCancel != nil && resp.AllowUndo != nil {
			parser.onGreSelectTargetsReq(*resp.Prompt, *resp.NonDecisionPlayerPrompt, *resp.SelectTargetsReq, *resp.AllowCancel, *resp.AllowUndo)
		}
	case "GREMessageType_IntermissionReq":
		if parser.onGreIntermissionReq != nil && resp.IntermissionReq != nil {
			parser.onGreIntermissionReq(*resp.IntermissionReq)
		}
	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed gre log: %s", resp.Type))
		}
	}
}
