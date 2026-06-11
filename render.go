package main

import (
	"fmt"
	"strings"
)

func render(input string, lines []string) {
	if len(strings.ReplaceAll(input, "\\n", "")) == 0 {
		fmt.Print(strings.ReplaceAll(input, "\\n", "\n"))
		return
	}
	for _, seg := range strings.Split(input, "\\n") {
		if seg == "" {
			fmt.Println()
			continue
		}
		for row := 0; row < 8; row++ {
			for i := 0; i < len(seg); i++ {
				pos := int(seg[i]-32)*9 + 1 + row
				if pos < len(lines) {
					fmt.Print(lines[pos])
				}
			}
			fmt.Println()
		}
	}
}
