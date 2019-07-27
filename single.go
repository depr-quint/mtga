package mtga

import (
	"encoding/json"
	"fmt"
	panic "log"
	"regexp"
	"strconv"
	"strings"

	"github.com/di-wu/mtga/thread/single"
)

type Single struct {
	onSkinSeen     func(skins single.Skins)
	onCardNotExist func(card single.NotExist)
	onNullEntity   func(null single.NullEntity)
	onStateChange  func(from, to string)
}

func (parser *Parser) parseSingleTreadLog(log string) {
	switch {
	case strings.HasPrefix(log, "Skins seen:"):
		if parser.onSkinSeen != nil {
			s := strings.Split(strings.TrimPrefix(log, "Skins seen:"), " ")
			skins := make(single.Skins)
			for _, v := range s {
				split := strings.Split(v, "=")
				if len(split) != 2 {
					continue
				}

				id, err := strconv.Atoi(split[0])
				if err != nil {
					panic.Println(err)
				}
				txt := split[1]
				skins[id] = txt
			}
			parser.onSkinSeen(skins)
		}
	case strings.HasPrefix(log, "Card #"):
		if parser.onCardNotExist != nil {
			str := regexp.MustCompile(`#([0-9]*?) \(\"([a-zA-Z ,\'-]*?)\"\) had ParentId #([0-9]*?) `).FindStringSubmatch(log)
			cardId, _ := strconv.Atoi(str[1])
			parentId, _ := strconv.Atoi(str[3])
			parser.onCardNotExist(single.NotExist{
				CardId:   cardId,
				CardName: str[2],
				ParentId: parentId,
			})
		}
	case strings.HasPrefix(log, "NULL entity on"):
		if parser.onNullEntity != nil {
			raw := []byte(strings.TrimPrefix(log, "NULL entity on"))
			var null single.NullEntity
			err := json.Unmarshal(raw, &null)
			if err != nil {
				panic.Fatalln(err)
			}
			parser.onNullEntity(null)
		}
	case strings.HasPrefix(log, "STATE CHANGED"):
		if parser.onStateChange != nil {
			str := strings.Split(strings.TrimPrefix(log, "STATE CHANGED "), " -> ")
			parser.onStateChange(str[0], str[1])
		}
	default:
		if parser.onUnknownLog != nil {
			parser.onUnknownLog(fmt.Sprintf("Unparsed single log: %s", log))
		}
	}
}

func (single *Single) OnSkinsSeen(callback func(skins single.Skins)) {
	single.onSkinSeen = callback
}

func (single *Single) OnCardNotExist(callback func(card single.NotExist)) {
	single.onCardNotExist = callback
}

func (single *Single) OnNullEntity(callback func(null single.NullEntity)) {
	single.onNullEntity = callback
}

func (single *Single) OnStateChange(callback func(from, to string)) {
	single.onStateChange = callback
}
