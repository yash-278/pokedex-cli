package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("location name cannot be empty")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.pokedex[pokemonName]

	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}
	return nil

}
