package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokecache"
	"strings"
	"time"
)

type Config struct {
	Next     string
	Previous string
}

var config Config

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *Config) error
}

var commands map[string]CliCommand
var cache *pokecache.Cache

func init() {

	cache = pokecache.NewCache(5 * time.Second)
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	commands = map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "display a list the next 20 locations. Will start with the first 20 locations when first ran.",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "display a list the last 20 locations. Will start with the first 20 locations when first ran.",
			Callback:    commandMapb,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		firstWord := cleanInput(scanner.Text())[0]
		command, ok := commands[firstWord]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			command.Callback(&config)
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
