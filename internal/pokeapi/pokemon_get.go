package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func (c *Client) GetPokemon(url string) (Pokemon, error) {
	cacheData, ok := c.cache.Get(url)
	if ok {
		var pokemon Pokemon
		if err := json.Unmarshal(cacheData, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, data)
	var pokemon Pokemon
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
