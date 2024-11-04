package repl

import (
	"time"

	"github.com/remcous/bootdev_pokedex/internal/pokeapi"
)

type Config struct {
	apiClient         pokeapi.Client
	LocationAreasNext *string
	LocationAreasPrev *string
}

func NewConfig(clientTimeout, cacheInterval time.Duration) *Config {
	client := pokeapi.NewClient(clientTimeout, cacheInterval)

	return &Config{
		apiClient: client,
	}
}
