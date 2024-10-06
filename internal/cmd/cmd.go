package cmd

import (
	"errors"
	"fmt"
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
		NextURL: "https://pokeapi.co/api/v2/location",
		PreviousURL: "",
		Cache: c,
		Mutex: &sync.RWMutex{},
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

	cl[Help.Name] = Help
	cl[Exit.Name] = Exit
	cl[Map.Name] = Map
	cl[MapB.Name] = MapB

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

func (cl *CommandList) HandleCommand(input string) error {
	_, ok := (*cl)[input]
	if !ok {
		return errors.New("Command not found: " + input)
	}
	
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

	return nil
}
