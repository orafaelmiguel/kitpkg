package commands

import (
	"strings"
)

type UpperCommand struct{}

func (u UpperCommand) Name() string {
	return "upper"
}

func (u UpperCommand) Description() string {
	return "Converts text to uppercase"
}

func (u UpperCommand) Execute(input string, args []string) string {
	var data string

	if input != "" {
		data = input
	} else {
		data = strings.Join(args, " ")
	}

	if data == "" {
		return ""
	}

	return strings.ToUpper(data)
}