package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func cleanInput(text string) []string {
	var output []string
	rawSplits := strings.Fields(text)

	for i := range rawSplits {
		output = append(output, strings.ToLower(rawSplits[i]))
	}

	return output
}

func commandExit() error {
	fmt.Println("Closing the Pokedex ... Goodbye!")
	os.Exit(0)
	return nil
}

func main() {
	cliCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		command := input[0]

		cliCommand, ok := cliCommands[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cliCommand.callback()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}
