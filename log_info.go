package main

import (
	"encoding/json"
	"log"
	"time"
)

type logInfo struct {
	MessageName   logInfoMessage         `json:"messageName"`
	HumanContext  string                 `json:"humanContext"`
	Payload       map[string]interface{} `json:"payloadObject"`
	TransactionId string                 `json:"transactionId"`
}

func (info logInfo) parse() {
	payload, err := json.Marshal(info.Payload)
	if err != nil {
		log.Fatal(err)
	}

	switch info.MessageName {
	case sceneChangeMsg:
		var c sceneChange
		if err := json.Unmarshal(payload, &c); err != nil {
			log.Fatal(err)
		}
	case performanceReportMsg:
		var r performanceReport
		if err := json.Unmarshal(payload, &r); err != nil {
			log.Fatal(err)
		}
		if len(r.TotalMemoryOddities) != 0 {
			log.Printf("Unparsed fields: %s", r.TotalMemoryOddities)
		}
	case bootSequenceReportMsg:
		var r bootSequenceReport
		if err := json.Unmarshal(payload, &r); err != nil {
			log.Fatal(err)
		}
	case userDeviceSpecsMsg:
		var s userDeviceSpecs
		if err := json.Unmarshal(payload, &s); err != nil {
			log.Fatal(err)
		}
	case connectedMsg:
		var c connected
		if err := json.Unmarshal(payload, &c); err != nil {
			log.Fatal(err)
		}
	case inventoryReportMsg:
		var r inventoryReport
		if err := json.Unmarshal(payload, &r); err != nil {
			log.Fatal(err)
		}
	case purchaseFunnelMsg:
		var f purchaseFunnel
		if err := json.Unmarshal(payload, &f); err != nil {
			log.Fatal(err)
		}
	case gameStartMsg:
		var s gameStart
		if err := json.Unmarshal(payload, &s); err != nil {
			log.Fatal(err)
		}
	case pregameSequenceReportMsg:
		var r pregameSequenceReport
		if err := json.Unmarshal(payload, &r); err != nil {
			log.Fatal(err)
		}
	case gameStopMsg:
		var s gameStop
		if err := json.Unmarshal(payload, &s); err != nil {
			log.Fatal(err)
		}
		if len(s.MulliganedHands) != 0 {
			log.Printf("Unparsed fields: %s", s.MulliganedHands)
		}
	case endOfMatchReportMsg:
		var r endOfMatchReport
		if err := json.Unmarshal(payload, &r); err != nil {
			log.Fatal(err)
		}
	case emotesUsedReportMsg:
		var r emotesUsedReport
		if err := json.Unmarshal(payload, &r); err != nil {
			log.Fatal(err)
		}
		if len(r.Emotes) != 0 {
			log.Printf("Unparsed fields: %s", r.Emotes)
		}

	default:
		log.Fatalf("Unparsed outgoing thread log: %s.\n%s\n", info.MessageName, payload)
	}
}

type logInfoMessage string

const (
	sceneChangeMsg           logInfoMessage = "Client.SceneChange"
	performanceReportMsg     logInfoMessage = "Client.PerformanceReport"
	bootSequenceReportMsg    logInfoMessage = "Client.BootSequenceReport"
	userDeviceSpecsMsg       logInfoMessage = "Client.UserDeviceSpecs"
	connectedMsg             logInfoMessage = "Client.Connected"
	inventoryReportMsg       logInfoMessage = "Client.InventoryReport"
	purchaseFunnelMsg        logInfoMessage = "Client.PurchaseFunnel"
	gameStartMsg             logInfoMessage = "DuelScene.GameStart"
	pregameSequenceReportMsg logInfoMessage = "Client.PregameSequenceReport"
	gameStopMsg              logInfoMessage = "DuelScene.GameStop"
	endOfMatchReportMsg      logInfoMessage = "DuelScene.EndOfMatchReport"
	emotesUsedReportMsg      logInfoMessage = "DuelScene.EmotesUsedReport"
)

// Client changed scenes.
type sceneChange struct {
	From      string    `json:"fromSceneName"`
	To        string    `json:"toSceneName"`
	Timestamp time.Time `json:"timestamp"`
	Duration  string    `json:"duration"`
	Initiator string    `json:"initiator"`
	Context   string    `json:"context"`
	PlayerId  string    `json:"playerId"`
}

// Session Performance Analysis.
type performanceReport struct {
	FrameRateAverage            int     `json:"FrameRateAverage"`
	FrameRateMinimum            int     `json:"FrameRateMinimum"`
	FrameRateDeviation          int     `json:"FrameRateDeviation"`
	FirstFrameRateFrame         int     `json:"FirstFrameRateFrame"`
	LastFrameRateFrame          int     `json:"LastFrameRateFrame"`
	FrameRateOddities           []frame `json:"FrameRateOddities"`
	TotalMemoryAverage          int     `json:"TotalMemoryAverage"`
	TotalMemoryMaximum          int     `json:"TotalMemoryMaximum"`
	TotalMemoryDeviation        int     `json:"TotalMemoryDeviation"`
	FirstTotalMemorySampleFrame int     `json:"FirstTotalMemorySampleFrame"`
	LastTotalMemorySampleFrame  int     `json:"LastTotalMemorySampleFrame"`
	TotalMemoryOddities         []interface{} // no example found in log, just an []
	Bookmarks                   []bookmark       `json:"Bookmarks"`
	QualitySettings             []qualitySetting `json:"QualitySettings"`
	CpuBenchmark                benchmark        `json:"CpuBenchmark"`
	GpuBenchmark                benchmark        `json:"GpuBenchmark"`
	PlayerId                    string           `json:"playerId"`
}

