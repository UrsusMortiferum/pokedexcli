package main

import "fmt"

func commandHelp(cfg *config) error {
	commandOrder := []string{"help", "map", "mapb", "exit"}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, commandName := range commandOrder {
		command := getCommands()[commandName]
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
