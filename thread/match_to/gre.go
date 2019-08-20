package match_to

import "github.com/di-wu/mtga/thread"

type GreToClientEvent struct {
	GreToClientMessages []Response `json:"greToClientMessages"`
}

type Response struct {
	Type                    thread.LogMethod     `json:"type"`
	SystemSeatIds           []int                `json:"systemSeatIds"`
	MsgId                   int                  `json:"msgId"`
	GameStateId             int                  `json:"gameStateId"`
	ConnectResp             *ConnectResp         `json:"connectResp"`
	DieRollResultsResp      *DieRollResultsResp  `json:"dieRollResultsResp"`
	GameStateMessage        *GameStateMessage    `json:"gameStateMessage"`
	GetSettingsResp         *Settings            `json:"getSettingsResp"`
	SetSettingsResp         *Settings            `json:"setSettingsResp"`
	Prompt                  *Prompt              `json:"prompt"`
	MulliganReq             *MulliganReq         `json:"mulliganReq"`
	NonDecisionPlayerPrompt *Prompt              `json:"nonDecisionPlayerPrompt"`
	TimerStateMessage       *TimerStateMessage   `json:"timerStateMessage"`
	UiMessage               *UiMessage           `json:"uiMessage"`
	ActionsAvailableReq     *ActionsAvailableReq `json:"actionsAvailableReq"`
	DeclareAttackersReq     *DeclareAttackersReq `json:"declareAttackersReq"`
	SubmitTargetsResp       *Submit              `json:"submitTargetsResp"`
	SubmitAttackersResp     *Submit              `json:"submitAttackersResp"`
	SelectTargetsReq        *Select              `json:"selectTargetsReq"`
	AllowCancel             *string              `json:"allow_cancel"`
	AllowUndo               *bool                `json:"allow_undo"`
	IntermissionReq         *IntermissionReq     `json:"intermissionReq"`
}

type ConnectResp struct {
	Status      string      `json:"status"`
	MajorVer    int         `json:"majorVer"`
	RevisionVer int         `json:"revisionVer"`
	BuildVer    int         `json:"buildVer"`
	ProtoVer    string      `json:"protoVer"`
	Settings    Settings    `json:"settings"`
	DeckMessage DeckMessage `json:"deckMessage"`
}

type Settings struct {
	Stops                                  []Stop `json:"stops"`
	AutoPassOption                         string `json:"autoPassOption"`
	GraveyardOrder                         string `json:"graveyardOrder"`
	ManaSelectionType                      string `json:"manaSelectionType"`
	DefaultAutoPassOption                  string `json:"defaultAutoPassOption"`
	SmartStopsSetting                      string `json:"smartStopsSetting"`
	AutoTapStopsSetting                    string `json:"autoTapStopsSetting"`
	AutoOptionalPaymentCancellationSetting string `json:"autoOptionalPaymentCancellationSetting"`
	TransientStops                         []Stop `json:"transientStops"`
}

type Stop struct {
	StopType  string `json:"stopType"`
	AppliesTo string `json:"appliesTo"`
	Status    string `json:"status"`
}

type DeckMessage struct {
	DeckCards      []int `json:"deckCards"`
	SideboardCards []int `json:"sideboardCards"`
}

type DieRollResultsResp struct {
	PlayerDieRolls []DieRoll `json:"playerDieRolls"`
}

type DieRoll struct {
	SystemSeatId int `json:"systemSeatId"`
	RollValue    int `json:"rollValue"`
}

type GameStateMessage struct {
	Type                   string       `json:"type"`
	GameStateId            int          `json:"gameStateId"`
	GameInfo               GameInfo     `json:"gameInfo"`
	Teams                  []Team       `json:"teams"`
	Players                []Player     `json:"players"`
	TurnInfo               TurnInfo     `json:"turnInfo"`
	Zones                  []Zone       `json:"zones"`
	GameObjects            []GameObject `json:"gameObjects"`
	Annotations            []Annotation `json:"annotations"`
	DiffDeletedInstanceIds []int        `json:"diffDeletedInstanceIds"`
	PrevGameStateId        int          `json:"prevGameStateId"`
	Timers                 []Timer      `json:"timers"`
	Update                 string       `json:"update"`
	Actions                []Action     `json:"actions"`
}

type GameInfo struct {
	MatchID            string             `json:"matchID"`
	GameNumber         int                `json:"gameNumber"`
	Stage              string             `json:"stage"`
	Type               string             `json:"type"`
	Variant            string             `json:"variant"`
	MatchState         string             `json:"matchState"`
	MatchWinCondition  string             `json:"matchWinCondition"`
	MaxTimeoutCount    int                `json:"maxTimeoutCount"`
	MaxPipCount        int                `json:"maxPipCount"`
	TimeoutDurationSec int                `json:"timeoutDurationSec"`
	SuperFormat        string             `json:"superFormat"`
	MulliganType       string             `json:"mulliganType"`
	DeckConstraintInfo DeckConstraintInfo `json:"deckConstraintInfo"`
}

type DeckConstraintInfo struct {
	MinDeckSize      int `json:"minDeckSize"`
	MaxDeckSize      int `json:"maxDeckSize"`
	MaxSideboardSize int `json:"maxSideboardSize"`
}

type Team struct {
	Id        int   `json:"id"`
	PlayerIds []int `json:"playerIds"`
}

