package main

import "fmt"

func commandPokedex(_ *Config, _ string) error {
	fmt.Println("Your Pokedex")
	if len(pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon yet! You can 'explore' a 'map' to try to 'catch' one!")
	}
	for _, pokemon := range pokedex {
		fmt.Println(" -", pokemon.Name)
	}
	return nil
}
