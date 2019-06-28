package front_door

type ConnectionDetails struct {
	SessionId string `json:"sessionId"`
	IsQueued  string `json:"isQueued"`
}
