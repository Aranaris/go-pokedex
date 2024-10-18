package cmd

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"internal/pokeapi"
	"internal/pokecache"
)

type Command struct{
	Name string
	Description string
	Config *pokeapi.APIConfig
}

type CommandList map[string]Command

func InitializeCommands() (*CommandList, error) {
	cl := make(CommandList)
	c, _ := pokecache.NewCache(time.Duration(10 * time.Second))

	cfg := pokeapi.APIConfig{
		NextURL: "https://pokeapi.co/api/v2/location-area/",
		PreviousURL: "",
		Cache: c,
		Mutex: &sync.RWMutex{},
		Pokedex: map[string]*pokeapi.Pokemon{},
	}

	Help := Command{
		Name: "help",
		Description: "Prints out all valid commands with descriptions for the pokedex",
	}

	Exit := Command{
		Name: "exit",
		Description: "Exits out of the program",
	}

	Map := Command{
		Name: "map",
		Description: "Displays the next 20 locations of the pokemon map",
		Config: &cfg,
	}

	MapB := Command{
		Name: "mapb",
		Description: "Displays the previous 20 locations of the pokemon map",
		Config: &cfg,
	}

	Explore := Command{
		Name: "explore",
		Description: "Displays a list of all pokemon in a given location area",
		Config: &cfg,
	}

	Catch := Command{
		Name: "catch",
		Description: "Attempt to catch a pokemon",
		Config: &cfg,
	}

	Inspect := Command{
		Name: "inspect",
		Description: "Get pokemon info from pokedex",
		Config: &cfg,
	}

	cl[Help.Name] = Help
	cl[Exit.Name] = Exit
	cl[Map.Name] = Map
	cl[MapB.Name] = MapB
	cl[Explore.Name] = Explore
	cl[Catch.Name] = Catch
	cl[Inspect.Name] = Inspect

	return &cl, nil
}

func (cl *CommandList) CommandHelp() error {
	fmt.Println("List of all valid commands:")
	for _, v := range *cl {
		fmt.Printf("Command: %s || Description: %s", v.Name, v.Description)
		fmt.Println("")
	}
	return nil
}

func (cl *CommandList) CommandExit() error {
	fmt.Println("Closing...")
	return nil
}

func (cl *CommandList) CommandMap() error {
	fmt.Println("Showing locations...")
	cfg := (*cl)["map"].Config
	locations, err := cfg.GetNextLocations()
	if err != nil {
		return err
	}
	
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func (cl *CommandList) CommandMapB() error {
	fmt.Println("Showing locations...")
	cfg := (*cl)["mapb"].Config
	locations, err := cfg.GetPreviousLocations()
	if err != nil {
		return err
	}
	
	for _, location := range locations {
		fmt.Println(location.Name)
	}
	return nil
}

func (cl *CommandList) CommandExplore(location string) error {
	fmt.Printf("Showing pokemon at location %s: ", location)
	cfg := (*cl)["explore"].Config
	pokemonlist, err := cfg.GetPokemonFromLocation(location)
	if err != nil {
		fmt.Printf("error retrieving pokemon list: %s", err)
		return err
	}

	for _, pokemon := range pokemonlist {
		fmt.Println(pokemon.Name)
	}

	return nil
}

func (cl *CommandList) CommandCatch(pokemon string) error {
	fmt.Printf("Attempting to catch %s...", pokemon)
	fmt.Println("")
	cfg := (*cl)["catch"].Config
	c, err := cfg.CatchPokemon(pokemon)
	if err != nil {
		fmt.Printf("error catching pokemon: %s", err)
	}

	if c {
		fmt.Printf("%s was caught!", pokemon)
	} else {
		fmt.Printf("%s escaped!", pokemon)
	}
	
	return nil
}

func (cl *CommandList) CommandInspect(pokemon string) error {
	fmt.Printf("Getting %s info from pokedex...", pokemon)
	fmt.Println("")
	cfg := (*cl)["inspect"].Config
	pd, err := cfg.InspectPokemon(pokemon)
	if err != nil {
		fmt.Printf("error getting pokemon info: %s", err)
	} else {
		fmt.Println(pd)
	}
	
	return nil
}

func (cl *CommandList) HandleCommand(input string) error {
	if input == "help" {
		cl.CommandHelp()
		return nil
	}

	if input == "exit" {
		cl.CommandExit()
		return nil
	}

	if input == "map" {
		cl.CommandMap()
		return nil
	}

	if input == "mapb" {
		cl.CommandMapB()
		return nil
	}
	
	inputs := strings.Split(input, " ")

	if inputs[0] == "explore" {
		if len(inputs) <= 1 {
			return errors.New("no location provided to explore")
		}
		cl.CommandExplore(inputs[1])
		return nil
	}

	if inputs[0] == "catch" {
		if len(inputs) <= 1 {
			return errors.New("no pokemon name provided to catch")
		}
		cl.CommandCatch(inputs[1])
		return nil
	}

	if inputs[0] == "inspect" {
		if len(inputs) <= 1 {
			return errors.New("no pokemon name provided to inspect")
		}
		cl.CommandInspect(inputs[1])
		return nil
	}
	return errors.New("Command not found: " + input)
}
