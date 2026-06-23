package main

import (
	"fmt"
	"math/rand"

	"github.com/rcopra/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("must include the name of the pokemon you want to catch")
	}

	pokemonName := args[0]
	url := pokeapi.BasePokemonURL + pokemonName

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	res, err := cfg.client.GetPokemon(url)
	if err != nil {
		return err
	}

	roll := rand.Intn(500)
	chance := res.BaseExperience
	if roll > chance {
		fmt.Printf("%v was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = res
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}

	return nil
}
