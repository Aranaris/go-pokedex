package pokeapi

import (
	"encoding/json"
	"fmt"
	"internal/pokecache"
	"io"
	"net/http"
	"sync"
)

type Location struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type APIConfig struct {
	NextURL string
	PreviousURL string
	Cache *pokecache.Cache
	Mutex *sync.RWMutex
}

type LocationResponse struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []Location `json:"results"`
}


func (cfg *APIConfig) GetNextLocations() ([]Location, error) {

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
			return nil, fmt.Errorf("error cacheing location data: %w", err)
		}
	}

	locationResponse := LocationResponse{}
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling locations from PokeAPI: %w", err)
	}

	if cfg.PreviousURL == "" {
		cfg.PreviousURL = "https://pokeapi.co/api/v2/location"
	} else {
		cfg.PreviousURL = locationResponse.Previous
	}

	cfg.NextURL = locationResponse.Next
	
	return locationResponse.Results, nil
}

func (cfg *APIConfig) GetPreviousLocations() ([]Location, error) {

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

	locationResponse := LocationResponse{}
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling locations from PokeAPI: %w", err)
	}

	cfg.NextURL = locationResponse.Next
	cfg.PreviousURL = locationResponse.Previous
	
	return locationResponse.Results, nil
}
