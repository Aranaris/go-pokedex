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
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := preProcessInput(reader.Text())
		f := cmd.HandleCommand(text)
		f()
		fmt.Println()

		if text == "exit" {
			return
		}
		printPrompt()
	}

	fmt.Println()

}
