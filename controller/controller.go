package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"pokeapi/backend"
	"pokeapi/structs"
	"pokeapi/templates"
	InitTemplate "pokeapi/templates"
	"strings"
	"time"
)

func Home(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "home", nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "about", nil)
}

func Collection(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "collection", nil)
}

func Favoris(w http.ResponseWriter, r *http.Request) {
	templates.Temp.ExecuteTemplate(w, "favoris", nil)
}

func GetRandomPokemon() ([]structs.PokemonDetails, error) {

	rand.Seed(time.Now().UnixNano())

	var pokemons []structs.PokemonDetails

	for i := 0; i < 20; i++ {

		pokemonID := rand.Intn(1024) + 1
		pokemonURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d", pokemonID)

		id, name, height, weight, types, abilities, locationAreaEncounters, image, error := GetPokeDetails(pokemonURL)
		if error != nil {
			fmt.Printf("Erreur de récupération du Pokémon numéro %d, %v", pokemonID, error)
			continue
		}

		pokemon := structs.PokemonDetails{
			ID:                     id,
			Name:                   name,
			Height:                 height,
			Weight:                 weight,
			Type:                   types,
			Abilities:              abilities,
			LocationAreaEncounters: locationAreaEncounters,
			Image:                  image,
		}
		pokemons = append(pokemons, pokemon)

	}
	return pokemons, nil
}

func Search(w http.ResponseWriter, r *http.Request) {

	searchQuery := r.URL.Query().Get("query")
	if searchQuery == "" {
		http.Error(w, "Veuillez entrer une donnée", http.StatusBadRequest)
		return
	}

	pokemonURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", searchQuery)

	id, name, height, weight, types, abilities, locationAreaEncounters, image, err := GetPokeDetails(pokemonURL)
	if err != nil {
		fmt.Printf("Impossible de récupérer les données pour %s: %v", searchQuery, err)
		if strings.Contains(err.Error(), "404") {
			InitTemplate.Temp.ExecuteTemplate(w, "search", map[string]string{
				"Erreur:": "Aucun Pokémon trouvé",
			})
		} else {
			http.Error(w, fmt.Sprintf("Erreur lors de la récupération des détails de Pokémon: %v", err), http.StatusInternalServerError)
			http.Redirect(w, r, "error", 404)
		}
		return
	}

	pokemon := structs.PokemonDetails{
		ID:                     id,
		Name:                   name,
		Height:                 height,
		Weight:                 weight,
		Type:                   types,
		LocationAreaEncounters: locationAreaEncounters,
		Abilities:              abilities,
		Image:                  image,
	}

	fmt.Printf("Données récupérées pour %s: %+v", searchQuery, pokemon)
	InitTemplate.Temp.ExecuteTemplate(w, "pokemon", pokemon)
}

func Pokemon(w http.ResponseWriter, r *http.Request) {
	pokemons, err := GetRandomPokemon()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des Pokémon", http.StatusInternalServerError)
		return
	}
	fmt.Printf("Nombre de Pokémons récupérés : %d\n", len(pokemons))

	InitTemplate.Temp.ExecuteTemplate(w, "categorie", pokemons)
}

