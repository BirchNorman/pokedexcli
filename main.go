package main

import (
	"time"

	"github.com/BirchNorman/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		commands:      getCommands(),
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}

// next:     "https://pokeapi.co/api/v2/location-area/1",
// previous: "",
