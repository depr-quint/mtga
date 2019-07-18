package event

type Payout struct {
	EventPayouts interface{} `json:"eventPayouts"`
	SeasonPayout string      `json:"seasonPayout"`
}
