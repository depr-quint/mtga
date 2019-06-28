package duel_scene

type GameStart struct {
	PlayerId   string `json:"playerId"`
	SeatId     int    `json:"seatId"`
	TeamId     int    `json:"teamId"`
	GameNumber int    `json:"gameNumber"`
	MatchId    string `json:"matchId"`
	EventId    string `json:"eventId"`
}
