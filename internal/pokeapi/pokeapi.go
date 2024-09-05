package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

var POKEAPI_BASE_URL = "https://pokeapi.co/api/v2/"

func GetLocations() ([]Location, error) {
	res, err := http.Get(POKEAPI_BASE_URL + "location")
	if err != nil {
		return nil, fmt.Errorf("error retrieving locations from PokeAPI: %w", err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error parsing PokeAPI response body: %w", err)
	}

	fmt.Println(string(body))
	return nil, nil
}
