package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())

		commandName := input[0]

		if cmd, ok := GetCommands()[commandName]; ok {
			err := cmd.Callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
