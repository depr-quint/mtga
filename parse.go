package main

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

var MTGATime = "1/2/2006 3:04:05 PM"

func parseBody(body []string) {
	if len(body) <= 1 {
		return
	}

	switch first, remaining := body[0], body[1:]; {
	case strings.HasPrefix(first, "[UnityCrossThreadLogger]"):
		date := strings.TrimPrefix(first, "[UnityCrossThreadLogger]")
		t, err := time.Parse(MTGATime, date)
		if err != nil {
			if !strings.HasPrefix(first, "[UnityCrossThreadLogger]Received unhandled GREMessageType") {
				log.Printf("Unparsed thread log: %s\n%s\n", first, remaining)
			}
			return
		}
		parseTreadLogger(t, remaining)
	case strings.HasPrefix(first, "[Client GRE]"):
		parts := strings.Split(strings.TrimPrefix(first, "[Client GRE]"), ":")
		t, err := time.Parse(MTGATime, strings.Join(parts[0:3], ":"))
		if err != nil {
			return
		}
		parseClient(t, clientMethod(strings.TrimSpace(parts[4])), remaining)
	case strings.HasPrefix(first, "[Get SKUs]") ||
		strings.HasPrefix(first, "[Store - Auth - Edit Payment]"):
		// ignore for the time being
	case strings.HasPrefix(first, "WARNING") ||
		strings.HasPrefix(first, "BIError") ||
		strings.HasPrefix(first, "Unloading"):
		// ignore warnings/unloading

	default:
		log.Fatalf("Unparsed log: %s.\n%s\n", first, remaining)
	}
}

func parseClient(t time.Time, method clientMethod, body []string) {
	if body[0] != "{" {
		log.Printf("Unparsed client message: %s\n%s\n", method, body)
		return
	}

	raw := []byte(strings.Join(body, " "))
	switch method {
	case authRequest, connRequest, clientToGre, clientToGreUi:
		var request request
		if err := json.Unmarshal(raw, &request); err != nil {
			log.Fatal(err)
		}
	case authResponse:
		var response response
		if err := json.Unmarshal(raw, &response); err != nil {
			log.Fatal(err)
		}
	case greToClient:
		var event messageEvent
		if err := json.Unmarshal(raw, &event); err != nil {
			log.Fatal(err)
		}
	case roomState:
		var event stateEvent
		if err := json.Unmarshal(raw, &event); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatalf("Unparsed client log: %s.\n%s\n", method, body)
	}
}

type request struct {
	RequestId   int    `json:"requestId"`
	MessageType string `json:"clientToMatchServiceMessageType"`
	Payload     string `json:"payload"`
}

type response struct {
	TransactionId string               `json:"transactionId"`
	RequestId     int                  `json:"requestId"`
	AuthResponse  authenticateResponse `json:"authenticateResponse"`
}

type authenticateResponse struct {
	ClientId   string `json:"clientId"`
	SessionId  string `json:"sessionId"`
	ScreenName string `json:"screenName"`
}

type messageEvent struct {
	TransactionId string      `json:"transactionId"`
	Timestamp     string      `json:"timestamp"`
	Messages      interface{} `json:"greToClientEvent"`
}

type stateEvent struct {
	TransactionId string      `json:"transactionId"`
	Timestamp     string      `json:"timestamp"`
	RoomState     interface{} `json:"matchGameRoomStateChangedEvent"`
}

type clientMethod string

const (
	greToClient   clientMethod = "GreToClientEvent"
	authRequest   clientMethod = "ClientToMatchServiceMessageType_AuthenticateRequest"
	authResponse  clientMethod = "AuthenticateResponse"
	connRequest   clientMethod = "ClientToMatchServiceMessageType_ClientToMatchDoorConnectRequest"
	roomState     clientMethod = "MatchGameRoomStateChangedEvent"
	clientToGre   clientMethod = "ClientToMatchServiceMessageType_ClientToGREMessage"
	clientToGreUi clientMethod = "ClientToMatchServiceMessageType_ClientToGREUIMessage"
)
