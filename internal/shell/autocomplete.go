package shell

import (
	"strings"

	"kitpkg/internal/commands"
)

func HandleTab(input string, commandMap map[string]commands.Command) string {
	args := strings.Fields(input)

	if len(args) == 0 {
		return input
	}

	if len(args) == 1 {
		prefix := args[0]

		matches := []string{}
		for name := range commandMap {
			if strings.HasPrefix(name, prefix) {
				matches = append(matches, name)
			}
		}

		if len(matches) == 1 {
			return matches[0] + " "
		}

		return input
	}

	return input
}