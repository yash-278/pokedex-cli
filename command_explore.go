package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, arg []string) error {
	if len(arg) == 0 {
		return errors.New("location name cannot be empty")
	}

	locationName := arg[0]

	fmt.Printf("Exploring %v...\n", locationName)

	locationResp, err := cfg.pokeapiClient.ExploreArea(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationResp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
