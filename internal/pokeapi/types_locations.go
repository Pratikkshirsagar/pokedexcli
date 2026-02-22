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

type Location struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon NamedAPIResource `json:"pokemon"`
}

type Pokemon struct {
	Name           string  `json:"name"`
	Height         int     `json:"height"`
	Weight         int     `json:"weight"`
	BaseExperience int     `json:"base_experience"`
	Stats          []State `json:"stats"`
	Types          []Type  `json:"types"`
}

type State struct {
	BaseState int              `json:"base_stat"`
	Effort    int              `json:"effort"`
	Stat      NamedAPIResource `json:"stat"`
}

type Type struct {
	Slot int              `json:"slot"`
	Type NamedAPIResource `json:"type"`
}
