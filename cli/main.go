package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"kitpkg/internal/commands"
	"kitpkg/internal/parser"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	commandMap := make(map[string]commands.Command)
	cmdList := []commands.Command{
		commands.EchoCommand{},
		commands.ExitCommand{},
	}

	for _, cmd := range cmdList {
		commandMap[cmd.Name()] = cmd
	}

	helpCmd := commands.HelpCommand{ Commands: commandMap }
	commandMap[helpCmd.Name()] = helpCmd

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR: ", err)
			continue
		}

		input = strings.TrimSpace(input)
		args := parser.Parse(input)

		if len(args) == 0 {
			continue
		}

		commandName := args[0]
		params := args[1:]

		if cmd, ok := commandMap[commandName]; ok {
			cmd.Execute(params)
		} else {
			fmt.Println("Command not recognized.")
		}
	}
}