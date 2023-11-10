package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, arg []string) error {

	if len(arg) == 0 {
		return errors.New("pokemon name cannot be empty")
	}

	pokemonName := arg[0]

	fmt.Printf("Throwing a Pokeball at %v\n", pokemonName)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	if pokemon.Name != "" {
		catchValue := rand.Intn(pokemon.BaseExperience)

		if catchValue > 40 {
			fmt.Printf("%v was caught!\n", pokemonName)
			cfg.pokedex[pokemonName] = pokemon

			return nil
		}
	}

	fmt.Printf("%v escaped!\n", pokemonName)

	return nil
}
