package commands

import (
	"fmt"
	"os"
)

type MkdirCommand struct{}

func (m MkdirCommand) Name() string {
	return "mkdir"
}

func (m MkdirCommand) Description() string {
	return "Creates a directory"
}

func (m MkdirCommand) Execute(input string, args []string) string {
	if len(args) == 0 {
		fmt.Println("usage: mkdir <directory>")
		return ""
	}

	dirname := args[0]

	err := os.Mkdir(dirname, 0755)
	if err != nil {
		fmt.Println("error creating directory:", err)
		return ""
	}

	return ""
}