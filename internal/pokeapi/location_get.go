package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationArea(url string) (LocationArea, error) {
	cacheData, ok := c.cache.Get(url)
	if ok {
		var location LocationArea
		if err := json.Unmarshal(cacheData, &location); err != nil {
			return LocationArea{}, err
		}
		return location, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(url, data)
	var location LocationArea
	if err := json.Unmarshal(data, &location); err != nil {
		return LocationArea{}, err
	}
	return location, nil
}
