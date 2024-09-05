/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"internal/cmd"
)

func printPrompt() {
	fmt.Print("pokedex", "> ")
}

func preProcessInput(input string) string {
	cleaned := strings.TrimSpace(input)
	cleaned = strings.ToLower(cleaned)
	return cleaned
}

func main() {
	cl, err := cmd.InitializeCommands()
	if err != nil {
		fmt.Printf("Error initializing commands: %s", err)
	}

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	
	exec:
	for reader.Scan() {
		text := preProcessInput(reader.Text())
		err := cl.HandleCommand(text)
		if err != nil {
			fmt.Printf("Error running command: %s, Error Out: %s", text, err)
		}
		fmt.Println("")

		if text == "exit" {
			break exec
		}
		printPrompt()
	}
}
