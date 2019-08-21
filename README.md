# MTGA Output Log Parser
This is a tool for monitoring and parsing the MTGA output_log.txt file. 

[![CircleCI](https://circleci.com/gh/di-wu/mtga/tree/master.svg?style=svg)](https://circleci.com/gh/di-wu/mtga/tree/master)
[![GoDoc](https://godoc.org/github.com/di-wu/mtga?status.svg)](https://godoc.org/github.com/di-wu/mtga)
## Getting Started
### Log Location
```gotemplate
filepath.Join(os.Getenv("APPDATA"), "..", "LocalLow", "Wizards Of The Coast", "MTGA", "output_log.txt")
```

#### Log Example Files
Jul 25, Update *v0.17*: 
[PDF](http://magic.wizards.com/sites/mtg/files/output_log_arena.pdf) |
[TXT](testdata/output_log_0.17.txt)

### Example
```go
package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"
	
    "github.com/di-wu/mtga"
    "github.com/di-wu/mtga/thread/outgoing"
    "github.com/di-wu/mtga/thread/outgoing/log/client"
)

var filePath string

func main()  {
    flag.StringVar(&filePath, "file", filepath.Join(os.Getenv("APPDATA"), "..", "LocalLow", "Wizards Of The Coast", "MTGA", "output_log.txt"), "Location to the MTGAs log file.")
    flag.Parse()
	
    parser := mtga.Parser{}
    parser.OnAuthenticate(func(auth outgoing.Authenticate) {
        fmt.Println("Authenticated!")
    })
    parser.OnInventoryReport(func(report client.InventoryReport) {
        fmt.Printf("Gold: %d, Gems %d\n", report.Gold, report.Gems)
    })
 
    t, _ := mtga.NewTail(filePath)
 
    for l := range t.Logs() {
        parser.Parse(l)
    }
}
```

### Unimplemented Logs
There are probably still some logs that don't have a callback or are just not that relevant.
```gotemplate
parser.OnUnknownLog(func(message string) {
    log.Println(message)
})
```

## MTGA Card API
[Scryfall API](https://scryfall.com/docs/api) |
[Arena Endpoint](https://scryfall.com/docs/api/cards/arena)

```go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	
	"github.com/di-wu/mtga"
)

type ScryfallResponse struct {
	Object     string `json:"object"`
	TotalCards int    `json:"total_cards"`
	HasMore    bool   `json:"has_more"`
	NextPage   string `json:"next_page"`
	Data       []Card `json:"data"`
}

type Card struct {
	ArenaID int    `json:"arena_id"`
	Name    string `json:"name"`
	Rarity  Rarity `json:"rarity"`
	Set     string `json:"set"`
}

type Rarity string

const (
	C Rarity = "common"
	U Rarity = "uncommon"
	R Rarity = "rare"
	M Rarity = "mythic"
)

// Scryfall returns all the cards from the set matching the given code (e.g. m20)
func Scryfall(code string) (cards []Card) {
	response, err := http.Get(fmt.Sprintf("https://api.scryfall.com/cards/search?q=set:%s", code))
	if err != nil {
		log.Fatalf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	var resp ScryfallResponse
	_ = json.Unmarshal(data, &resp)
	for _, v := range resp.Data {
		cards = append(cards, v)
	}
	if resp.HasMore {
		cards = append(cards, Scryfall(resp.NextPage)...)
	}
	return
}
```
