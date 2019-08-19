package minus_one

type MatchCreated struct {
	ControllerFabricUri            string   `json:"controllerFabricUri"`
	MatchEndpointHost              string   `json:"matchEndpointHost"`
	MatchEndpointPort              int      `json:"matchEndpointPort"`
	OpponentScreenName             string   `json:"opponentScreenName"`
	OpponentIsWotc                 bool     `json:"opponentIsWotc"`
	MatchId                        string   `json:"matchId"`
	OpponentRankingClass           string   `json:"opponentRankingClass"`
	OpponentRankingTier            int      `json:"opponentRankingTier"`
	OpponentMythicPercentile       float64  `json:"opponentMythicPercentile"`
	OpponentMythicLeaderboardPlace int      `json:"opponentMythicLeaderboardPlace"`
	EventId                        string   `json:"eventId"`
	OpponentAvatarSelection        string   `json:"opponentAvatarSelection"`
	OpponentCardBackSelection      string   `json:"opponentCardBackSelection"`
	OpponentPetSelection           string   `json:"opponentPetSelection"`
	OpponentPetModSelections       []string `json:"opponentPetModSelections"`
	AvatarSelection                string   `json:"avatarSelection"`
	CardbackSelection              string   `json:"cardbackSelection"`
	PetModSelections               []string `json:"petModSelections"`
	Battlefield                    string   `json:"battlefield"`
}
