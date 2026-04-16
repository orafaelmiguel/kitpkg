package commands

import (
	"strings"
)

type ReplaceCommand struct{}

func (r ReplaceCommand) Name() string {
	return "replace"
}

func (r ReplaceCommand) Description() string {
	return "Replaces text: replace <old> <new>"
}

func (r ReplaceCommand) Execute(input string, args []string) string {
	if len(args) < 2 {
		return "usage: replace <old> <new>"
	}

	old := args[0]
	new := args[1]

	var data string

	if input != "" {
		data = input
	} else {
		data = strings.Join(args[2:], " ")
	}

	if data == "" {
		return "no input provided"
	}

	result := strings.ReplaceAll(data, old, new)

	return result
}