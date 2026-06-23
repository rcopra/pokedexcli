package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type PokemonPage struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func (c *Client) ListPokemon(url string) (PokemonPage, error) {
	cacheData, ok := c.cache.Get(url)
	if ok {
		var pokemon PokemonPage
		if err := json.Unmarshal(cacheData, &pokemon); err != nil {
			return PokemonPage{}, err
		}
		return pokemon, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return PokemonPage{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonPage{}, err
	}
	c.cache.Add(url, data)
	var pokemon PokemonPage
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return PokemonPage{}, err
	}
	return pokemon, nil
}
