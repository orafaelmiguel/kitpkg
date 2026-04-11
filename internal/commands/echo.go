package commands

import (
	"fmt"
	"strings"
)

type EchoCommand struct{}

func (c EchoCommand) Name() string {
	return "echo"
}

func (c EchoCommand) Description() string {
	return "Prints text in the terminal."
}

func (c EchoCommand) Execute(args []string) {
	fmt.Println(strings.Join(args, " "))
}