package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	// You will implement your solution here
	// The current return value is just a hint, you can replace it
	words := strings.Fields(s)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
}

func main() {
	fmt.Println(WordCount("The quick brown fox jumped over the lazy dog that liked the word the."))
}
