package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		userInput := ""
		if scanner.Scan() {
			userInput = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		cleanedInput := cleanInput(userInput)[0]
		fmt.Printf("Your command was: %s\n", cleanedInput)
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	wordsToLower := []string{}
	for _, word := range words {
		wordToLower := strings.ToLower(word)
		wordsToLower = append(wordsToLower, wordToLower)
	}
	return wordsToLower
}
