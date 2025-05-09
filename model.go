package main

import (
	"fmt"
	"strings"
)

type Pokedex map[string]*Pokemon

func (dex Pokedex) String() string {
	var sb strings.Builder

	sb.WriteString("Your Pokedex:\n")

	for pokemonName := range dex {
		sb.WriteString(fmt.Sprintf("  - %s\n", pokemonName))
	}

	return sb.String()
}

type cliCommand struct {
	name        string
	description string
	callback    func(...string) error
}

type apiResp struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type LocationAreasResp struct {
	apiResp
	Results []*LocationArea `json:"results"`
}

type LocationArea struct {
	ID                int                 `json:"id"`
	Name              string              `json:"name"`
	PokemonEncounters []*PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon *Pokemon `json:"pokemon"`
}

type Pokemon struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	BaseExperience int            `json:"base_experience"`
	Height         int            `json:"height"`
	Weight         int            `json:"weight"`
	Stats          []*PokemonStat `json:"stats"`
	Types          []*PokemonType `json:"types"`
}

func (pokemon *Pokemon) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(`Name: %s
Height: %d
Weight: %d
`, pokemon.Name, pokemon.Height, pokemon.Weight))

	sb.WriteString("Stats:\n")
	for _, pStat := range pokemon.Stats {
		sb.WriteString(fmt.Sprintf("  -%s: %d\n", pStat.Stat.Name, pStat.BaseStat))
	}

	sb.WriteString("Types:\n")
	for _, pType := range pokemon.Types {
		sb.WriteString(fmt.Sprintf("  -%s\n", pType.Type.Name))
	}

	return sb.String()
}

type PokemonStat struct {
	Stat     *Stat `json:"stat"`
	BaseStat int   `json:"base_stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}
