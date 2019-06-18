package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var filePath string

func main() {
	flag.StringVar(&filePath, "file", filepath.Join(os.Getenv("APPDATA"), "..", "LocalLow", "Wizards Of The Coast", "MTGA", "output_log.txt"), "Location to the MTGAs log file.")
	flag.Parse()

	t, err := NewTail(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var loggedIn bool
	var accountName, accountNumber string
	var body []string
	for line := range t.lines {
		switch text := strings.TrimSpace(string(line)); {
		case strings.HasPrefix(text, "Initialize engine version"):
			loggedIn = false

		case strings.HasPrefix(text, "[Accounts - Startup] Successfully logged in to account:"):
			account := strings.Split(text, ": ")[1]
			parts := strings.Split(account, "#")
			if len(parts) != 2 {
				log.Fatalln("Could not parse account name.")
			}
			accountName = parts[0]
			accountNumber = parts[1]

			fmt.Printf("Account Name: %s (#%s)\n", accountName, accountNumber)
			loggedIn = true

		case !loggedIn:
			// skip

		case text == "":
			parseBody(body)
			body = nil

		default:
			body = append(body, text)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
