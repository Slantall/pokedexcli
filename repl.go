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
	Callback    func(config *Config, parameter string) error
}

var commands map[string]CliCommand
var cache *pokecache.Cache
var pokedex map[string]Pokemon

func init() {

	cache = pokecache.NewCache(5 * time.Second)
	pokedex = make(map[string]Pokemon)
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
			Description: "Display a list the next 20 locations. Will start with the first 20 locations when first ran.",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display a list the last 20 locations. Will show nothing if map hasn't ben ran yet.",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Requires a name of a location to be input. Display a list available Pokemon in a given area.",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Allows you to attempt to catch a Pokemon to add to your Pokedex.",
			Callback:    commandCatch,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := cleanInput(scanner.Text())
		firstWord := userInput[0]
		secondWord := ""
		if len(userInput) > 1 {
			secondWord = userInput[1]
		}
		command, ok := commands[firstWord]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			command.Callback(&config, secondWord)
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
