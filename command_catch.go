package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	baseExperience := pokemon.BaseExperience

	// Calculate the catch probability (higher baseExperience = lower chance)
	catchProbability := 1 / (1 + float64(baseExperience)/100)

	// Generate a random number between 0 and 1
	randomNumber := rand.Float64()
	catch := false
	// Compare random number with catch probability
	catch = randomNumber < catchProbability

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	if !catch {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil
}
