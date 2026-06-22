package main

import (
	"bufio"
	"fmt"
	"github.com/rcopra/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	Next     *string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{

		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display location information from the Pokemon world",
			callback:    commandMap,
		},
		"explore": {
			name:        "explore",
			description: "Explore Pokemon in the location using the id or name",
			callback:    commandExplore,
		},
	}
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, c := range getCommands() {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}
	return nil
}

func commandMap(cfg *config, args ...string) error {
	url := pokeapi.BaseLocationAreaURL
	if cfg.Next != nil {
		url = *cfg.Next
	}

	res, err := pokeapi.ListLocationAreas(url)
	if err != nil {
		return err
	}

	for _, area := range res.Results {
		fmt.Println(area.Name)
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous

	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Must include name or id to explore")
	}
	url := pokeapi.BaseLocationAreaURL + args[0]
	fmt.Printf("Exploring %s...\n", args[0])
	res, err := pokeapi.GetLocationArea(url)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range res.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{}
	for {
		fmt.Print("Pokedex > ")
		ok := scanner.Scan()
		if !ok {
			break
		}
		usrInput := scanner.Text()
		words := cleanInput(usrInput)
		if len(words) == 0 {
			fmt.Println("Empty input, try again")
			continue
		}
		commandName := words[0]
		cmd, ok := getCommands()[commandName]
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
