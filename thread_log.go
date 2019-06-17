package main

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type threadLog struct {
	Method outgoingThreadLogType  `json:"method"`
	Params map[string]interface{} `json:"params"`
	Id     string                 `json:"id"`
}

func parseTreadLogger(t time.Time, body []string) {
	switch first := body[0]; {

	// outgoing thread logs
	case strings.HasPrefix(first, "==>"):
		parseOutgoing(body[1:])

	// incoming thread logs
	case strings.HasPrefix(first, "<=="):
		parts := strings.Split(strings.TrimSuffix(strings.TrimPrefix(first, "<== "), ")"), "(")
		parseIncoming(parts[0], parts[1], body[1:])

	default:
		// fmt.Println(first)
	}
}

func parseOutgoing(body []string) {
	var l threadLog
	raw := []byte(strings.Join(body, " "))
	if err := json.Unmarshal(raw, &l); err != nil {
		log.Fatal(err)
	}

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

	default:
		log.Fatalf("Unparsed outgoing thread log: %s(%s).\n", l.Method, l.Id)
	}
}

type outgoingThreadLogType string

const (
	authenticateType   outgoingThreadLogType = "Authenticate"
	productCatalogType outgoingThreadLogType = "PlayerInventory.GetProductCatalog"
	trackDetailType    outgoingThreadLogType = "Quest.GetTrackDetail"
	playerCourseType   outgoingThreadLogType = "Event.GetPlayerCourseV2"
	joinQueueType      outgoingThreadLogType = "Event.JoinQueue"

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

type joinQueue struct {
	QueueId string `json:"queueId"`
	Avatar  string `json:"avatar"`
}

func parseIncoming(method, id string, body []string) {
	// TODO
}
