package pokeapi

import (
	"encoding/json"
	"github.com/rcopra/pokedexcli/internal/pokecache"
	"io"
	"net/http"
	"time"
)

const BaseLocationAreaURL = "https://pokeapi.co/api/v2/location-area/"

type LocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type LocationAreas struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type Client struct {
	cache pokecache.Cache
}

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		cache: *pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) ListLocationAreas(url string) (LocationAreas, error) {
	/*TODO: Check Cache
	check c.cache.Get(url)
	if found -> unmarshal cached bytes and return
	if not found -> do the HTTP request, then c.cache.Add(url, data) before returning
	*/

	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	var locations LocationAreas
	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationAreas{}, err
	}
	return locations, nil
}

func (c *Client) GetLocationArea(url string) (LocationArea, error) {
	/*TODO: Check Cache
	check c.cache.Get(url)
	if found -> unmarshal cached bytes and return
	if not found -> do the HTTP request, then c.cache.Add(url, data) before returning
	*/
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var location LocationArea
	if err := json.Unmarshal(data, &location); err != nil {
		return LocationArea{}, err
	}
	return location, nil
}
