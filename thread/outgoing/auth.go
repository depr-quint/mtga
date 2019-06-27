package outgoing

type Authenticate struct {
	ClientVersion string `json:"clientVersion"`
	Ticket        string `json:"ticket"`
}
