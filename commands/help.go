package commands

import (
	"fmt"
)

func help(args ...string) error {
	fmt.Println("Usage: filecust <command> [arguments]")
	fmt.Println("The commands are:")
	fmt.Println()

	maxLen := 0

	for _, command := range GetCommands() {
		if len(command.Name) > maxLen {
			maxLen = len(command.Name)
		}
	}
	for _, command := range GetCommands() {
		fmt.Printf("%-*s   %s\n", maxLen, command.Name, command.Description)
	}

	return nil

}