func GetPokeDetails(pokemonURL string) (id int, name string, height int, weight int, types []string, abilities []structs.Ability, locationAreaEncounters string, image string, err error) {
	resp, err := http.Get(pokemonURL)
	if err != nil {
		return 0, "", 0, 0, nil, nil, "", "", err
	}
	defer resp.Body.Close()

	var details struct {
		ID                     int      `json:"id"`
		Name                   string   `json:"name"`
		Image                  string   `json:"image"`
		Height                 int      `json:"height"`
		Weight                 int      `json:"weight"`
		Type                   []string `json:"type"`
		LocationAreaEncounters string   `json:"location_area_encounters"`
		Forms                  []struct {
			Name string `json:"name"`
		} `json:"forms"`
		Species struct {
			Name string `json:"name"`
		} `json:"species"`
		Sprites struct {
			Other struct {
				OfficialArtwork struct {
					FrontDefault string `json:"front_default"`
				} `json:"official-artwork"`
			} `json:"other"`
		} `json:"sprites"`
		Types []struct {
			Type struct {
				Name string `json:"name"`
			} `json:"type"`
		} `json:"types"`
		Abilities []struct {
			Ability struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"ability"`
		} `json:"abilities"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&details); err != nil {
		return 0, "", 0, 0, nil, nil, "", "", err
	}

	id = details.ID
	name = details.Name
	image = details.Sprites.Other.OfficialArtwork.FrontDefault
	height = details.Height
	weight = details.Weight
	types = details.Type
	locationAreaEncounters = details.LocationAreaEncounters

	for _, t := range details.Types {
		types = append(types, t.Type.Name)
	}

	for _, a := range details.Abilities {
		abilities = append(abilities, structs.Ability{
			Name: a.Ability.Name,
		})
	}

	return id, name, height, weight, types, abilities, locationAreaEncounters, image, nil
}

func PokeDetails(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/pokemon/")

	id, name, height, weight, types, abilities, locationAreaEncounters, image, err := GetPokeDetails("https://pokeapi.co/api/v2/pokemon/" + name) // Ajouté abilities dans la récupération
	if err != nil {
		http.Error(w, "Pokémon introuvable", http.StatusNotFound)
		http.Redirect(w, r, "error", 404)
		return
	}
	if err != nil {
		fmt.Printf("Erreur lors de la récupération des données: %v", err)
	}

	pokemon := structs.PokemonDetails{
		ID:                     id,
		Name:                   name,
		Height:                 height,
		Weight:                 weight,
		Type:                   types,
		LocationAreaEncounters: locationAreaEncounters,
		Abilities:              abilities,
		Image:                  image,
	}

	InitTemplate.Temp.ExecuteTemplate(w, "pokemon", pokemon)
}

func GetPokemonDetails(name string) (structs.PokemonDetails, error) {
	id, name, height, weight, types, abilities, locationAreaEncounters, image, err := GetPokeDetails("https://pokeapi.co/api/v2/pokemon/" + name)
	if err != nil {
		return structs.PokemonDetails{}, err
	}

	return structs.PokemonDetails{
		ID:                     id,
		Name:                   name,
		Height:                 height,
		Weight:                 weight,
		Type:                   types,
		LocationAreaEncounters: locationAreaEncounters,
		Abilities:              abilities,
		Image:                  image,
	}, nil
}

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	//Retrieve types for passing them to the template.
	types, err := FetchPokemonTypes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*	//Retrieve generations for passing them to the template.
		generations, err := FetchPokemonGenerations()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Retrieve location area encounters for passing them to the template.
		color, err := FetchPokemonColor()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}*/

	data := struct {
		Types []string
		/*Generations      []string
		Color            []string*/
		PokemonsByFilter []structs.PokemonDetails
	}{
		Types: types,
		/*Generations: generations,
		Color:       color,*/
	}

	if r.Method == "POST" {
		typeName := r.FormValue("PokeType")
		/*generationName := r.FormValue("generation")
		colorName := r.FormValue("color")*/

		pokemons, err := FetchPokemonsByFilters(typeName /*, generationName, colorName */)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data.PokemonsByFilter = pokemons
	}

	fmt.Println(data.PokemonsByFilter)

	//Render the same template with the processed data
	InitTemplate.Temp.ExecuteTemplate(w, "collection", data)
}