type Player struct {
	LifeTotal          int    `json:"lifeTotal"`
	SystemSeatNumber   int    `json:"systemSeatNumber"`
	MaxHandSize        int    `json:"maxHandSize"`
	TeamId             int    `json:"teamId"`
	TimerIds           []int  `json:"timerIds"`
	ControllerSeatId   int    `json:"controllerSeatId"`
	ControllerType     string `json:"controllerType"`
	PendingMessageType string `json:"pendingMessageType"`
}

type TurnInfo struct {
	ActivePlayer   int `json:"activePlayer"`
	DecisionPlayer int `json:"decisionPlayer"`
}

type Zone struct {
	ZoneId            int    `json:"zoneId"`
	Type              string `json:"type"`
	Visibility        string `json:"visibility"`
	OwnerSeatId       int    `json:"ownerSeatId"`
	ObjectInstanceIds []int  `json:"objectInstanceIds"`
	Viewers           []int  `json:"viewers"`
}

type GameObject struct {
	InstanceId       int      `json:"instanceId"`
	GrpId            int      `json:"grpId"`
	Type             string   `json:"type"`
	ZoneId           int      `json:"zoneId"`
	Visibility       string   `json:"visibility"`
	OwnerSeatId      int      `json:"ownerSeatId"`
	ControllerSeatId int      `json:"controllerSeatId"`
	SuperTypes       []string `json:"superTypes"`
	CardTypes        []string `json:"cardTypes"`
	Subtypes         []string `json:"subtypes"`
	Color            []string `json:"color"`
	Power            Value    `json:"power"`
	Toughness        Value    `json:"toughness"`
	Viewers          []int    `json:"viewers"`
	Name             int      `json:"name"`
	Abilities        []int    `json:"abilities"`
	OverlayGrpId     int      `json:"overlayGrpId"`
}

type Value struct {
	Value int `json:"value"`
}

type Annotation struct {
	Id          int      `json:"id"`
	AffectorId  int      `json:"affectorId"`
	AffectedIds []int    `json:"affectedIds"`
	Type        []string `json:"type"`
}

type Timer struct {
	TimerId             int    `json:"timerId"`
	Type                string `json:"type"`
	DurationSec         int    `json:"durationSec"`
	Running             bool   `json:"running"`
	Behavior            string `json:"behavior"`
	WarningThresholdSec int    `json:"warningThresholdSec"`
	ElapsedMs           int    `json:"elapsedMs"`
}

type ActionSeat struct {
	SeatId int    `json:"seatId"`
	Action Action `json:"action"`
}

type Action struct {
	ActionType string     `json:"actionType"`
	GrpId      int        `json:"grpId"`
	InstanceId int        `json:"instanceId"`
	Grouping   string     `json:"grouping"`
	ManaCost   []ManaCost `json:"manaCost"`
	ShouldStop bool       `json:"shouldStop"`
}

type ManaCost struct {
	Color  []string `json:"color"`
	Count  int      `json:"count"`
	CostId int      `json:"costId"`
}

type Prompt struct {
	PromptId   int          `json:"promptId"`
	Parameters []Parameters `json:"parameters"`
}

type Parameters struct {
	ParameterName string    `json:"parameterName"`
	Type          string    `json:"type"`
	NumberValue   int       `json:"numberValue"`
	Reference     Reference `json:"reference"`
}

type Reference struct {
	Type string `json:"type"`
	Id   int    `json:"id"`
}

type MulliganReq struct {
	MulliganType string `json:"mulligan_type"`
}

type TimerStateMessage struct {
	SeatId int     `json:"seatId"`
	Timers []Timer `json:"timers"`
}

type UiMessage struct {
	SeatIds []int   `json:"seatIds"`
	OnHover OnHover `json:"onHover"`
}

type OnHover struct {
	ObjectId int `json:"objectId"`
}

type ActionsAvailableReq struct {
	Actions []Action `json:"actions"`
}

type DeclareAttackersReq struct {
	Attackers []Attacker `json:"attackers"`
}

type Attacker struct {
	AttackerInstanceId      int                `json:"attackerInstanceId"`
	LegalDamageRecipients   []DamageRecipients `json:"legalDamageRecipients"`
	SelectedDamageRecipient DamageRecipients   `json:"selectedDamageRecipient"`
}

type DamageRecipients struct {
	Type               string `json:"type"`
	PlayerSystemSeatId int    `json:"playerSystemSeatId"`
}

type Submit struct {
	Result string `json:"result"`
}

type Select struct {
	Targets  []Target `json:"targets"`
	SourceId int      `json:"source_id"`
}

type Target struct {
	TargetIdx  int            `json:"target_idx"`
	Targets    []TargetTarget `json:"targets"`
	MinTargets int            `json:"minTargets"`
	MaxTargets int            `json:"maxTargets"`
	Prompt     Prompt         `json:"prompt"`
}

type TargetTarget struct {
	TargetInstanceId int    `json:"targetInstanceId"`
	LegalAction      string `json:"legal_action"`
	Highlight        int    `json:"highlight"`
}

type IntermissionReq struct {
	Options            []Option `json:"options"`
	IntermissionPrompt Prompt   `json:"intermissionPrompt"`
	GameResultType     string   `json:"gameResultType"`
	WinningTeamId      int      `json:"winningTeamId"`
	Result             Result   `json:"result"`
}

type Option struct {
	OptionPrompt Prompt `json:"optionPrompt"`
	ResponseType string `json:"responseType"`
}

type Result struct {
	Scope         string `json:"scope"`
	Result        string `json:"result"`
	WinningTeamId int    `json:"winningTeamId"`
	Reason        string `json:"reason"`
}
