package main

import (
	"time"

	"github.com/yash-278/pokedexcli/internal/pokedexapi"
)

func main() {
	pokeClient := pokedexapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       map[string]pokedexapi.Pokemon{},
	}

	startRepl(cfg)
}
