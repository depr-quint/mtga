package main

import (
	"flag"
	"fmt"
	"github.com/di-wu/mtga"
	"github.com/di-wu/mtga/thread/outgoing"
	"log"
	"os"
	"path/filepath"
)

var filePath string

func main() {
	flag.StringVar(&filePath, "file", filepath.Join(os.Getenv("APPDATA"), "..", "LocalLow", "Wizards Of The Coast", "MTGA", "output_log.txt"), "Location to the MTGAs log file.")
	flag.Parse()

	parser := mtga.Parser{}
	parser.OnAuthenticate(func(auth outgoing.Authenticate) {
		fmt.Println("Authenticated!")
	})

	t, err := mtga.NewTail(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for l := range t.Logs() {
		parser.ParseRawLog(l)
	}
}
