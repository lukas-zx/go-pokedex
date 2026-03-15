package main

import (
	"strings"
)

func getCliCommands() map[string]cliCommand {
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
			description: "Prints next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Prints previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Prints all Pokemon in given location area, usage: explore <location name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon, usage: catch <pokemon name>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon, usage: inspect <pokemon name>",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show all caught Pokemon",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	res := []string{}
	words := strings.FieldsSeq(text)
	for word := range words {
		res = append(res, strings.ToLower(strings.TrimSpace(word)))
	}
	return res
}
