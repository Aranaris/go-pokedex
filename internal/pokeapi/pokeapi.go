package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type APIConfig struct {
	NextURL string
	PreviousURL string
}

type LocationResponse struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []Location `json:"results"`
}


func (cfg *APIConfig) GetLocations() ([]Location, error) {
	res, err := http.Get(cfg.NextURL)
	if err != nil {
		return nil, fmt.Errorf("error retrieving locations from PokeAPI: %w", err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error parsing PokeAPI response body: %w", err)
	}

	locationResponse := LocationResponse{}
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling locations from PokeAPI: %w", err)
	}

	if cfg.PreviousURL == "" {
		cfg.PreviousURL = cfg.NextURL
	} else {
		cfg.PreviousURL = locationResponse.Previous
	}
	
	cfg.NextURL = locationResponse.Next

	return locationResponse.Results, nil
}
