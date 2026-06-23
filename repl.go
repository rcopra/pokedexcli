package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rcopra/pokedexcli/internal/pokeapi"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	pokeClient := pokeapi.NewClient(5 * time.Minute)
	cfg := &config{
		client:  pokeClient,
		pokedex: make(map[string]pokeapi.Pokemon),
	}
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			break
		}
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			fmt.Println("Empty input, try again")
			continue
		}
		commandName := words[0]
		cmd, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(cfg, words[1:]...); err != nil {
			fmt.Println(err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
