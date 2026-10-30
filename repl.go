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
	pokedex       map[string]pokeapi.Pokemon
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
		params := []string{}
		if len(cleanedInput) > 1 {
			params = cleanedInput[1:]
		}
		register(cfg, command, params...)
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

func register(cfg *config, command string, params ...string) {
	unknown := true
	for _, c := range cfg.commands {
		if c.name == command {
			c.callback(cfg, params...)
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
	callback    func(*config, ...string) error
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
			description: "Display the names of all Pokemon found in that area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Adds the Pokemon to the user's Pokedex.",
			callback:    commandCatch,
		},
		"inpect": {
			name:        "inspect",
			description: "Prints the name, height, weight, stats and type(s) of the Pokemon.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints a list of all teh names of the Pokemon the user has caught.",
			callback:    commandPokedex,
		},
	}
}
