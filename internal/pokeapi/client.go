package pokeapi

import (
	"net/http"
	"time"

	"github.com/remcous/bootdev_pokedex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	HttpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		HttpClient: http.Client{
			Timeout: timeout,
		},
	}
}
