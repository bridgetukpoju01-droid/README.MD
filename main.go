package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func countAndPrint(text string, fromFile bool) {
	characterCount := make(map[rune]int)

	for _, ch := range text {
		if ch == '\n' || ch == '\r' {
			continue
		}
		lowercaseCh := unicode.ToLower(ch)
		characterCount[lowercaseCh]++
	}

	availabe := make(map[rune]bool)

	for _, ch := range text {
		if ch == '\n' || ch == '\r' {
			continue
		}

		lowercaseCh := unicode.ToLower(ch)

		if availabe[lowercaseCh] {
			continue
		}

		availabe[lowercaseCh] = true

		displayCh := lowercaseCh
		if fromFile {
			displayCh = ch
		}

		fmt.Printf("%c > %d\n", displayCh, characterCount[lowercaseCh])
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <string|filename>")
		return
	}

	userInput := os.Args[1]

	fileContents, err := os.ReadFile(userInput)

	if err != nil {
		countAndPrint(userInput, false)
		return
	}

	text := strings.TrimRight(string(fileContents), "\n\r")

	if strings.TrimSpace(text) == "" {
		fmt.Println("An Empty Data File")
		return
	}

	countAndPrint(text, true)
}