func FetchPokemonsByFilters(typeName /*, generationName, colorName */ string) ([]structs.PokemonDetails, error) {
	// Call the API to get all Pokémon filtered by type, generation, and location area
	url := "https://pokeapi.co/api/v2/type/" + typeName

	/*if typeName != "" {
		url += "?type=" + typeName

		if generationName != "" {
			url += "&generation=" + generationName
		}
		if colorName != "" {
			url += "&color=" + colorName
		}

	}*/
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response struct {
		PokemonsByFilter []struct {
			PokemonsByFilter struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"pokemon"`
		} `json:"pokemon"`
	}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	var pokemon []structs.PokemonDetails
	for i, p := range response.PokemonsByFilter {
		if i >= 20 {
			break
		}

		id, name, height, weight, types, abilities, locationAreaEncounters, image, err := GetPokeDetails(p.PokemonsByFilter.Name)
		if err != nil {
			continue
		}

		pokemon = append(pokemon, structs.PokemonDetails{
			ID:                     id,
			Name:                   name,
			Height:                 height,
			Weight:                 weight,
			Type:                   types,
			LocationAreaEncounters: locationAreaEncounters,
			Abilities:              abilities,
			Image:                  image,
		})
	}
	return pokemon, nil
}

func FetchPokemonTypes() ([]string, error) {
	url := "https://pokeapi.co/api/v2/type/"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response structs.TypesResponse

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	var types []string
	for _, t := range response.Results {
		types = append(types, t.Name)
	}

	return types, nil
}

/*
func FetchPokemonGenerations() ([]string, error) {

		url := "https://pokeapi.co/api/v2/generation/"
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var response structs.GenerationsResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, err
		}

		var generations []string
		for _, g := range response.Results {
			generations = append(generations, g.Name)
		}

		return generations, nil
	}

func FetchPokemonColor() ([]string, error) {

		url := "https://pokeapi.co/api/v2/pokemon-color/"
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var response structs.ColorResponse
		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return nil, err
		}

		var color []string
		for _, l := range response.Results {
			color = append(color, l.Name)
		}

		return color, nil
	}
*/
func RenderPokemonPage(w http.ResponseWriter, r *http.Request) {

	pokemonsByType, err := FetchPokemonTypes()
	if err != nil {
		http.Error(w, "Failed to fetch Pokémon data by type", http.StatusInternalServerError)
		return
	}
	/*
		pokemonsByColor, err := FetchPokemonColor()
		if err != nil {
			http.Error(w, "Failed to fetch Pokémon data by color", http.StatusInternalServerError)
			return
		}

		pokemonsByGeneration, err := FetchPokemonGenerations()
		if err != nil {
			http.Error(w, "Failed to fetch Pokémon data by generation", http.StatusInternalServerError)
			return
		}
	*/

	data := struct {
		PokeType []string
		/*PokeColor        []string
		PokeGeneration   []string*/
		PokemonsByFilter []structs.PokemonDetails
	}{
		PokeType: pokemonsByType,
		/*	PokeColor:      pokemonsByColor,
			PokeGeneration: pokemonsByGeneration,*/
	}

	if r.Method == "POST" {
		typeName := r.FormValue("PokeType")
		/*	generationName := r.FormValue("PokeGeneration")
			colorName := r.FormValue("PokeColor")*/

		pokemons, err := FetchPokemonsByFilters(typeName /*, generationName, colorName*/)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data.PokemonsByFilter = pokemons
	}

	fmt.Println(data.PokemonsByFilter)

	if err := InitTemplate.Temp.ExecuteTemplate(w, "collection", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddToFavoritesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	pokemonID := r.FormValue("pokemon_id")

	favorites, err := backend.ReadFavorites()
	if err != nil {
		http.Error(w, "Failed to read favorites data", http.StatusInternalServerError)
		return
	}

	favorites = append(favorites, pokemonID)

	err = backend.SaveFavorites(favorites)
	if err != nil {
		http.Error(w, "Failed to save favorites data", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/favoris", http.StatusSeeOther)
}

func ViewFavoritesHandler(w http.ResponseWriter, r *http.Request) {

	favoriteIDs, err := ReadFavoritesFromJSONFile("favorites.json")
	if err != nil {
		if os.IsNotExist(err) {
			favoriteIDs = []string{}
		} else {
			fmt.Println("Error reading favorites:", err)
			http.Error(w, "Failed to read favorites", http.StatusInternalServerError)
			return
		}
	}

	var favoritePokemons []structs.Pokemon
	for _, id := range favoriteIDs {
		pokemon, err := FetchPokemonByID(id)
		if err != nil {
			fmt.Printf("Error fetching Pokémon data for ID %s: %v\n", id, err)
			continue
		}
		favoritePokemons = append(favoritePokemons, pokemon)
	}
	if err := InitTemplate.Temp.ExecuteTemplate(w, "favoris", favoriteIDs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ReadFavoritesFromJSONFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var favoriteIDs []string
	err = json.NewDecoder(file).Decode(&favoriteIDs)
	if err != nil {
		return nil, err
	}

	return favoriteIDs, nil
}

func FetchPokemonByID(id string) (structs.Pokemon, error) {

	apiUrl := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return structs.Pokemon{}, err
	}
	defer resp.Body.Close()

	var pokemon structs.Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return structs.Pokemon{}, err
	}

	return pokemon, nil
}
