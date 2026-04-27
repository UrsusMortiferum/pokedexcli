package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	shallowLocations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = shallowLocations.Next
	cfg.prevLocationsURL = shallowLocations.Previous

	for _, result := range shallowLocations.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	shallowLocations, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = shallowLocations.Next
	cfg.prevLocationsURL = shallowLocations.Previous

	for _, result := range shallowLocations.Results {
		fmt.Println(result.Name)
	}
	return nil
}
