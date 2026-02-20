package pokeapi

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type RespLocation struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon NamedAPIResource `json:"pokemon"`
}
