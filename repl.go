package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) error {
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
		register(command, cfg)
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

func register(command string, cfg *config) {
	unknown := true
	for _, c := range cfg.commands {
		if c.name == command {
			c.callback(cfg)
			unknown = false
			break
		}
	}
	if unknown == true {
		fmt.Println("Unknown command")
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	var usage strings.Builder
	for _, c := range cfg.commands {
		usage.WriteString("\n" + c.name + ": " + c.description)
	}
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n%s", usage.String())
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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

type config struct {
	commands map[string]cliCommand
}
