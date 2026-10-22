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
		fmt.Print("\nPokedex > ")
		userInput := ""
		if scanner.Scan() {
			userInput = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
		command := cleanInput(userInput)[0]
		register(command)
		// fmt.Printf("Your command was: %s\n", cleanedInput)
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

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func register(command string) {
	unknown := true
	commands := getCommands()
	for _, c := range commands {
		if c.name == command {
			c.callback()
			unknown = false
			break
		}
	}
	if unknown == true {
		fmt.Println("Unknown command")
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	commands := getCommands()
	var usage strings.Builder
	for _, c := range commands {
		usage.WriteString("\n" + c.name + ": " + c.description)
	}
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n%s", usage.String())
	return nil
}
