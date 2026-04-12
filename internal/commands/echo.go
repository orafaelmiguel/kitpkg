package commands

import (
	"strings"
)

type EchoCommand struct{}

func (c EchoCommand) Name() string {
	return "echo"
}

func (c EchoCommand) Description() string {
	return "Prints text in the terminal."
}

func (c EchoCommand) Execute(input string, args []string) string {
	return strings.Join(args, " ")
}