package main

import (
	"fmt"
	"os"
	"strings"
)

func commandExit(cfg *config, parameter string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, parameter string) error {
	var usage strings.Builder
	for _, c := range cfg.commands {
		fmt.Fprintf(&usage, "\n%s: %s", c.name, c.description)
	}
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n%s", usage.String())
	return nil
}

func commandMap(cfg *config, parameter string) error {
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

func commandMapb(cfg *config, parameter string) error {
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

func commandExplore(cfg *config, parameter string) error {
	area_name := parameter // nil

	PokemonList, err := cfg.pokeapiClient.ListPokemon(area_name)
	if err != nil {
		return err
	}

	for _, Pokemon := range PokemonList.PokemonEncountersField {
		fmt.Printf("\n- %s", Pokemon.Pokemon.Name)

	}
	return nil
}
