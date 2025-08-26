package main

import "fmt"

func commandPokedex(config *config, arguments ...string) error {
	if len(config.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range config.Pokedex {
		fmt.Printf("  - %s\n", name)
	}

	return nil
}
