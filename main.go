package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	cliCommands := getCliCommands()
	scanner := bufio.NewScanner(os.Stdin)
	config := &config{
		Next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		Previous: nil,
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		commandString := cleanedInput[0]
		if command, ok := cliCommands[commandString]; ok {
			command.callback(config)
		} else {
			fmt.Printf("unknown command: %s\n", commandString)
		}
	}
}
