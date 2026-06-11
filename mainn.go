package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		return
	}
	lines, err := loadBanner("standard.txt")
	if err != nil {
		os.Exit(1)
	}
	render(os.Args[1], lines)
}
