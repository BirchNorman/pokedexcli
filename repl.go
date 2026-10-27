package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/BirchNorman/pokedexcli/internal/pokeapi"
	"github.com/BirchNorman/pokedexcli/internal/pokecache"
)

type config struct {
	commands      map[string]cliCommand
	pokeapiClient pokeapi.Client
	next          *string
	prev          *string
	pokeCache     pokecache.Cache
}

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
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas.",
			callback:    commandMapb,
		},
	}
}
