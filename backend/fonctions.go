package backend

import (
	"encoding/json"
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

	fileData, err := os.ReadFile("favorites.json")
	if err != nil {

		if os.IsNotExist(err) || len(fileData) == 0 {
			return []string{}, nil
		}
		return nil, err
	}

	var favorites []string
	err = json.Unmarshal(fileData, &favorites)
	if err != nil {
		return nil, err
	}

	return favorites, nil
}

func SaveFavorites(favorites []string) error {

	jsonData, err := json.Marshal(favorites)
	if err != nil {
		return err
	}

	if err := os.WriteFile("favorites.json", jsonData, 0644); err != nil {
		return err
	}

	return nil
}
