package cmd

import "fmt"

type Command struct{
	Name string
	Description string
}

type CommandList map[string]Command

func InitializeCommands() (*CommandList, error) {
	cl := make(CommandList)
	var Help = new(Command)
	Help.Name = "help"
	Help.Description = "Prints out all valid commands with descriptions for the pokedex"

	cl[Help.Name] = *Help
	return &cl, nil
}

func (cl *CommandList) HandleCommand(input string) (func(), error) {
	return func() {fmt.Printf("run command: %s", input)}, nil
}
