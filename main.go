package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputString := scanner.Text()
		inputList := cleanInput(inputString)
		fmt.Printf("Your command was: %v\n", inputList[0])
	}
}
