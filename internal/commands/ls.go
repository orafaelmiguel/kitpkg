package commands

import (
	"fmt"
	"os"
)

type LsCommand struct{}

func (l LsCommand) Name() string {
	return "ls"
}

func (l LsCommand) Description() string {
	return "Lists directory contents"
}

func (l LsCommand) Execute(input string, args []string) string {
	path := "."

	if len(args) > 0 {
		path = args[0]
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("error reading directory:", err)
		return ""
	}

	var result string

	for _, entry := range entries {
		name := entry.Name()

		if entry.IsDir() {
			name += "/"
		}

		result += name + "\n"
	}

	return result
}