package unhandled

type DieRollResults struct {
	PlayerDieRolls []Roll `json:"playerDieRolls"`
}

type Roll struct {
	SystemSeatId int `json:"systemSeatId"`
	RollValue    int `json:"rollValue"`
}
