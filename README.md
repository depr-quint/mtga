# MTGA Output Log Parser
This is a tool for tailing and parsing the MTGA output_log.txt file. 

[![GoDoc](https://godoc.org/github.com/di-wu/mtga?status.svg)](https://godoc.org/github.com/di-wu/mtga)
## Getting Started
### Log Location
```gotemplate
filepath.Join(os.Getenv("APPDATA"), "..", "LocalLow", "Wizards Of The Coast", "MTGA", "output_log.txt")
```
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

## MTGA Card API
[Scryfall API](https://scryfall.com/docs/api) |
[Arena Endpoint](https://scryfall.com/docs/api/cards/arena)
