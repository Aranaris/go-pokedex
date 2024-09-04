package cmd

import "fmt"

func HandleCommand(input string) func() {
	return func() {fmt.Printf("run command: %s", input)}
}
