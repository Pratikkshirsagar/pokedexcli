package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemon, err := cfg.pokeapiClient.GetInformation(pokemonName)
	if err != nil {
		return err
	}

	roll := cfg.rng.Intn(100)
	catchChance := 100 - pokemon.BaseExperience
	if catchChance < 5 {
		catchChance = 5
	}
	if catchChance > 95 {
		catchChance = 95
	}

	if roll < catchChance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
