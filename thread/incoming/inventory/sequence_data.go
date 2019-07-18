package inventory

type SequenceData struct {
	PlayerId          string `json:"playerId"`
	DailySequence     int    `json:"dailySequence"`
	WeeklySequence    int    `json:"weeklySequence"`
	DailyLastAwarded  string `json:"dailyLastAwarded"`
	WeeklyLastAwarded string `json:"weeklyLastAwarded"`
}