type frame struct {
	Frame     int `json:"Frame"`
	Value     int `json:"Value"`
	Deviation int `json:"Deviation"`
}

type bookmark struct {
	Frame int    `json:"Frame"`
	Value string `json:"Value"`
}

type qualitySetting struct {
	Name  string `json:"Name"`
	Value int    `json:"Value"`
}

type benchmark struct {
	Median        float64 `json:"Median"`
	Score         float64 `json:"Score"`
	AdjustedScore float64 `json:"AdjustedScore"`
}

// Duration of the application launch process including granular durations of notable events within.
// Durations are in seconds.
type bootSequenceReport struct {
	BuildVersion               string `json:"buildVersion"`
	EndpointHash               string `json:"EndpointHash"`
	HashExistingManifest       string `json:"SimpleAssetDownloader_HashExistingManifest"`
	LoadManifest               string `json:"SimpleAssetDownloader_LoadManifest"`
	DeleteOldOrModifiedFiles   string `json:"SimpleAssetDownloader_DeleteOldOrModifiedFiles"`
	DetermineAssetsToDownload  string `json:"SimpleAssetDownloader_DetermineAssetsToDownload"`
	SimpleAssetDownloader      string `json:"SimpleAssetDownloader"`
	AssetBundles               string `json:"MdnAssetLibraryLoad_AssetBundles"`
	SetupPayloadTypeDictionary string `json:"MdnAssetLibrary_SetupPayloadTypeDictionary"`
	MdnAssetLibraryLoad        string `json:"MdnAssetLibraryLoad"`
	FindObjects                string `json:"MDNGlobals_FindObjects"`
	CreateObjectPool           string `json:"MDNGlobals_CreateObjectPool"`
	CardDatabase               string `json:"MDNGlobals_CardDatabase"`
	MDNGlobals                 string `json:"MDNGlobals"`
	SceneLoad                  string `json:"SceneLoad"`
	ObjectPooling              string `json:"ObjectPooling"`
	SplashScreen               string `json:"SplashScreen"`
	BootSequenceDuration       string `json:"BootSequenceDuration"`
	PlayerId                   string `json:"playerId"`
}

// User Device Specs
type userDeviceSpecs struct {
	GraphicsDeviceName          string              `json:"graphicsDeviceName"`
	GraphicsDeviceType          string              `json:"graphicsDeviceType"`
	GraphicsDeviceVendor        string              `json:"graphicsDeviceVendor"`
	GraphicsDeviceVersion       string              `json:"graphicsDeviceVersion"`
	GraphicsMemorySize          int                 `json:"graphicsMemorySize"`
	GraphicsMultiThreaded       bool                `json:"graphicsMultiThreaded"`
	GraphicsShaderLevel         int                 `json:"graphicsShaderLevel"`
	DeviceUniqueIdentifier      string              `json:"deviceUniqueIdentifier"`
	DeviceModel                 string              `json:"deviceModel"`
	DeviceType                  string              `json:"deviceType"`
	OperatingSystem             string              `json:"operatingSystem"`
	OperatingSystemFamily       string              `json:"operatingSystemFamily"`
	ProcessorCount              int                 `json:"processorCount"`
	ProcessorFrequency          int                 `json:"processorFrequency"`
	ProcessorType               string              `json:"processorType"`
	SystemMemorySize            int                 `json:"systemMemorySize"`
	MaxTextureSize              int                 `json:"maxTextureSize"`
	IsWindowed                  bool                `json:"isWindowed"`
	GameResolution              resolution          `json:"gameResolution"`
	MonitorResolution           resolution          `json:"monitorResolution"`
	MonitorSupportedResolutions []monitorResolution `json:"monitorSupportedResolutions"`
	PlayerId                    string              `json:"playerId"`
}

type resolution struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type monitorResolution struct {
	resolution
	ValidForWindow     bool `json:"validForWindow"`
	ValidForFullscreen bool `json:"validForFullscreens"`
}

// Client connected to FrontDoor
type connected struct {
	PlayerId      string   `json:"playerId"`
	ScreenName    string   `json:"screenName"`
	Timestamp     int      `json:"timestamp"`
	ClientVersion string   `json:"clientVersion"`
	Settings      settings `json:"settings"`
}

type settings struct {
	Gameplay gameplay `json:"gameplay"`
	Audio    audio    `json:"audio"`
	Language language `json:"language"`
}

