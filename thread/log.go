package thread

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// ThreadLog
//
//	// FORMAT
//	[UnityCrossThreadLogger] ...
//	==> Method(Id):
//	{Json}
type Log struct {
	Type   LogType
	Method LogMethod
	Id     int
	Json   []byte
}

func NewLog(heading string, body []string) Log {
	first, remaining := body[0], body[1:]

	var typ LogType
	if strings.HasPrefix(first, Outgoing) {
		typ = Outgoing
	} else if strings.HasPrefix(first, Incoming) {
		typ = Incoming
	}

	switch typ {
	case Outgoing, Incoming:
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

		return Log{typ, method, id, raw}
	}
	return Log{}
}

type LogMethod string

const (
	LogInfoMethod LogMethod = "Log.Info"
)

type LogType string

const (
	Outgoing = "==>"
	Incoming = "<=="
)
