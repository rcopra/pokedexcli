package pokeapi

import (
	"time"

	"github.com/rcopra/pokedexcli/internal/pokecache"
)

const BaseLocationAreaURL = "https://pokeapi.co/api/v2/location-area/"
const BasePokemonURL = "https://pokeapi.co/api/v2/pokemon/"

type Client struct {
	cache pokecache.Cache
}

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		cache: *pokecache.NewCache(cacheInterval),
	}
}
