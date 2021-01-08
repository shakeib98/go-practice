package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordCount(strings.Split("HHHOOOOO", "")))
}

func wordCount(str []string) map[string]int {
	keyMap := make(map[string]int)
	for i := 0; i < len(str); i++ {
		keyMap[str[i]] = keyMap[str[i]] + 1
	}
	return keyMap
}
