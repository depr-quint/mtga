package match_to

import "github.com/di-wu/mtga/thread"

const (
	AuthenticateResponseMethod           thread.LogMethod = "AuthenticateResponse"
	GreToClientEventMethod               thread.LogMethod = "GreToClientEvent"
	MatchGameRoomStateChangedEventMethod thread.LogMethod = "MatchGameRoomStateChangedEvent"

	GreConnectRespMethod            thread.LogMethod = "GREMessageType_ConnectResp"
	GreDieRollResultsRespMethod     thread.LogMethod = "GREMessageType_DieRollResultsResp"
	GreGameStateMessageMethodMethod thread.LogMethod = "GREMessageType_GameStateMessage"
	GreQueuedGameStateMessageMethod thread.LogMethod = "GREMessageType_QueuedGameStateMessage"
	GreGetSettingsRespMethod        thread.LogMethod = "GREMessageType_GetSettingsResp"
	GreSetSettingsRespMethod        thread.LogMethod = "GREMessageType_SetSettingsResp"
	GrePromptReqMethod              thread.LogMethod = "GREMessageType_PromptReq"
	GreMulliganReqMethod            thread.LogMethod = "GREMessageType_MulliganReq"
	GreTimerStateMessageMethod      thread.LogMethod = "GREMessageType_TimerStateMessage"
	GreUIMessageMethod              thread.LogMethod = "GREMessageType_UIMessage"
	GreActionsAvailableReqMethod    thread.LogMethod = "GREMessageType_ActionsAvailableReq"
	GreDeclareAttackersReMethod     thread.LogMethod = "GREMessageType_DeclareAttackersReq"
	GreSubmitAttackersRespMethod    thread.LogMethod = "GREMessageType_SubmitAttackersResp"
	GreSubmitTargetsRespMethod      thread.LogMethod = "GREMessageType_SubmitTargetsResp"
	GreSelectTargetsReqMethod       thread.LogMethod = "GREMessageType_SelectTargetsReq"
	GreIntermissionReqMethod        thread.LogMethod = "GREMessageType_IntermissionReq"
)
