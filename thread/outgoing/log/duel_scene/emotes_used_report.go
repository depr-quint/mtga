package duel_scene

type EmotesUsedReport struct {
	MatchId  string  `json:"matchId"`
	Emotes   []Emote `json:"emotes"`
	PlayerId string  `json:"playerId"`
}

type Emote struct {
	EmoteName    string `json:"emoteName"`
	EmoteMessage string `json:"emoteMessage"`
	EmoteCount   int    `json:"emoteCount"`
}
