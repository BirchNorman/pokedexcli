package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func commandExit(cfg *config, params ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, params ...string) error {
	var usage strings.Builder
	for _, c := range cfg.commands {
		fmt.Fprintf(&usage, "\n%s: %s", c.name, c.description)
	}
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n%s", usage.String())
	return nil
}

func commandMap(cfg *config, params ...string) error {
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

func commandMapb(cfg *config, params ...string) error {
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

func commandExplore(cfg *config, params ...string) error {
	if len(params) != 1 {
		return errors.New("You must provide a location name.")
	}

	area_name := params[0]

	PokemonList, err := cfg.pokeapiClient.ListPokemon(area_name)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Printf("\nExploring %s...\nFound Pokemon:", area_name)
	for _, Pokemon := range PokemonList.PokemonEncountersField {
		fmt.Printf("\n- %s", Pokemon.Pokemon.Name)

	}
	return nil
}

func commandCatch(cfg *config, params ...string) error {
	if len(params) != 1 {
		return errors.New("You must provide a Pokemon name.")
	}

	PokemonName := params[0]

	PokemonInfo, err := cfg.pokeapiClient.GetPokemonInfo(PokemonName)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Printf("\nThrowing a Pokeball at %s...", PokemonName)
	randomNumber := rand.Intn(PokemonInfo.BaseExperience)

	if randomNumber < 40 {
		fmt.Printf("\n%s was caught!", PokemonName)
		cfg.pokedex[PokemonName] = PokemonInfo
	} else {
		fmt.Printf("\n%s escaped!", PokemonName)
	}
	return nil
}

func commandInspect(cfg *config, params ...string) error {
	if len(params) != 1 {
		return errors.New("You must provide a Pokemon name.")
	}

	PokemonName := params[0]

	PokemonInfo, ok := cfg.pokedex[PokemonName]
	if ok != true {
		fmt.Print("\nYou have not caught that pokemon.")
		return nil
	}

	fmt.Printf("Name: %s\n", PokemonInfo.Name)
	fmt.Printf("Height: %d\n", PokemonInfo.Height)
	fmt.Printf("Weight: %d\n", PokemonInfo.Weight)

	fmt.Print("Stats:\n")
	for _, Stat := range PokemonInfo.Stats {
		fmt.Printf("  -%s: %d\n", Stat.Stat.Name, Stat.BaseStat)
	}

	fmt.Print("Types:\n")
	for _, Type := range PokemonInfo.Types {
		fmt.Printf("  - %s\n", Type.Type.Name)
	}

	return nil
}
