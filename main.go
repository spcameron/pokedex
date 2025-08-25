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

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    displayHelp,
		},
	}
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
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func displayHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	cliCommands := getCommands()
	for _, v := range cliCommands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		command := input[0]

		cliCommands := getCommands()
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
