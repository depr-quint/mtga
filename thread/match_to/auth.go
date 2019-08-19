package match_to

type AuthenticateResponse struct {
	ClientId   string `json:"clientId"`
	ScreenName string `json:"screenName"`
	SessionId  string `json:"sessionId"`
}
