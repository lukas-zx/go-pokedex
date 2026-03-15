package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"

	"github.com/lukas-zx/go-pokedex/internal/pokeapi"
)

func commandExit(config *config, params []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config, params []string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCliCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func getResponseBody(config *config, url string) ([]byte, error) {
	if data, ok := config.Cache.Get(url); ok {
		// fmt.Printf("using cached value for url %v", url)
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("error making http request to %v: %v", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("http request to %v failed with status code %v", url, res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading response body: %v", err)
	}

	config.Cache.Add(url, body)
	return body, nil
}


func mapHelper(config *config, url string) error {
	body, err := getResponseBody(config, url)
	if err != nil {
		return err
	}

	locationAreas := pokeapi.LocationAreas{}
	if err = json.Unmarshal(body, &locationAreas); err != nil {
		return fmt.Errorf("error unmarshalling response body: %v", err)
	}

	config.Next = locationAreas.Next
	config.Previous = locationAreas.Previous
	for _, location := range locationAreas.Results {
		fmt.Println(location.Name)
	}

	return nil
}
func commandMap(config *config, params []string) error {
	if err := mapHelper(config, config.Next); err != nil {
		return err
	}
	return nil
}
func commandMapb(config *config, params []string) error {
	if config.Previous == nil {
		return fmt.Errorf("no previous data available")
	}

	if err := mapHelper(config, *config.Previous); err != nil {
		return err
	}

	return nil
}

func commandExplore(config *config, params []string) error {
	location := params[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", location)
	body, err := getResponseBody(config, url)
	if err != nil {
		return err
	}

	locationArea := pokeapi.LocationArea{}
	if err = json.Unmarshal(body, &locationArea); err != nil {
		return fmt.Errorf("error unmarshalling response body: %v", err)
	}

	for _, entry := range locationArea.PokemonEncounters {
    fmt.Println(entry.Pokemon.Name)
  }

	return nil
}

func commandCatch(config *config, params []string) error {
	pokemon := params[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemon)
	body, err := getResponseBody(config, url)
	if err != nil {
		return err
	}

	pokeStruct := pokeapi.Pokemon{}
	if err = json.Unmarshal(body, &pokeStruct); err != nil {
		return fmt.Errorf("error unmarshalling response body: %v", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokeStruct.Name)

	catchChance := int(float64(pokeStruct.BaseExperience) * 0.8)
  if rand.Intn(pokeStruct.BaseExperience) < catchChance {
		config.Pokedex[pokeStruct.Name] = pokeStruct
		fmt.Printf("%s was caught!\n", pokeStruct.Name)
  } else {
		fmt.Printf("%s escaped!\n", pokeStruct.Name)
	}

	return nil
}

func commandInspect(config *config, params []string) error {
	name := params[0]
	pokemon, ok := config.Pokedex[name]
	if !ok {
		fmt.Printf("%s has not been caught yet\n", name)
		return nil
	} 

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Heigth: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pType := range pokemon.Types {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}

	return nil
}

func commandPokedex(config *config, params []string) error {
	fmt.Println("Your Pokedex:")
  for name := range config.Pokedex {
    fmt.Printf("  - %s\n", name)
  }

	return nil
}
