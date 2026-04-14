package commands

import (
	"fmt"
	"os"
)

type CdCommand struct{}

func (c CdCommand) Name() string {
	return "cd"
}

func (c CdCommand) Description() string {
	return "Changes current directory"
}

func (c CdCommand) Execute(input string, args []string) string {
	if len(args) == 0 {
		fmt.Println("usage: cd <directory>")
		return ""
	}

	err := os.Chdir(args[0])
	if err != nil {
		fmt.Println("error changing directory:", err)
	}

	return ""
}