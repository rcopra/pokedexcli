package main

import "github.com/rcopra/pokedexcli/internal/pokeapi"

// config holds the state shared across REPL commands.
type config struct {
	next    *string
	client  *pokeapi.Client
	pokedex map[string]pokeapi.Pokemon
}
