package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := make(map[string]int)
	for i := 0; i < len(words); i++ {
		count[ words[i] ]++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
