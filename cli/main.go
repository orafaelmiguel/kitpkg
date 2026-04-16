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

func redraw(input string, cursorPos int) {
	fmt.Print("\r\033[K")
	fmt.Print(shell.GetPrompt() + input)

	moveBack := len(input) - cursorPos
	if moveBack > 0 {
		fmt.Printf("\033[%dD", moveBack)
	}
}

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
		commands.CpCommand{},
		commands.MvCommand{},
		commands.ReplaceCommand{},
		commands.UpperCommand{},
		commands.LowerCommand{},
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
		cursorPos := 0

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
				if cursorPos > 0 {
					input = input[:cursorPos-1] + input[cursorPos:]
					cursorPos--
					redraw(input, cursorPos)
				}
			case '\t':
				input = shell.HandleTab(input, commandMap)
				cursorPos = len(input)
				redraw(input, cursorPos)
			case 27:
				next1, _ := reader.ReadByte()
				next2, _ := reader.ReadByte()

				if next1 == 91 {
					switch next2 {
					case 65:
						if len(history) > 0 && historyIndex < len(history)-1 {
							historyIndex++
							input = history[len(history)-1-historyIndex]
							cursorPos = len(input)
						}
					case 66:
						if historyIndex > 0 {
							historyIndex--
							input = history[len(history)-1-historyIndex]
						} else {
							historyIndex = -1
							input = ""
						}
						cursorPos = len(input)
					case 67:
						if cursorPos < len(input) {
							cursorPos++
						}
					case 68:
						if cursorPos > 0 {
							cursorPos--
						}
					}

					redraw(input, cursorPos)
				}
			default:
				historyIndex = -1
				input = input[:cursorPos] + string(char) + input[cursorPos:]
				cursorPos++
				redraw(input, cursorPos)
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

		redir := parser.ParseRedirection(input)
		parts := strings.Split(redir.Command, "|")

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

		if redir.File != "" {
			var file *os.File
			var err error

			if redir.Append {
				file, err = os.OpenFile(redir.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			} else {
				file, err = os.Create(redir.File)
			}

			if err != nil {
				fmt.Println("error writing file:", err)
				return
			}
			defer file.Close()

			if result != "" {
				file.WriteString(result + "\n")
			}
		} else {
			if result != "" {
				fmt.Println(result)
			}
		}
	}
}