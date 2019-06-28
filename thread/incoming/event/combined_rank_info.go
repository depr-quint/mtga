package event

type CombinedRankInfo struct {
	PlayerId                    string  `json:"playerId"`
	ConstructedSeasonOrdinal    int     `json:"constructedSeasonOrdinal"`
	ConstructedClass            string  `json:"constructedClass"`
	ConstructedLevel            int     `json:"constructedLevel"`
	ConstructedStep             int     `json:"constructedStep"`
	ConstructedMatchesWon       int     `json:"constructedMatchesWon"`
	ConstructedMatchesLost      int     `json:"constructedMatchesLost"`
	ConstructedMatchesDrawn     int     `json:"constructedMatchesDrawn"`
	LimitedSeasonOrdinal        int     `json:"limitedSeasonOrdinal"`
	LimitedClass                string  `json:"limitedClass"`
	LimitedLevel                int     `json:"limitedLevel"`
	LimitedStep                 int     `json:"limitedStep"`
	LimitedMatchesWon           int     `json:"limitedMatchesWon"`
	LimitedMatchesLost          int     `json:"limitedMatchesLost"`
	LimitedMatchesDrawn         int     `json:"limitedMatchesDrawn"`
	ConstructedPercentile       float64 `json:"constructedPercentile"`
	ConstructedLeaderboardPlace int     `json:"constructedLeaderboardPlace"`
	LimitedPercentile           float64 `json:"limitedPercentile"`
	LimitedLeaderboardPlace     int     `json:"limitedLeaderboardPlace"`
}
