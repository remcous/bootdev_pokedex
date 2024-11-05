package repl

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"catch": {
			Name:        "catch <location_name>",
			Description: "Attempt to catch a pokemon",
			Callback:    CommandCatch,
		},
		"explore": {
			Name:        "explore <location_name>",
			Description: "Explore a location",
			Callback:    CommandExplore,
		},
		"map": {
			Name:        "map",
			Description: "Display next map locations",
			Callback:    CommandMapNext,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display previous map locations",
			Callback:    CommandMapPrev,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    CommandExit,
		},
	}
}

func CommandHelp(cfg *Config, _ ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()

	return nil
}

func CommandExit(cfg *Config, _ ...string) error {
	os.Exit(0)
	return nil
}

func CommandMapNext(cfg *Config, _ ...string) error {
	locationAreas, err := cfg.apiClient.GetLocationList(cfg.LocationAreasNext)
	if err != nil {
		return err
	}

	cfg.LocationAreasNext = locationAreas.Next
	cfg.LocationAreasPrev = locationAreas.Prev

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func CommandMapPrev(cfg *Config, _ ...string) error {
	if cfg.LocationAreasPrev == nil {
		return errors.New("you're on the first page")
	}

	locationAreas, err := cfg.apiClient.GetLocationList(cfg.LocationAreasPrev)
	if err != nil {
		return err
	}

	cfg.LocationAreasNext = locationAreas.Next
	cfg.LocationAreasPrev = locationAreas.Prev

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func CommandExplore(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("you must specify where to explore")
	}

	locationName := args[0]
	location, err := cfg.apiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}

func CommandCatch(cfg *Config, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("you must specify pokemon to try and catch")
	}

	pokemonName := args[0]
	pokemon, err := cfg.apiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s\n", pokemon.Name)
	caught := rand.Intn(pokemon.BaseExperience) < 40

	if caught {
		fmt.Printf("%s was caught!\n", pokemon.Name)

		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
