package main

import (
	"time"

	"github.com/remcous/bootdev_pokedex/repl"
)

const (
	clientTimeout = time.Second
	cacheInterval = 5 * time.Minute
)

func main() {
	cfg := repl.NewConfig(clientTimeout, cacheInterval)

	repl.StartRepl(cfg)
}
