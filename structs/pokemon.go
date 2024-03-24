package structs

type PokemonDetails struct {
	ID                     int            `json:"id"`
	Name                   string         `json:"name"`
	Height                 int            `json:"height"`
	Weight                 int            `json:"weight"`
	Forms                  PokemonForm    `json:"forms"`
	Species                PokemonSpecies `json:"species"`
	Type                   []string       `json:"types"`
	Image                  string         `json:"image"`
	Abilities              []Ability      `json:"abilities"`
	LocationAreaEncounters string         `json:"location_area_encounters"`
	Generation             string         `json:"generation"`
	Color                  string         `json:"color"`
}

type Json struct {
	Type                   []string `json:"types"`
	Generation             string   `json:"generation"`
	LocationAreaEncounters string   `json:"location_area_encounters"`
}

type PokemonForm struct {
	Name string `json:"name"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

type Ability struct {
	Name string `json:"name"`
}

type Generation struct {
	Name string `json:"name"`
}

type TypesResponse struct {
	Results []TypeResult `json:"results"`
}

// TypeResult represents a single type result from the PokeAPI.
type TypeResult struct {
	Name string `json:"name"`
}

// LocationAreasResponse represents the response structure from the PokeAPI for location areas.
type ColorResponse struct {
	Results []ColorResult `json:"results"`
}

// LocationAreaResult represents a single location area result from the PokeAPI.
type ColorResult struct {
	Name string `json:"name"`
}

// GenerationsResponse represents the response structure from the PokeAPI for generations.
type GenerationsResponse struct {
	Results []GenerationResult `json:"results"`
}

// GenerationResult represents a single generation result from the PokeAPI.
type GenerationResult struct {
	Name string `json:"name"`
}

// PokemonListResponse represents the response structure from the PokeAPI for a list of Pokémon.
type PokemonListResponse struct {
	Results []PokemonResult `json:"results"`
}

// PokemonResult represents a single Pokémon result from the PokeAPI.
type PokemonResult struct {
	Type []string `json:"types"`
}
