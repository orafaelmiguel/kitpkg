package commands

import (
	"strings"
)

type LowerCommand struct{}

func (l LowerCommand) Name() string {
	return "lower"
}

func (l LowerCommand) Description() string {
	return "Converts text to lowercase"
}

func (l LowerCommand) Execute(input string, args []string) string {
	var data string

	if input != "" {
		data = input
	} else {
		data = strings.Join(args, " ")
	}

	if data == "" {
		return ""
	}

	return strings.ToLower(data)
}