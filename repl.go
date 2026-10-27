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
		cleanedInput := cleanInput(userInput)
		command := cleanedInput[0]
		parameter := "" // list
		if len(cleanedInput) > 1 {
			parameter = cleanedInput[1]
		}
		register(command, parameter, cfg)
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

func register(command string, parameter string, cfg *config) {
	unknown := true
	for _, c := range cfg.commands {
		if c.name == command {
			c.callback(cfg, parameter)
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
	callback    func(*config, string) error
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
		"explore": {
			name:        "explore",
			description: "Display the names of all Pokémon found in that area.",
			callback:    commandExplore,
		},
	}
}
