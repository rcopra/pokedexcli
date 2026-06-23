package main

import (
	"fmt"

	"github.com/rcopra/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("must include a name or id to explore")
	}
	url := pokeapi.BaseLocationAreaURL + args[0]
	fmt.Printf("Exploring %s...\n", args[0])
	res, err := cfg.client.GetLocationArea(url)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range res.PokemonEncounters {
		fmt.Printf(" - %v\n", encounter.Pokemon.Name)
	}
	return nil
}
