package cmd

import (
	"errors"
	"fmt"
)

type Command struct{
	Name string
	Description string
	Function func() error
}

type CommandList map[string]Command

func InitializeCommands() (*CommandList, error) {
	cl := make(CommandList)

	Help := Command{
		Name: "help",
		Description: "Prints out all valid commands with descriptions for the pokedex",
	}

	Exit := Command{
		Name: "exit",
		Description: "Exits out of the program",
	}

	cl[Help.Name] = Help
	cl[Exit.Name] = Exit
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

	return nil
}
