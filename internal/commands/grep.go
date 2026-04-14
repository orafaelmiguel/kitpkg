package commands

import (
	"fmt"
	"strings"
)

type GrepCommand struct{}

func (g GrepCommand) Name() string {
	return "grep"
}

func (g GrepCommand) Description() string {
	return "Filters lines containing a pattern"
}

func (g GrepCommand) Execute(input string, args []string) string {
	if len(args) == 0 {
		fmt.Println("usage: grep <pattern>")
		return ""
	}

	pattern := args[0]

	if input == "" {
		fmt.Println("no input provided")
		return ""
	}

	lines := strings.Split(input, "\n")
	var result []string

	for _, line := range lines {
		if strings.Contains(line, pattern) {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}