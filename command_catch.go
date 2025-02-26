package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func commandCatch(_ *Config, name string) error {

	url := "https://pokeapi.co/api/v2/pokemon/" + name

	data, err := getAndCache(url)
	if data == nil {
		fmt.Println("Pokemon not found. Use 'explore' to find valid Pokemon. Use 'map' for valid locations to explore.")
		return nil
	}
	if err != nil {
		return fmt.Errorf("error attempting to get or retreive cache: %w", err)
	}
	pokemonInfo := Pokemon{}
	err = json.Unmarshal(data, &pokemonInfo)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonInfo.Name)
	catchChance := 1 - (float32(pokemonInfo.BaseExperience) / 300)
	if catchChance < 0.05 {
		catchChance = 0.05
	}
	roll := float32(rand.Intn(100)) / 100
	//fmt.Println(catchChance, "catch chance", roll, "roll")
	if catchChance > roll {
		fmt.Println(pokemonInfo.Name, "was caught!")
		pokedex[pokemonInfo.Name] = pokemonInfo
	} else {
		fmt.Println(pokemonInfo.Name, "escaped!")
	}

	return nil
}
