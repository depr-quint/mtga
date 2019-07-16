package inventory

import "github.com/di-wu/mtga/thread/incoming"

type RewardSchedule struct {
	DailyReset    string   `json:"dailyReset"`
	WeeklyReset   string   `json:"weeklyReset"`
	DailyRewards  []Reward `json:"dailyRewards"`
	WeeklyRewards []Reward `json:"weeklyRewards"`
}

type Reward struct {
	Wins             int                  `json:"wins"`
	AwardDescription incoming.Description `json:"awardDescription"`
}
