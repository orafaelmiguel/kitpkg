package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"

	"kitpkg/internal/commands"
	"kitpkg/internal/parser"
	"kitpkg/internal/shell"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	oldState, err := term.MakeRaw(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(syscall.Stdin), oldState)

	commandMap := make(map[string]commands.Command)

	cmdList := []commands.Command{
		commands.EchoCommand{},
		commands.ExitCommand{},
		commands.CatCommand{},
		commands.GrepCommand{},
		commands.LsCommand{},
		commands.PwdCommand{},
		commands.CdCommand{},
		commands.TouchCommand{},
		commands.MkdirCommand{},
		commands.RmCommand{},
	}

	for _, cmd := range cmdList {
		commandMap[cmd.Name()] = cmd
	}

	helpCmd := commands.HelpCommand{
		Commands: commandMap,
	}
	commandMap[helpCmd.Name()] = helpCmd

	aliases := make(map[string]string)

	commandMap["alias"] = commands.AliasCommand{Aliases: aliases}
	commandMap["unalias"] = commands.UnaliasCommand{Aliases: aliases}

	history := []string{}
	commandMap["history"] = commands.HistoryCommand{History: &history}

	historyIndex := -1 

	for {
		fmt.Print(shell.GetPrompt())

		var input string

		for {
			char, err := reader.ReadByte()
			if err != nil {
				fmt.Println("\nERROR:", err)
				break
			}

			switch char {
			case '\r', '\n':
				fmt.Println()
				goto EXECUTE
			case 127:
				if len(input) > 0 {
					input = input[:len(input)-1]
					fmt.Print("\b \b")
				}
			case '\t':
				input = shell.HandleTab(input, commandMap)

				fmt.Print("\r\033[K")
				cwd, _ := os.Getwd()
				fmt.Printf("%s> %s", cwd, input)
			case 27:
				next1, _ := reader.ReadByte()
				next2, _ := reader.ReadByte()

				if next1 == 91 {
					switch next2 {
					case 65:
						if len(history) > 0 && historyIndex < len(history)-1 {
							historyIndex++
							input = history[len(history)-1-historyIndex]
						}
					case 66:
						if historyIndex > 0 {
							historyIndex--
							input = history[len(history)-1-historyIndex]
						} else {
							historyIndex = -1
							input = ""
						}
					}
					fmt.Print("\r\033[K")
					cwd, _ := os.Getwd()
					fmt.Printf("%s> %s", cwd, input)
				}

			default:
				historyIndex = -1 
				input += string(char)
				fmt.Print(string(char))
			}
		}

	EXECUTE:

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