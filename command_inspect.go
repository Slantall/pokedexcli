package main

import "fmt"

func commandInspect(_ *Config, name string) error {
	pokemon, ok := pokedex[name]
	if !ok {
		fmt.Println("Pokemon not found in Pokedex.")
		return nil
	}
	fmt.Println("Name:", pokemon.Name, "\nHeight:", pokemon.Height, "\nWegiht:", pokemon.Weight, "\nStats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}
	return nil
}
