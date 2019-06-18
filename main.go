package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var filePath string

func main() {
	flag.StringVar(&filePath, "file", filepath.Join(os.Getenv("APPDATA"), "..", "LocalLow", "Wizards Of The Coast", "MTGA", "output_log.txt"), "Location to the MTGAs log file.")
	flag.Parse()

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var loggedIn bool
	var accountName, accountNumber string

	var body []string

	var line int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++

		// Skip all lines before the account has been logged in successfully.
		switch text := strings.TrimSpace(scanner.Text()); {
		case strings.HasPrefix(text, "[Accounts - Startup] Successfully logged in to account:"):
			account := strings.Split(text, ": ")[1]
			parts := strings.Split(account, "#")
			if len(parts) != 2 {
				fmt.Println("Could not parse account name.")
				return
			}
			accountName = parts[0]
			accountNumber = parts[1]

			fmt.Printf("Account Name: %s (#%s)\n", accountName, accountNumber)
			loggedIn = true
			break
		case !loggedIn:
			continue // skip

		case text == "":
			parseBody(body)
			body = nil

		default:
			body = append(body, text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

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
			log.Printf("Unparsed Tread Log: %s\n%s\n", first, remaining)
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
	case strings.HasPrefix(first, "[Get SKUs]"):
		// ignore for the time being
	case strings.HasPrefix(first, "WARNING") ||
		strings.HasPrefix(first, "Unloading"):
		// ignore warnings/unloading

	default:
		log.Fatalf("Unparsed outgoing thread log: %s.\n%s\n", first, remaining)

	}
}

func parseClient(t time.Time, method clientMethod, body []string) {
	if body[0] != "{" {
		log.Printf("Unparsedd Client Message: %s\n%s\n", method, body)
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
		log.Fatalf("Unparsed outgoing thread log: %s.\n%s\n", method, body)
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

type event struct {
	TransactionId string `json:"transactionId"`
	Timestamp     string `json:"timestamp"`
}

type messageEvent struct {
	event
	Messages interface{} `json:"greToClientEvent"`
}

type stateEvent struct {
	event
	RoomState interface{} `json:"matchGameRoomStateChangedEvent"`
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
