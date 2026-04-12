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

	helpCmd := commands.HelpCommand{
		Commands: commandMap,
	}
	commandMap[helpCmd.Name()] = helpCmd

	aliases := make(map[string]string)

	aliasCmd := commands.AliasCommand{
		Aliases: aliases,
	}
	commandMap[aliasCmd.Name()] = aliasCmd

	unaliasCmd := commands.UnaliasCommand{
		Aliases: aliases,
	}
	commandMap[unaliasCmd.Name()] = unaliasCmd

	history := []string{}

	historyCmd := commands.HistoryCommand{
		History: &history,
	}
	commandMap[historyCmd.Name()] = historyCmd

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}

		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if input != "history" {
			history = append(history, input)
		}

		parts := strings.Split(input, "|")

		var result string

		for _, part := range parts {
			part = strings.TrimSpace(part)

			args := parser.Parse(part)
			if len(args) == 0 {
				continue
			}

			commandName := args[0]
			params := args[1:]

			if realCmd, exists := aliases[commandName]; exists {
				commandName = realCmd
			}

			cmd, ok := commandMap[commandName]
			if !ok {
				fmt.Println("Command not recognized:", commandName)
				result = ""
				break
			}

			result = cmd.Execute(result, params)
		}

		if result != "" {
			fmt.Println(result)
		}
	}
}