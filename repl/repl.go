package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	clicommand "github.com/remcous/bootdev_pokedex/cliCommand"
)

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())

		commandName := input[0]

		if cmd, ok := clicommand.GetCommands()[commandName]; ok {
			err := cmd.Callback()
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