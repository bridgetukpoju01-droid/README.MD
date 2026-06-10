package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("error reading file")
		return
	}

	text := strings.ToLower(string(data))

	count := make(map[rune]int)

	for _, ch := range text {
		if ch != '\n' {
			count[ch]++
		}
	}

	ex := make(map[rune]bool)

	for _, ch := range text {
		if ch != '\n' && !ex[ch] {
			fmt.Printf("%c -> %d\n", ch, count[ch])
			ex[ch] = true
		}
	}
}
