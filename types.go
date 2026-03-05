package main

type config struct {
	Next     string
	Previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
