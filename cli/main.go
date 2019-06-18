package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var filePath string

func main() {
	flag.StringVar(&filePath, "file", filepath.Join(os.Getenv("APPDATA"), "..", "LocalLow", "Wizards Of The Coast", "MTGA", "output_log.txt"), "Location to the MTGAs log file.")
	flag.Parse()

	t, err := NewTail(filePath)
	if err != nil {
		log.Fatal(err)
	}

	for line := range t.lines {
		fmt.Println(string(line))
	}

	if t.err != nil {
		log.Fatal(err)
	}
}
