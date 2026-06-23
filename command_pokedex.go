package main

import (
	"fmt"
	"sort"
)

func commandPokedex(cfg *config, args ...string) error {
	names := make([]string, 0, len(cfg.pokedex))

	for name := range cfg.pokedex {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Println("-", name)
	}
	return nil
}
