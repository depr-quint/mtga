package event

import (
	"github.com/di-wu/mtga/thread/incoming"
)

type SeasonRankAndDetail struct {
	CurrentSeason       CurrentSeason `json:"currentSeason"`
	LimitedRankInfo     []RankInfo    `json:"limitedRankInfo"`
	ConstructedRankInfo []RankInfo    `json:"constructedRankInfo"`
}

type CurrentSeason struct {
	SeasonOrdinal            int    `json:"seasonOrdinal"`
	SeasonStartTime          string `json:"seasonStartTime"`
	SeasonEndTime            string `json:"seasonEndTime"`
	SeasonLimitedRewards     Season `json:"seasonLimitedRewards"`
	SeasonConstructedRewards Season `json:"seasonConstructedRewards"`
	MinMatches               int    `json:"minMatches"`
}

type Season struct {
	Bronze   incoming.Description `json:"bronze"`
	Silver   incoming.Description `json:"silver"`
	Gold     incoming.Description `json:"gold"`
	Platinum incoming.Description `json:"platinum"`
	Diamond  incoming.Description `json:"diamond"`
	Mythic   incoming.Description `json:"mythic"`
}

type RankInfo struct {
	RankClass string `json:"rankClass"`
	Level     int    `json:"level"`
	Steps     int    `json:"steps"`
}
