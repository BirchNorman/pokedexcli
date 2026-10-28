package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/BirchNorman/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration, interval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(interval),
	}
}

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreas, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	val, ok := c.cache.Get(url)

	data := []byte{}

	if ok == true {
		data = val
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreas{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreas{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreas{}, err
		}
		c.cache.Add(url, data)
	}

	LocationAreasList := LocationAreas{}
	err := json.Unmarshal(data, &LocationAreasList)
	if err != nil {
		return LocationAreas{}, err
	}

	return LocationAreasList, nil
}

func (c *Client) ListPokemon(area_name string) (PokemonEncountersType, error) {
	url := baseURL + "/location-area/" + area_name

	val, ok := c.cache.Get(url)

	data := []byte{}

	if ok == true {
		data = val
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return PokemonEncountersType{}, err
		}
		res, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonEncountersType{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return PokemonEncountersType{}, err
		}
		c.cache.Add(url, data)
	}

	PokemonList := PokemonEncountersType{}
	err := json.Unmarshal(data, &PokemonList)
	if err != nil {
		return PokemonEncountersType{}, err
	}

	fmt.Printf("\nExploring %s...\nFound Pokemon:", area_name)
	return PokemonList, nil
}
