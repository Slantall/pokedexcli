package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func commandMap(config *Config) error {
	if config.Next == "" {
		config.Next = "https://pokeapi.co/api/v2/location-area/"
	}
	data, found := cache.Get(config.Next)
	if !found {
		fmt.Println("Cache miss - making HTTP request...")
		res, err := http.Get(config.Next)
		if err != nil {
			return fmt.Errorf("error making request: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if res.StatusCode > 299 {
			fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, data)
		}
		if err != nil {
			return fmt.Errorf("error reading response: %w", err)
		}
		cache.Add(config.Next, data)
	} else {
		fmt.Println("Cache hit!")
	}
	mapPages := MapPages{}
	err := json.Unmarshal(data, &mapPages)
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

func commandMapb(config *Config) error {
	fmt.Println(config.Previous)
	if config.Previous == "" {
		fmt.Println("you're on the first page")
	} else {
		config.Next = config.Previous
		commandMap(config)
	}
	return nil
}
