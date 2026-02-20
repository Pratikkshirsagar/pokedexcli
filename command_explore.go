package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {

	if len(args) < 1 {
		return fmt.Errorf("usage: explore <location-area-name>")
	}

	name := args[0]

	locationReso, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", name)
	fmt.Println("Found Pokemon:")
	for _, enc := range locationReso.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
