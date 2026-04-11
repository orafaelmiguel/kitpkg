package commands

import "os"

type ExitCommand struct{}

func (c ExitCommand) Name() string {
	return "exit"
}

func (c ExitCommand) Description() string {
	return "Exit CLI."
}

func (c ExitCommand) Execute(args []string) {
	os.Exit(0)
}