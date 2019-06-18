package main

import (
	"encoding/json"
	"log"
	"strings"
)

func parseOutgoing(body []string) {
	var l threadLog
	raw := []byte(strings.Join(body, " "))
	if err := json.Unmarshal(raw, &l); err != nil {
		log.Fatal(err)
	}

	// no params to parse further
	if len(l.Params) == 0 {
		return
	}

	params, err := json.Marshal(l.Params)
	if err != nil {
		log.Fatal(err)
	}

	switch l.Method {
	case authenticateType:
		var a authenticate
		if err := json.Unmarshal(params, &a); err != nil {
			log.Fatal(err)
		}

	case logInfoType, logErrorType:
		var l logInfo
		if err := json.Unmarshal(params, &l); err != nil {
			log.Fatal(err)
		}
		l.parse()

	case productCatalogType:
		var c productCatalog
		if err := json.Unmarshal(params, &c); err != nil {
			log.Fatal(err)
		}

	case trackDetailType:
		var d trackDetail
		if err := json.Unmarshal(params, &d); err != nil {
			log.Fatal(err)
		}

	case playerCourseType:
		var c playerCourse
		if err := json.Unmarshal(params, &c); err != nil {
			log.Fatal(err)
		}

	case joinQueueType:
		var q joinQueue
		if err := json.Unmarshal(params, &q); err != nil {
			log.Fatal(err)
		}

	case dropType:
		var d drop
		if err := json.Unmarshal(params, &d); err != nil {
			log.Fatal(err)
		}

	case joinType:
		var j join
		if err := json.Unmarshal(params, &j); err != nil {
			log.Fatal(err)
		}

	case deckSubmitType:
		var s deckSubmit
		if err := json.Unmarshal(params, &s); err != nil {
			log.Fatal(err)
		}

		var d deck
		if err := json.Unmarshal([]byte(s.Deck), &d); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("Unparsed outgoing thread log: %s(%s).\n%s\n", l.Method, l.Id, body)
	}
}

type outgoingThreadLogType string

const (
	authenticateType   outgoingThreadLogType = "Authenticate"
	productCatalogType outgoingThreadLogType = "PlayerInventory.GetProductCatalog"
	trackDetailType    outgoingThreadLogType = "Quest.GetTrackDetail"

	playerCourseType outgoingThreadLogType = "Event.GetPlayerCourseV2"
	deckSubmitType   outgoingThreadLogType = "Event.DeckSubmitV3"
	joinQueueType    outgoingThreadLogType = "Event.JoinQueue"
	dropType         outgoingThreadLogType = "Event.Drop"
	joinType         outgoingThreadLogType = "Event.Join"

	logInfoType  outgoingThreadLogType = "Log.Info"
	logErrorType outgoingThreadLogType = "Log.Error"
)

type authenticate struct {
	ClientVersion string `json:"clientVersion"`
	Ticket        string `json:"ticket"`
}

type productCatalog struct {
	CatalogName string `json:"catalogName"`
}

type trackDetail struct {
	TrackName string `json:"trackName"`
}

type playerCourse struct {
	EventName string `json:"eventName"`
}

type deckSubmit struct {
	EventName string `json:"eventName"`
	Deck      string `json:"deck"`
}

type deck struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Format      string `json:"format"`
	Description string `json:"description"`
	DeckTileId  int    `json:"deckTileId"`
	MainDeck    []int  `json:"mainDeck"`
	CardBack    string `json:"cardBack"`
	CardSkins   []skin `json:"cardSkins"`
	LastUpdated string `json:"lastUpdated"`
}

type skin struct {
	GroupId int    `json:"grpId"`
	Ccv     string `json:"ccv"`
}

type joinQueue struct {
	QueueId string `json:"queueId"`
	Avatar  string `json:"avatar"`
}

type drop struct {
	EventName string `json:"eventName"`
}

type join struct {
	EventName string `json:"eventName"`
}
