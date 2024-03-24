package backend

import (
	"encoding/json"
	"io"
	"os"
	"pokeapi/structs"
)

func LoadData() ([]structs.Json, error) {

	fileData, err := os.ReadFile("data.json")
	if err != nil {
		return nil, err
	}

	var forms []structs.Json

	if len(fileData) != 0 {
		err = json.Unmarshal(fileData, &forms)
		if err != nil {
			return nil, err
		}
	}

	return forms, nil
}

func ReadFavorites() ([]string, error) {
	// Open the JSON file
	file, err := os.Open("favorites.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the JSON data from the file
	jsonData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a slice of strings
	var favorites []string
	if err := json.Unmarshal(jsonData, &favorites); err != nil {
		return nil, err
	}

	return favorites, nil
}

// SaveFavoritesToJSONFile saves the favorites data to a JSON file.
func SaveFavorites(favorites []string) error {
	// Marshal the favorites data into JSON format
	jsonData, err := json.Marshal(favorites)
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	if err := os.WriteFile("favorites.json", jsonData, 0644); err != nil {
		return err
	}

	return nil
}
