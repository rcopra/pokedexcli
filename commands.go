package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
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
			description: "Usage: 'catch <pokemon name>' - attempt to catch a pokemon",
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
