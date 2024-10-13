package pokeapi

import (
	"encoding/json"
	"fmt"
	"internal/pokecache"
	"io"
	"net/http"
	"sync"
)

type LocationArea struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type LocationAreaResponse struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []LocationArea `json:"results"`
}

type APIConfig struct {
	NextURL string
	PreviousURL string
	Cache *pokecache.Cache
	Mutex *sync.RWMutex
}

type Pokemon struct {
	Name string `json:"name"`
	PokemonURL string `json:"url"`
}

type Encounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type PokemonLocResponse struct {
	EncounterList []Encounter `json:"pokemon_encounters"`
	Name string `json:"name"`
	ID int `json:"id"`
}

func (cfg *APIConfig) GetNextLocations() ([]LocationArea, error) {

	var body []byte

	c, err := cfg.Cache.Get(cfg.NextURL, cfg.Mutex)
	if err != nil {
		return nil, fmt.Errorf("error checking cache: %w", err)
	}
	if c != nil {
		fmt.Println("Retrieving values from cache...")
		body = c
	} else {
		
		res, err := http.Get(cfg.NextURL)

		if err != nil {
			return nil, fmt.Errorf("error retrieving locations from PokeAPI: %w", err)
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error parsing PokeAPI response body: %w", err)
		}

		err = cfg.Cache.Add(cfg.NextURL, body, cfg.Mutex)
		if err != nil {
			return nil, fmt.Errorf("error caching location data: %w", err)
		}
	}

	lar := LocationAreaResponse{}
	err = json.Unmarshal(body, &lar)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling locations from PokeAPI: %w", err)
	}

	if cfg.PreviousURL == "" {
		cfg.PreviousURL = "https://pokeapi.co/api/v2/location-area/"
	} else {
		cfg.PreviousURL = lar.Previous
	}

	cfg.NextURL = lar.Next

	return lar.Results, nil
}

func (cfg *APIConfig) GetPreviousLocations() ([]LocationArea, error) {

	if cfg.PreviousURL == "" {
		fmt.Println("Reached start of map")
		return nil, fmt.Errorf("no previous locations available")
	}

	var body []byte

	c, err := cfg.Cache.Get(cfg.PreviousURL, cfg.Mutex)
	if err != nil {
		return nil, fmt.Errorf("error checking cache: %w", err)
	}
	if c != nil {
		fmt.Println("Retrieving values from cache...")
		body = c
	} else {
		res, err := http.Get(cfg.PreviousURL)
		if err != nil {
			return nil, fmt.Errorf("error retrieving locations from PokeAPI: %w", err)
		}
	
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("error parsing PokeAPI response body: %w", err)
		}
	}

	lar := LocationAreaResponse{}
	err = json.Unmarshal(body, &lar)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling locations from PokeAPI: %w", err)
	}

	cfg.NextURL = lar.Next
	cfg.PreviousURL = lar.Previous
	
	return lar.Results, nil
}

func (cfg *APIConfig) GetPokemonFromLocation(location string) ([]Pokemon, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + location

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error retrieving pokemon from PokeAPI: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}	

	encounters := PokemonLocResponse{}

	err = json.Unmarshal(body, &encounters)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}

	pl := make([]Pokemon, len(encounters.EncounterList))

	for k, v := range(encounters.EncounterList) {
		pl[k] = v.Pokemon
	}

	return pl, nil
}
