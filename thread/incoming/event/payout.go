package event

type Payout struct {
	EventPayouts []string `json:"eventPayouts"`
	SeasonPayout string   `json:"seasonPayout"`
}
