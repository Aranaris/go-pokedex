module github.com/Aranaris/go-pokedex

go 1.22.5

require internal/cmd v1.0.0

replace internal/cmd => ./internal/cmd

require internal/pokeapi v1.0.0 // indirect

replace internal/pokeapi => ./internal/pokeapi

require internal/pokecache v1.0.0 // indirect

replace internal/pokecache => ./internal/pokecache
