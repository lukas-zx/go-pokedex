package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/lukas-zx/go-pokedex/internal/pokeapi"
	"github.com/lukas-zx/go-pokedex/internal/pokecache"
)

func main() {
	cliCommands := getCliCommands()
	scanner := bufio.NewScanner(os.Stdin)
	config := &config{
		Next:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		Previous: nil,
		Cache:    pokecache.NewCache(5 * time.Second),
		Pokedex:  make(map[string]pokeapi.Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		commandString := cleanedInput[0]
		params := cleanedInput[1:]
		if command, ok := cliCommands[commandString]; ok {
			command.callback(config, params)
		} else {
			fmt.Printf("unknown command: %s\n", commandString)
		}
	}
}
