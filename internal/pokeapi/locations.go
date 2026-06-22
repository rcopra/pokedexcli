package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const BaseLocationAreaURL = "https://pokeapi.co/api/v2/location-area/"

type LocationAreas struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func ListLocationAreas(url string) (LocationAreas, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	var locations LocationAreas
	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationAreas{}, err
	}
	return locations, nil
}
