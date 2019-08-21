package mtga

import (
	"github.com/di-wu/mtga/thread/match_to"
	"testing"
)

func TestMatchToAuthenticateResponse(t *testing.T) {
	l := []string{
		`[UnityCrossThreadLogger]8/10/2019 1:00:57 PM: Match to XXX: AuthenticateResponse`,
		`{ "transactionId": "a878bfc6-3310-4274-8627-92dc1c927b23", "requestId": 1, "authenticateResponse": { "clientId": "SJ2QOI43EBCP3DQMHPJG73YLNY", "sessionId": "08aba436-39ac-4880-8d1f-9a5aaaab2dbe", "screenName": "Diwu#93074" } }`,
	}
	var callback bool
	parser := Parser{}
	parser.OnAuthenticateResponse(func(response match_to.AuthenticateResponse) {
		callback = true
		if response.ScreenName != "Diwu#93074" {
			t.Error()
		}
	})
	parser.Parse(l)
	if !callback {
		t.Error()
	}
}
