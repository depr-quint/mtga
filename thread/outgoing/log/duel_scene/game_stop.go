package duel_scene

type GameStop struct {
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