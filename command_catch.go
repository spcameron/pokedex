package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/spcameron/pokedex/internal/pokeapi"
)

func commandCatch(config *config, arguments ...string) error {
	if len(arguments) == 0 {
		return fmt.Errorf("expected arguments but received none")
	}

	pokemon, err := config.apiClient.GetPokemon(arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	baseExp := pokemon.BaseExperience
	p := pokeapi.GetCatchProbability(baseExp)

	if rand.Float64() < p {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.Pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
