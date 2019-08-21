package progression

import "github.com/di-wu/mtga/thread/incoming"

type Track struct {
	Name        string `json:"name"`
	TrackLevels []struct {
		XpToComplete  int  `json:"xpToComplete"`
		IsPageStarter bool `json:"isPageStarter"`
		IsTentpole    bool `json:"isTentpole"`
	} `json:"trackLevels"`
	TrackRewardTiers [][]struct {
		Chest        incoming.Description `json:"chest"`
		OrbsRewarded int                  `json:"orbsRewarded"`
	} `json:"trackRewardTiers"`
	RewardWeb struct {
		AllNodes []struct {
			Id                int                  `json:"id"`
			UnlockQuestMetric interface{}          `json:"unlockQuestMetric"`
			UnlockMetricCount int                  `json:"unlockMetricCount"`
			Chest             incoming.Description `json:"chest"`
			UpgradePacket     struct {
				TargetDeckDescription string `json:"targetDeckDescription"`
				CardsAdded            []int  `json:"cardsAdded"`
			} `json:"upgradePacket"`
			ChildIds []int `json:"childIds"`
		} `json:"allNodes"`
		TopLevelNodeIds []int `json:"topLevelNodeIds"`
		Enabled         bool  `json:"enabled"`
	} `json:"rewardWeb"`
	Enabled bool `json:"enabled"`
}
