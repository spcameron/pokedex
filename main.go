package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %s\n", input[0])
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
