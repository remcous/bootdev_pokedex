package main

import (
	"github.com/remcous/bootdev_pokedex/repl"
)

func main() {
	cfg := repl.Config{}

	repl.StartRepl(&cfg)
}
