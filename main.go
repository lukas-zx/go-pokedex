package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	cliCommands := getCliCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		commandString := cleanedInput[0]
		if command, ok := cliCommands[commandString]; ok {
			command.callback()
		} else {
			fmt.Printf("unknown command: %s\n", commandString)
		}
	}
}
