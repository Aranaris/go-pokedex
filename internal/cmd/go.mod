module github.com/Aranaris/go-pokedex

go 1.22.5

require internal/pokeapi v1.0.0

replace internal/pokeapi => ../pokeapi

require internal/pokecache v1.0.0

replace internal/pokecache => ../pokecache
