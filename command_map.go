package main

import (
	"fmt"
)

func commandMapForward(config *config, arguments ...string) error {
	resp, err := config.apiClient.GetLocationAreaBatch(config.Next)
	if err != nil {
		return err
	}

	config.Next = resp.Next
	config.Previous = resp.Previous

	for _, locationArea := range resp.Results {
		fmt.Printf("%s\n", locationArea.Name)
	}

	return nil
}

func commandMapBackward(config *config, arguments ...string) error {
	if config.Previous == nil {
		return fmt.Errorf("You're on the first page!")
	}

	resp, err := config.apiClient.GetLocationAreaBatch(config.Previous)
	if err != nil {
		return err
	}

	config.Next = resp.Next
	config.Previous = resp.Previous

	for _, locationArea := range resp.Results {
		fmt.Printf("%s\n", locationArea.Name)
	}

	return nil

}
