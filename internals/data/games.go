package data

import (
	"eliest/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)


func AllGames(path string) ([]models.Game, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var loadedGames []models.Game
	err = json.Unmarshal(content, &loadedGames)
	if err != nil {
		return []models.Game{}, err
	}
	return loadedGames, nil
}

func FindGame(id, path string) (models.Game, error) {
	allGames, err := AllGames(path)

	if err != nil {
		return models.Game{}, err
	}
	for _, n := range allGames {
		if n.Id == id {
			return n, nil
		}
	}
	return models.Game{}, errors.New("Not Found")
}

func UseAllGames() ([]models.Game, error) {
	cwd := os.Getenv("WorkStation")
	genFilePath := filepath.Join(cwd, "games.json")
	return AllGames(genFilePath)
}

func UseFindGame(id string) (models.Game, error) {
	cwd := os.Getenv("WorkStation")
	genFilePath := filepath.Join(cwd, "games.json")
	return FindGame(id, genFilePath)
}
