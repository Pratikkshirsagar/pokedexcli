package main

import (
	"math/rand"
	"time"

	"github.com/Pratikkshirsagar/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		rng:           rand.New(rand.NewSource(time.Now().UnixNano())),
		pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
