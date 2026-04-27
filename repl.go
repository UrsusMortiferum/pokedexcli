package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		output := words[0]

		fmt.Printf("You command was: %s\n", output)
	}

}
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
