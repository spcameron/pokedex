package main

import (
	"fmt"
)

func commandMapForward(config *config) error {
	locationAreaResponse, err := config.apiClient.getLocationAreaBatch(config.Next)
	if err != nil {
		return err
	}

	config.Next = locationAreaResponse.Next
	config.Previous = locationAreaResponse.Previous

	for _, locationArea := range locationAreaResponse.Results {
		fmt.Printf("%s\n", locationArea.Name)
	}

	return nil
}

func commandMapBackward(config *config) error {
	if config.Previous == nil {
		return fmt.Errorf("You're on the first page!")
	}

	locationAreaResponse, err := config.apiClient.getLocationAreaBatch(config.Previous)
	if err != nil {
		return err
	}

	config.Next = locationAreaResponse.Next
	config.Previous = locationAreaResponse.Previous

	for _, locationArea := range locationAreaResponse.Results {
		fmt.Printf("%s\n", locationArea.Name)
	}

	return nil

}
