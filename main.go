package main

import (
	"fmt"
	"os"

	"github.com/nick5928/file_custodian/commands"
)

func main() {
	cmds := commands.GetCommands()
	if len(os.Args) < 2 {
		fmt.Println("Please give a command")
		cmds["help"].Callback()
		return
	}
	arg := os.Args[1]
	cmd, ok := cmds[arg]
	if !ok {
		fmt.Println("Command not found")
		cmds["help"].Callback()
		return
	}
	var params []string

	if len(os.Args) > 2 {
		params = os.Args[2:]
	}
	err := cmd.Callback(params...)

	if err != nil {
		fmt.Println(err)
	}

}
