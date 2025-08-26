package main

import "fmt"

func commandInspect(config *config, arguments ...string) error {
	if len(arguments) == 0 {
		return fmt.Errorf("expected arguments but received none")
	}

	pokemonName := arguments[0]
	pokemon, exists := config.Pokedex[pokemonName]
	if !exists {
		fmt.Printf("you have not caught %s\n", pokemonName)
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}
