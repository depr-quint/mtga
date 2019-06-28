package client

type PregameSequenceReport struct {
	BuildVersion              string `json:"buildVersion"`
	SceneLoadMatchScene       string `json:"SceneLoad_MatchScene"`
	MatchSceneManagerInit     string `json:"MatchSceneManager_Init"`
	PreGameSceneLoaded        string `json:"None_to_WaitingForMatch_transition_PreGameScene_Loaded"`
	NoneAlphaIn               string `json:"None_to_WaitingForMatch_transition_AlphaIn"`
	NoneExecuteBlock          string `json:"None_to_WaitingForMatch_transition_ExecuteBlock"`
	WaitingForMatchTransition string `json:"None_to_WaitingForMatch_transition"`
	WaitingForMatch           string `json:"WaitingForMatch"`
	BattlefieldLoaded         string `json:"WaitingForMatch_to_MatchReady_transition_Battlefield_Loaded"`
	GameManagerAwake          string `json:"GameManager_Awake"`
	DuelSceneLoaded           string `json:"WaitingForMatch_to_MatchReady_transition_DuelScene_Loaded"`
	WaitingAlphaIn            string `json:"WaitingForMatch_to_MatchReady_transition_AlphaIn"`
	WaitingExecuteBlock       string `json:"WaitingForMatch_to_MatchReady_transition_ExecuteBlock"`
	MatchReadyTransition      string `json:"WaitingForMatch_to_MatchReady_transition"`
	MatchReady                string `json:"MatchReady"`
	AlphaOut                  string `json:"MatchReady_to_DuelScene_transition_AlphaOut"`
	PreGameSceneUnloaded      string `json:"MatchReady_to_DuelScene_transition_PreGameScene_Unloaded"`
	MatchReadyAlphaIn         string `json:"match_ready_to_duel_scene_transition_alpha_in"`
	DuelSceneTransition       string `json:"MatchReady_to_DuelScene_transition"`
	PregameSequence           string `json:"PregameSequence"`
	BattlefieldSceneName      string `json:"BattlefieldSceneName"`
	PlayerId                  string `json:"playerId"`
	MatchId                   string `json:"matchId"`
}
