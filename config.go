package main

import "github.com/rcopra/pokedexcli/internal/pokeapi"

type config struct {
	next    *string
	client  *pokeapi.Client
	pokedex map[string]pokeapi.Pokemon
}
