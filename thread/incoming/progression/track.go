package progression

import "github.com/di-wu/mtga/thread/incoming"

type Track struct {
	Name             string       `json:"name"`
	TrackLevels      []TrackLevel `json:"trackLevels"`
	TrackRewardTiers [][]Reward   `json:"trackRewardTiers"`
	RewardWeb        RewardWeb    `json:"rewardWeb"`
	Enabled          bool         `json:"enabled"`
}

type TrackLevel struct {
	XpToComplete  int  `json:"xpToComplete"`
	IsPageStarter bool `json:"isPageStarter"`
	IsTentpole    bool `json:"isTentpole"`
}

type Reward struct {
	Chest        incoming.Description `json:"chest"`
	OrbsRewarded int                  `json:"orbsRewarded"`
}

type RewardWeb struct {
	AllNodes        []Node `json:"allNodes"`
	TopLevelNodeIds []int  `json:"topLevelNodeIds"`
	Enabled         bool   `json:"enabled"`
}

type Node struct {
	Id                int                  `json:"id"`
	UnlockQuestMetric interface{}          `json:"unlockQuestMetric"`
	UnlockMetricCount int                  `json:"unlockMetricCount"`
	Chest             incoming.Description `json:"chest"`
	UpgradePacket     UpgradePacket        `json:"upgradePacket"`
	ChildIds          []int                `json:"childIds"`
}

type UpgradePacket struct {
	TargetDeckDescription string `json:"targetDeckDescription"`
	CardsAdded            []int  `json:"cardsAdded"`
}
