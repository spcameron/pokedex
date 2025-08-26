package main

import (
	"time"

	"github.com/spcameron/pokedex/internal/pokeapi"
)

func main() {
	config := &config{
		apiClient: pokeapi.NewClient(5 * time.Minute),
		Pokedex:   map[string]pokeapi.Pokemon{},
	}

	startREPL(config)
}
