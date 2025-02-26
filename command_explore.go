package main

import (
	"encoding/json"
	"fmt"
)

type AreaInfo struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(_ *Config, name string) error {

	url := "https://pokeapi.co/api/v2/location-area/" + name

	data, err := getAndCache(url)
	if data == nil {
		fmt.Println("Location not found. Use 'map' for valid locations.")
		return nil
	}
	if err != nil {
		return fmt.Errorf("error attempting to get or retreive cache: %w", err)
	}
	areaInfo := AreaInfo{}
	err = json.Unmarshal(data, &areaInfo)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	for _, pokemon := range areaInfo.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
