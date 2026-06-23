package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreaPage struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func (c *Client) ListLocationAreas(url string) (LocationAreaPage, error) {
	cacheData, ok := c.cache.Get(url)
	if ok {
		var locations LocationAreaPage
		if err := json.Unmarshal(cacheData, &locations); err != nil {
			return LocationAreaPage{}, err
		}
		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreaPage{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaPage{}, err
	}
	c.cache.Add(url, data)
	var locations LocationAreaPage
	if err := json.Unmarshal(data, &locations); err != nil {
		return LocationAreaPage{}, err
	}
	return locations, nil
}
