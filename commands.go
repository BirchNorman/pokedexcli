package main

import (
	"fmt"
	"os"
	"strings"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	var usage strings.Builder
	for _, c := range cfg.commands {
		fmt.Fprintf(&usage, "\n%s: %s", c.name, c.description)
	}
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n%s", usage.String())
	return nil
}

func commandMap(cfg *config) error {
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.next)
	if err != nil {
		return err
	}
	cfg.next = res.Next
	cfg.prev = res.Prev

	for _, locationArea := range res.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prev == nil {
		fmt.Println("you're on the first page")
	}
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.prev)
	if err != nil {
		return err
	}
	cfg.next = res.Next
	cfg.prev = res.Prev

	for _, locationArea := range res.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}
