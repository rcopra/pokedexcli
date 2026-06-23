package main

import (
	"bufio"
	"fmt"
	"github.com/rcopra/pokedexcli/internal/pokeapi"
	"math/rand"
	"os"
	"strings"
	"time"
)

type config struct {
	Next     *string
	Previous *string
	pokeapi  *pokeapi.Client
	pokedex  map[string]pokeapi.Pokemon
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
		"catch": {
			name:        "catch",
			description: "Usage: 'throw <pokemon name> - attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a pokemons attributes",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "list all the pokemon you have caught in your pokedex",
			callback:    commandPokedex,
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

	res, err := cfg.pokeapi.ListLocationAreas(url)
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
	res, err := cfg.pokeapi.GetLocationArea(url)
	if err != nil {
		return err
	}
	fmt.Println("Found Pokemon:")
	for _, encounter := range res.PokemonEncounters {
		fmt.Printf(" - %v\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Must include name of pokemon you want to catch")
	}

	pokemonName := args[0]
	url := pokeapi.BasePokemonURL + pokemonName

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	res, err := cfg.pokeapi.GetPokemon(url)
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

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Must include name of pokemon to inspect")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that Pokemon")
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println(" -", typeInfo.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	names := make([]string, 0, len(cfg.pokedex))

	for name := range cfg.pokedex {
		names = append(names, name)
	}

	for _, name := range names {
		fmt.Println("-", name)
	}
	return nil
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	pokeClient := pokeapi.NewClient(5 * time.Minute)
	cfg := &config{
		pokeapi: pokeClient,
		pokedex: make(map[string]pokeapi.Pokemon),
	}
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
