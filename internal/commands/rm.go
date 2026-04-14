package commands

import (
	"fmt"
	"os"
)

type RmCommand struct{}

func (r RmCommand) Name() string {
	return "rm"
}

func (r RmCommand) Description() string {
	return "Removes a file"
}

func (r RmCommand) Execute(input string, args []string) string {
	if len(args) == 0 {
		fmt.Println("usage: rm <file>")
		return ""
	}

	filename := args[0]

	info, err := os.Stat(filename)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	if info.IsDir() {
		fmt.Println("cannot remove directory (use future -r flag)")
		return ""
	}

	err = os.Remove(filename)
	if err != nil {
		fmt.Println("error removing file:", err)
		return ""
	}

	return ""
}