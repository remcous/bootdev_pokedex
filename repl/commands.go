package repl

import (
	"errors"
	"fmt"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the pokedex",
			Callback:    CommandExit,
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
	}
}

func CommandHelp(cfg *Config) error {
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

func CommandExit(cfg *Config) error {
	os.Exit(0)
	return nil
}

func CommandMapNext(cfg *Config) error {
	locationResp, err := cfg.apiClient.GetLocationList(cfg.LocationAreasNext)
	if err != nil {
		return err
	}

	cfg.LocationAreasNext = locationResp.Next
	cfg.LocationAreasPrev = locationResp.Prev

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func CommandMapPrev(cfg *Config) error {
	if cfg.LocationAreasPrev == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.apiClient.GetLocationList(cfg.LocationAreasPrev)
	if err != nil {
		return err
	}

	cfg.LocationAreasNext = locationResp.Next
	cfg.LocationAreasPrev = locationResp.Prev

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
