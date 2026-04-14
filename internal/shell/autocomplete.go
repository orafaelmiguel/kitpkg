package shell

import (
	"os"
	"path/filepath"
	"strings"

	"kitpkg/internal/commands"
)

func HandleTab(input string, commandMap map[string]commands.Command) string {
	args := strings.Fields(input)

	if len(args) == 0 {
		return input
	}

	if len(args) == 1 && !strings.Contains(input, " ") {
		prefix := args[0]

		var matches []string
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

	lastArg := args[len(args)-1]

	dir := "."
	prefix := lastArg

	if strings.Contains(lastArg, "/") {
		dir = filepath.Dir(lastArg)
		prefix = filepath.Base(lastArg)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return input
	}

	var matches []string

	for _, entry := range entries {
		name := entry.Name()

		if strings.HasPrefix(strings.ToLower(name), strings.ToLower(prefix)) {
			if entry.IsDir() {
				name += "/"
			}

			if dir != "." {
				name = filepath.Join(dir, name)
			}

			matches = append(matches, name)
		}
	}

	if len(matches) == 1 {
		args[len(args)-1] = matches[0]
		return strings.Join(args, " ")
	}

	return input
}