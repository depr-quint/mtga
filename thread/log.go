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

	if strings.HasPrefix(first, Outgoing) {
		method := LogMethod(regexp.MustCompile(`[a-zA-Z0-9.]+`).FindStringSubmatch(first)[0])
		id, _ := strconv.Atoi(regexp.MustCompile(`\(([0-9]+)\)`).FindStringSubmatch(first)[1])
		str := strings.TrimSpace(strings.Join(remaining, " "))

		var raw []byte
		if strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}") {
			var m map[string]interface{}
			err := json.Unmarshal([]byte(str), &m)
			if err != nil {
				log.Fatalln(err)
			}
			raw, err = json.Marshal(m["params"])
			if err != nil {
				log.Fatalln(err)
			}
		}

		return Log{Outgoing, method, id, raw}
	}

	return Log{}
}

type LogMethod string

const (
	AuthenticateMethod   LogMethod = "Authenticate"
	LogInfoMethod        LogMethod = "Log.Info"
	ProductCatalogMethod LogMethod = "PlayerInventory.GetProductCatalog"
	TrackDetailMethod    LogMethod = "Quest.GetTrackDetail"

	PlayerCourseMethod LogMethod = "Event.GetPlayerCourseV2"
	JoinQueueMethod    LogMethod = "Event.JoinQueue"
)

type LogType string

const (
	Outgoing = "==>"
)
