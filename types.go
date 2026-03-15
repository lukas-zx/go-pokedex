package main

import (
	"github.com/lukas-zx/go-pokedex/internal/pokeapi"
	"github.com/lukas-zx/go-pokedex/internal/pokecache"
)

type config struct {
	Next     string
	Previous *string
	Cache    *pokecache.Cache
	Pokedex  map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}
