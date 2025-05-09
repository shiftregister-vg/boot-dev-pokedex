package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/shiftregister-vg/pokedexcli/internal/pokecache"
)

const (
	prompt  = "Pokedex > "
	baseURL = "https://pokeapi.co/api/v2"
)

var (
	cache             *pokecache.Cache
	commands          map[string]cliCommand
	errNotImplemented error = errors.New("not implemented")
	locationsNextPage int   = 0
	locationsLimit    int   = 20
	pokedex           Pokedex
)

func init() {
	cache = pokecache.NewCache(time.Second * 5)
	pokedex = make(Pokedex)
}

func main() {
	for {
		printPrompt(false)

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			var args []string

			input := cleanInput(scanner.Text())
			commandWord := input[0]

			if len(input) > 1 {
				args = input[1:]
			}

			if command, ok := commands[commandWord]; ok {
				if err := command.callback(args...); err != nil {
					fmt.Printf("Could not execute command [%s]: %v", commandWord, err)
				}
			} else {
				fmt.Println("Unknown command")
			}

			printPrompt(true)
		}

	}
}

func printPrompt(newline bool) {
	p := prompt
	if newline {
		p = "\n" + p
	}
	fmt.Print(p)
}

func cleanInput(text string) []string {
	return strings.Split(strings.TrimSpace(strings.ToLower(text)), " ")
}
