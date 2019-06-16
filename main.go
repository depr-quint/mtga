package main

import (
	"bufio"
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

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var loggedIn bool
	var accountName, accountNumber string

	var method, body string

	var line int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line++

		// Skip all lines before the account has been logged in successfully.
		switch text := scanner.Text(); {
		case strings.HasPrefix(text, "[Accounts - Startup] Successfully logged in to account:"):
			account := strings.Split(text, ": ")[1]
			parts := strings.Split(account, "#")
			if len(parts) != 2 {
				fmt.Println("Could not parse account name.")
				return
			}
			accountName = parts[0]
			accountNumber = parts[1]

			fmt.Printf("Account Name: %s (#%s)\n", accountName, accountNumber)
			loggedIn = true
			break
		case !loggedIn:
			continue // skip

		// Ignore all following lines:
		case strings.TrimSpace(text) == "" || // e.g. whitespace, new lines, etc.
			strings.HasPrefix(text, "(Filename:") || // e.g. (Filename: [Path] Line: [Int])
			strings.HasPrefix(text, "[UnityCrossThreadLogger]") || // e.g. [UnityCrossThreadLogger] [Date]
			strings.HasPrefix(text, "loading bundle") || // e.g. loading bundle [Path]
			strings.HasPrefix(text, "Timer"): // e.g.Timer PregameSequence start [Date Hr]
			continue

		// In/Out
		case strings.HasPrefix(text, "==>"):
			method = strings.TrimSuffix(strings.TrimPrefix(text, "==> "), ":")
		case strings.HasPrefix(text, "<=="):
			method = strings.TrimPrefix(text, "<== ")
		case strings.HasPrefix(text, "(-1)"):
			method = strings.TrimPrefix(text, "(-1) ")
			body += "["
		case strings.HasPrefix(text, "[Client GRE]"):
			method = strings.TrimPrefix(text, "[Client GRE]")

		case method != "":
			body += strings.TrimSpace(text)
			if text == "}" || text == "]" {
				// fmt.Println(method)
				// fmt.Println(body)
				method, body = "", ""
			}

		default:
			// fmt.Printf("line %d: %s\n", line, text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
