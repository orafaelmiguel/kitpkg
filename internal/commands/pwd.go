package commands

import (
	"fmt"
	"os"
)

type PwdCommand struct{}

func (p PwdCommand) Name() string {
	return "pwd"
}

func (p PwdCommand) Description() string {
	return "Prints current working directory"
}

func (p PwdCommand) Execute(input string, args []string) string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting directory:", err)
		return ""
	}

	return dir
}