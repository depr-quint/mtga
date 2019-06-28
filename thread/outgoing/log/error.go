package log

type Err struct {
	Message  string `json:"message"`
	PlayerId string `json:"playerId"`
	MatchId  string `json:"matchId"`
}
