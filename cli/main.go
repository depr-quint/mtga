package main

import (
	"flag"
	"fmt"
	"github.com/di-wu/mtga/thread/outgoing/log/client"
	panic "log"
	"os"
	"path/filepath"

	"github.com/di-wu/mtga"
	"github.com/di-wu/mtga/thread/outgoing"
)

var filePath string

func main() {
	flag.StringVar(&filePath, "file", filepath.Join(os.Getenv("APPDATA"), "..", "LocalLow", "Wizards Of The Coast", "MTGA", "output_log.txt"), "Location to the MTGAs log file.")
	flag.Parse()

	parser := mtga.Parser{}
	parser.OnAuthenticate(func(auth outgoing.Authenticate) {
		fmt.Println("Authenticated!")
	})
	parser.OnInventoryReport(func(report client.InventoryReport) {
		fmt.Printf("Gold: %d, Gems %d\n", report.Gold, report.Gems)
	})

	t, err := mtga.NewTail(filePath)
	if err != nil {
		panic.Fatal(err)
	}

	for l := range t.Logs() {
		parser.ParseRawLog(l)
	}
}
