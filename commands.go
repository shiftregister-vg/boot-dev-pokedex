package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	commands = map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "Catches named pokemon based on chance and pokemon experience.",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore",
			description: "Displays Pokemon in given area. Usage: `explore pastoria-city-area`",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays infor about a pokemon from your pokedex",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays the names of location around the world of Pokemon",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of names of location around the world of Pokemon",
			callback:    commandMapb,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints a list of the pokemon in your Pokedex.",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp(args ...string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")

	for _, command := range commands {
		fmt.Printf("\n%s: %s", command.name, command.description)
	}
	return nil
}

func commandExit(args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("os.Exit(0) failed, this is really odd behavior...")
}

func commandMap(args ...string) error {
	fmt.Println("")
	locationAreas := getLocationAreas(locationsLimit, locationsNextPage)
	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}
	locationsNextPage += locationsLimit
	return nil
}

func commandMapb(args ...string) error {
	offset := locationsNextPage - locationsLimit
	if offset < 0 {
		offset = 0
		locationsNextPage = 0
	} else {
		locationsNextPage = offset - locationsLimit
	}

	fmt.Println("")
	locationAreas := getLocationAreas(locationsLimit, locationsNextPage)
	for _, locationArea := range locationAreas {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func commandExplore(args ...string) error {
	if len(args) == 0 || len(args) > 1 {
		err := errors.New("Invalid number of arguments. Usage: `explore pastoria-city-area`.")
		return err
	}

	locationAreaName := args[0]
	locationArea := getLocationArea(locationAreaName)

	fmt.Println("")
	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}

// checkCatchSuccess takes a Pokemon's base experience value and returns true if the catch is successful.
// Higher base experience values decrease the probability of a successful catch.
func checkCatchSuccess(baseExperience int) bool {
	// Initialize the random number generator with current time as seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random number between 0 and 100
	randomNum := r.Intn(101)

	// Calculate catch chance using a formula that scales better with high experience values
	// We'll use: 100 / (1 + (baseExperience / 100))
	// This means:
	// - A Pokemon with 0 base experience has 100% catch rate
	// - A Pokemon with 100 base experience has 50% catch rate
	// - A Pokemon with 300 base experience has 25% catch rate
	// - A Pokemon with 900 base experience has 10% catch rate
	catchChance := 100 / (1 + (baseExperience / 100))

	// Ensure catch chance doesn't go below 1% or above 100%
	if catchChance < 1 {
		catchChance = 1
	} else if catchChance > 100 {
		catchChance = 100
	}

	// Return true if random number is less than catch chance
	return randomNum < catchChance
}

func commandCatch(args ...string) error {
	if len(args) == 0 || len(args) > 1 {
		err := errors.New("Invalid number of arguments. Usage: `catch pikachu`.")
		return err
	}

	pokemonName := args[0]

	fmt.Printf("\nThrowing a Pokeball at %s...\n", pokemonName)

	pokemon := getPokemon(pokemonName)
	caught := checkCatchSuccess(pokemon.BaseExperience)

	if caught {
		// add this pokemon to our pokedex
		pokedex[pokemonName] = pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func commandInspect(args ...string) error {
	if len(args) == 0 || len(args) > 1 {
		err := errors.New("Invalid number of arguments. Usage: `inspect pikachu`.")
		return err
	}

	pokemonName := args[0]

	pokemon, ok := pokedex[pokemonName]
	if !ok {
		return fmt.Errorf("%[1]s is not in your Pokedex. Try catching one with `catch %[1]s`", pokemonName)
	}

	fmt.Println(pokemon)
	return nil
}

func commandPokedex(args ...string) error {
	fmt.Println(pokedex)
	return nil
}
