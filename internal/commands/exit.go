package commands

import "os"

type ExitCommand struct{}

func (c ExitCommand) Name() string {
	return "exit"
}

func (c ExitCommand) Description() string {
	return "Exit CLI."
}

func (c ExitCommand) Execute(input string, args []string) string {
	os.Exit(0)
	return ""
}