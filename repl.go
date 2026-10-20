package main

import (
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Fields(text)
	wordsToLower := []string{}
	for _, word := range words {
		wordToLower := strings.ToLower(word)
		wordsToLower = append(wordsToLower, wordToLower)
	}
	return wordsToLower
}
