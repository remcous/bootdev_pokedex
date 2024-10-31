package repl

import "github.com/remcous/bootdev_pokedex/api"

type Config struct {
	apiClient         api.Client
	LocationAreasNext *string
	LocationAreasPrev *string
}
