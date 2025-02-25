package main

import (
	"encoding/json"
	"fmt"
)

type MapPages struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *Config, p string) error {
	if config.Next == "" {
		config.Next = "https://pokeapi.co/api/v2/location-area/"
	}
	data, err := getAndCache(config.Next)
	if err != nil {
		return fmt.Errorf("error attempting to get or retreive cache: %w", err)
	}
	mapPages := MapPages{}
	err = json.Unmarshal(data, &mapPages)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	config.Next = mapPages.Next
	config.Previous = mapPages.Previous

	for _, result := range mapPages.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapb(config *Config, p string) error {
	fmt.Println(config.Previous)
	if config.Previous == "" {
		fmt.Println("you're on the first page")
	} else {
		config.Next = config.Previous
		commandMap(config, "")
	}
	return nil
}
