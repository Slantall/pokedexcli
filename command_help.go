package main

import "fmt"

func commandHelp(_ *Config, _ string) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}
	return nil
}