type gameplay struct {
	DisableEmotes                bool `json:"disableEmotes"`
	EvergreenKeywordReminders    bool `json:"evergreenKeywordReminders"`
	AutoTap                      bool `json:"autoTap"`
	AutoOrderTriggeredAbilities  bool `json:"autoOrderTriggeredAbilities"`
	AutoChooseReplacementEffects bool `json:"autoChooseReplacementEffects"`
	ShowPhaseLadder              bool `json:"showPhaseLadder"`
	AllPlayModesToggle           bool `json:"allPlayModesToggle"`
}

type audio struct {
	Master           float64 `json:"master"`
	Music            float64 `json:"music"`
	Effects          float64 `json:"effects"`
	Voice            float64 `json:"voice"`
	Ambience         float64 `json:"ambience"`
	PlayInBackground bool    `json:"playInBackground"`
}

type language struct {
	Language string `json:"language"`
}

// Summary of inventory.
type inventoryReport struct {
	Gold             int    `json:"gold"`
	Gems             int    `json:"gems"`
	WildcardCommon   int    `json:"wcCommon"`
	WildcardUncommon int    `json:"wcUncommon"`
	WildcardRare     int    `json:"wcRare"`
	WildcardMythic   int    `json:"wcMythic"`
	DraftTokens      int    `json:"draftTokens"`
	SealedTokens     int    `json:"sealedTokens"`
	PlayerId         string `json:"playerId"`
}

// Updated available store SKUs.
type purchaseFunnel struct {
	Context  string `json:"context"`
	PlayerId string `json:"playerId"`
}

// General information about a game that has started within a match.
type gameStart struct {
	PlayerId   string `json:"playerId"`
	SeatId     int    `json:"seatId"`
	TeamId     int    `json:"teamId"`
	GameNumber int    `json:"gameNumber"`
	MatchId    string `json:"matchId"`
	EventId    string `json:"eventId"`
}

// Duration of the matchmaking process including granular durations of notable events within. Durations are in seconds.
type pregameSequenceReport struct {
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

// General information about a game that has ended within a match
type gameStop struct {
	PlayerId                  string        `json:"playerId"`
	SeatId                    int           `json:"seatId"`
	TeamId                    int           `json:"teamId"`
	GameNumber                int           `json:"gameNumber"`
	MatchId                   string        `json:"matchId"`
	EventId                   string        `json:"eventId"`
	StartingTeamId            int           `json:"startingTeamId"`
	WinningTeamId             int           `json:"winningTeamId"`
	WinningReason             string        `json:"winningReason"`
	MulliganedHands           []interface{} `json:"mulliganedHands"`
	TurnCount                 int           `json:"turnCount"`
	TurnCountInFullControl    int           `json:"turnCountInFullControl"`
	SecondsCount              int           `json:"secondsCount"`
	SecondsCountInFullControl int           `json:"secondsCountInFullControl"`
	RopeShownCount            int           `json:"ropeShownCount"`
	RopeExpiredCount          int           `json:"ropeExpiredCount"`
}

// End of match report.
type endOfMatchReport struct {
	MatchId                               string   `json:"matchId"`
	MaxCreatures                          int      `json:"maxCreatures"`
	MaxLands                              int      `json:"maxLands"`
	MaxArtifactsAndEnchantments           int      `json:"maxArtifactsAndEnchantments"`
	LongestPassPriorityWaitTimeInSeconds  string   `json:"longestPassPriorityWaitTimeInSeconds"`
	ShortestPassPriorityWaitTimeInSeconds string   `json:"shortestPassPriorityWaitTimeInSeconds"`
	AveragePassPriorityWaitTimeInSeconds  float64  `json:"averagePassPriorityWaitTimeInSeconds"`
	ReceivedPriorityCount                 int      `json:"receivedPriorityCount"`
	PassedPriorityCount                   int      `json:"passedPriorityCount"`
	RespondedToPriorityCount              int      `json:"respondedToPriorityCount"`
	SpellsCastWithAutoPayCount            int      `json:"spellsCastWithAutoPayCount"`
	SpellsCastWithManualManaCount         int      `json:"spellsCastWithManualManaCount"`
	SpellsCastWithMixedPayManaCount       int      `json:"spellsCastWithMixedPayManaCount"`
	AbilityUseByGrpId                     string   `json:"abilityUseByGrpId"`
	AbilityCanceledByGrpId                string   `json:"abilityCanceledByGrpId"`
	AverageActionsByLocalPhaseStep        string   `json:"averageActionsByLocalPhaseStep"`
	AverageActionsByOpponentPhaseStep     string   `json:"averageActionsByOpponentPhaseStep"`
	InteractionCount                      []string `json:"interactionCount"`
	PlayerId                              string   `json:"playerId"`
}

// A tally of emotes used by a player during a match.
type emotesUsedReport struct {
	MatchId  string        `json:"matchId"`
	Emotes   []interface{} `json:"emotes"`
	PlayerId string        `json:"playerId"`
}
