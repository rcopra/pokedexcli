package main

import (
	"fmt"

	"github.com/rcopra/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config, args ...string) error {
	url := pokeapi.BaseLocationAreaURL
	if cfg.next != nil {
		url = *cfg.next
	}

	res, err := cfg.client.ListLocationAreas(url)
	if err != nil {
		return err
	}

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}

	cfg.next = res.Next

	return nil
}
