package thread

import (
	"bytes"
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ThreadLog
//
//	// FORMAT
//	[UnityCrossThreadLogger] ...
//	==> Method(Id):
//	{Raw}
type Log struct {
	Type   LogType
	Method LogMethod
	Id     int
	Time   time.Time
	Raw    []byte
}

func NewLog(heading string, body []string) Log {
	first, remaining := body[0], body[1:]

	var typ LogType
	if strings.HasPrefix(first, Outgoing) {
		typ = Outgoing
	} else if strings.HasPrefix(first, Incoming) {
		typ = Incoming
	}

	if typ == Outgoing || typ == Incoming {
		t, _ := time.Parse("1/02/2006 03:04:05 PM", heading)
		method := LogMethod(regexp.MustCompile(`[a-zA-Z0-9.]+`).FindStringSubmatch(first)[0])
		id, _ := strconv.Atoi(regexp.MustCompile(`\(([0-9]+)\)`).FindStringSubmatch(first)[1])
		str := strings.TrimSpace(strings.Join(remaining, " "))
		var raw []byte
		if strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}") && typ == Outgoing {
			var m map[string]interface{}
			err := json.Unmarshal([]byte(str), &m)
			if err != nil {
				log.Fatalln(err)
			}
			raw, err = json.Marshal(m["params"])
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			raw = []byte(str)
		}
		return Log{Type: typ, Method: method, Id: id, Time: t, Raw: raw}
	}

	if strings.HasPrefix(heading, "ConnectResp") {
		return Log{Type: ConnectResp, Raw: []byte(strings.Join(append([]string{"{"}, body...), " "))}
	}

	if strings.HasPrefix(heading, "Received unhandled GREMessageType") {
		method := LogMethod(strings.TrimPrefix(heading, "Received unhandled GREMessageType: "))
		return Log{Type: Unhandled, Method: method, Raw: []byte(strings.Join(body, " "))}
	}

	if len(strings.Split(heading, ":")) == 3 {
		t, _ := time.Parse("1/02/2006 03:04:05 PM", heading)
		str := strings.Split(first, " ")
		method, surplus := LogMethod(str[2]), str[3]
		if surplus == "[]" {
			return Log{}
		}
		return Log{Type: MinusOne, Method: method, Time: t, Raw: []byte(strings.Join(append([]string{surplus}, remaining...), " "))}
	}

	if str := strings.Split(heading, ": "); len(str) == 3 {
		t, _ := time.Parse("1/02/2006 03:04:05 PM", str[0])
		var raw []byte
		var m map[string]interface{}
		err := json.Unmarshal([]byte(strings.Join(body, " ")), &m)
		if err != nil {
			log.Fatalln(err)
		}
		match, method := str[1], LogMethod(str[2])
		if strings.HasPrefix(match, MatchTo) {
			bts := []byte(method) // lowercase first letter
			lc := bytes.ToLower([]byte{bts[0]})
			rest := bts[1:]
			raw, err = json.Marshal(m[string(bytes.Join([][]byte{lc, rest}, nil))])
			if err != nil {
				log.Fatalln(err)
			}
			typ = MatchTo
		} else if strings.HasSuffix(match, ToMatch) {
			raw, err = json.Marshal(m["payload"])
			if err != nil {
				log.Fatalln(err)
			}
			typ = ToMatch
		}
		return Log{Type: typ, Method: method, Time: t, Raw: raw}
	}

	return Log{}
}

type LogMethod string

const (
	LogInfoMethod  LogMethod = "Log.Info"
	LogErrorMethod LogMethod = "Log.Error"
)

type LogType string

const (
	Outgoing    = "==>"
	Incoming    = "<=="
	ConnectResp = "ConnectResp"
	Unhandled   = "Received unhandled GREMessageType"
	MinusOne    = "(-1)"
	ToMatch     = "to Match"
	MatchTo     = "Match to"
)
