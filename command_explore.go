package main

import (
	"fmt"
)

func commandExplore(config *config, arguments ...string) error {
	if len(arguments) == 0 {
		return fmt.Errorf("expected arguments but received none")
	}

	locationAreaResponse, err := config.apiClient.GetLocationArea(arguments[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", arguments[0])
	fmt.Println("Found Pokemon:")
	encounters := locationAreaResponse.PokemonEncounters
	for _, encounter := range encounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
