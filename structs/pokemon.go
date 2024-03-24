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

type Pokemon struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Abilities []Ability `json:"abilities"`
	Types     []string  `json:"types"`
}

type PokeID struct {
	ID int `json:"id"`
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

type TypeResult struct {
	Name string `json:"name"`
}

type ColorResponse struct {
	Results []ColorResult `json:"results"`
}

type ColorResult struct {
	Name string `json:"name"`
}

type GenerationsResponse struct {
	Results []GenerationResult `json:"results"`
}

type GenerationResult struct {
	Name string `json:"name"`
}

type PokemonListResponse struct {
	Results []PokemonResult `json:"results"`
}

type PokemonResult struct {
	Type []string `json:"types"`
}
