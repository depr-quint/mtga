package match_to

type RoomStateChange struct {
	GameRoomInfo GameRoomInfo `json:"gameRoomInfo"`
}

type GameRoomInfo struct {
	GameRoomConfig   GameRoomConfig    `json:"gameRoomConfig"`
	StateType        string            `json:"stateType"`
	Players          []PlayerInfo      `json:"players"`
	FinalMatchResult *FinalMatchResult `json:"finalMatchResult"`
}

type GameRoomConfig struct {
	EventId                     string           `json:"eventId"`
	ReservedPlayers             []ReservedPlayer `json:"reservedPlayers"`
	MatchId                     string           `json:"matchId"`
	MatchConfig                 MatchConfig      `json:"matchConfig"`
	GreConfig                   GreConfig        `json:"greConfig"`
	GreHostLoggerLevel          string           `json:"greHostLoggerLevel"`
	JoinRoomTimeoutSecs         int              `json:"joinRoomTimeoutSecs"`
	PlayerDisconnectTimeoutSecs int              `json:"playerDisconnectTimeoutSecs"`
}

type PlayerInfo struct {
	UserId       string `json:"userId"`
	SystemSeatId int    `json:"systemSeatId"`
}

type ReservedPlayer struct {
	UserId         string         `json:"userId"`
	PlayerName     string         `json:"playerName"`
	SystemSeatId   int            `json:"systemSeatId"`
	TeamId         int            `json:"teamId"`
	ConnectionInfo ConnectionInfo `json:"connectionInfo"`
	CourseId       string         `json:"courseId"`
}

type ConnectionInfo struct {
	ConnectionState string `json:"connectionState"`
}

type MatchConfig struct {
}

type GreConfig struct {
	GameStateRedactorConfiguration GameStateRedactorConfiguration `json:"gameStateRedactorConfiguration"`
	ClipsConfiguration             ClipsConfiguration             `json:"clipsConfiguration"`
	CheckpointConfiguration        CheckpointConfiguration        `json:"checkpointConfiguration"`
}

type GameStateRedactorConfiguration struct {
	EnableRedaction bool `json:"enableRedaction"`
	EnableForceDiff bool `json:"enableForceDiff"`
}

type ClipsConfiguration struct {
}

type CheckpointConfiguration struct {
}

type FinalMatchResult struct {
	MatchId              string       `json:"matchId"`
	MatchCompletedReason string       `json:"matchCompletedReason"`
	ResultList           []ResultInfo `json:"resultList"`
}

type ResultInfo struct {
	Scope         string `json:"scope"`
	Result        string `json:"result"`
	WinningTeamId int    `json:"winningTeamId"`
}
