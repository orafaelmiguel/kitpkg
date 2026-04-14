package commands

import (
	"fmt"
	"os"
)

type TouchCommand struct{}

func (t TouchCommand) Name() string {
	return "touch"
}

func (t TouchCommand) Description() string {
	return "Creates a file if it does not exist"
}

func (t TouchCommand) Execute(input string, args []string) string {
	if len(args) == 0 {
		fmt.Println("usage: touch <file>")
		return ""
	}

	filename := args[0]

	file, err := os.OpenFile(filename, os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("error creating file:", err)
		return ""
	}

	file.Close()

	return ""
}