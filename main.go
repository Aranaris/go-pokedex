/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func printPrompt() {
	fmt.Print("pokedex", "> ")
}

func main() {
	reader := bufio.NewScanner(os.Stdin)
	printPrompt()
	for reader.Scan() {
		text := reader.Text()
		fmt.Printf("input is: %s", text)
		fmt.Println()
		
		if text == "exit" {
			return
		}
	}

	fmt.Println()

}
