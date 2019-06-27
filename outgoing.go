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
	case payEntryType:
		var e payEntry
		if err := json.Unmarshal(params, &e); err != nil {
			log.Fatal(err)
		}
	case draftType:
		var d draft
		if err := json.Unmarshal(params, &d); err != nil {
			log.Fatal(err)
		}
	case completeDraftType:
		var c completeDraft
		if err := json.Unmarshal(params, &c); err != nil {
			log.Fatal(err)
		}
	case draftStatusType:
		var s draftStatus
		if err := json.Unmarshal(params, &s); err != nil {
			log.Fatal(err)
		}
	case draftMakePickType:
		var p draftMakePick
		if err := json.Unmarshal(params, &p); err != nil {
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
	case claimPrizeType:
		var c claimPrize
		if err := json.Unmarshal(params, &c); err != nil {
			log.Fatal(err)
		}
	case crackBoostersType:
		var c crackBoosters
		if err := json.Unmarshal(params, &c); err != nil {
			log.Fatal(err)
		}
	case purchaseProductType:
		var p purchaseProduct
		if err := json.Unmarshal(params, &p); err != nil {
			log.Fatal(err)
		}
	case updateDeckType:
		var u updateDeck
		if err := json.Unmarshal(params, &u); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("Unparsed outgoing thread log: %s(%s).\n%s\n", l.Method, l.Id, body)
	}
}

type outgoingThreadLogType string

const (
	authenticateType    outgoingThreadLogType = "Authenticate"
	productCatalogType  outgoingThreadLogType = "PlayerInventory.GetProductCatalog"
	crackBoostersType   outgoingThreadLogType = "PlayerInventory.CrackBoostersV3"
	trackDetailType     outgoingThreadLogType = "Quest.GetTrackDetail"
	updateDeckType      outgoingThreadLogType = "Deck.UpdateDeckV3"
	purchaseProductType outgoingThreadLogType = "Mercantile.PurchaseProduct"

	playerCourseType  outgoingThreadLogType = "Event.GetPlayerCourseV2"
	deckSubmitType    outgoingThreadLogType = "Event.DeckSubmitV3"
	joinQueueType     outgoingThreadLogType = "Event.JoinQueue"
	dropType          outgoingThreadLogType = "Event.Drop"
	joinType          outgoingThreadLogType = "Event.Join"
	payEntryType      outgoingThreadLogType = "Event.PayEntry"
	draftType         outgoingThreadLogType = "Event.Draft"
	completeDraftType outgoingThreadLogType = "Event.CompleteDraft"
	claimPrizeType    outgoingThreadLogType = "Event.ClaimPrize"

	draftStatusType   outgoingThreadLogType = "Draft.DraftStatus"
	draftMakePickType outgoingThreadLogType = "Draft.MakePick"

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

type payEntry struct {
	EventName    string `json:"eventName"`
	CurrencyType string `json:"currencyType"`
}

type draft struct {
	EventName string `json:"eventName"`
}

type draftStatus struct {
	DraftId string `json:"draftId"`
}

type draftMakePick struct {
	DraftId    string `json:"draftId"`
	CardId     string `json:"cardId"`
	PackNumber string `json:"packNumber"`
	PickNumber string `json:"pickNumber"`
}

type completeDraft struct {
	EventName string `json:"eventName"`
}

type claimPrize struct {
	EventName string `json:"eventName"`
}

type crackBoosters struct {
	CollationId string `json:"collationId"`
	Count       string `json:"count"`
}

type purchaseProduct struct {
	ProductId string `json:"productId"`
	Quantity  string `json:"quantity"`
}

type updateDeck struct {
	Deck string `json:"deck"`
}
