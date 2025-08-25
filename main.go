package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	var output []string
	rawSplits := strings.Fields(text)

	for i := range rawSplits {
		output = append(output, strings.ToLower(rawSplits[i]))
	}

	return output
}
