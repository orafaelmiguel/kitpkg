package commands

import (
	"fmt"
	"os"
)

type CatCommand struct{}

func (c CatCommand) Name() string {
	return "cat"
}

func (c CatCommand) Description() string {
	return "Reads and outputs file content"
}

func (c CatCommand) Execute(input string, args []string) string {
	if input != "" && len(args) == 0 {
		return input
	}

	if len(args) == 0 {
		fmt.Println("usage: cat <file>")
		return ""
	}

	filename := args[0]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading file:", err)
		return ""
	}

	return string(data)
}