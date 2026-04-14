package shell

import (
	"os"
	"strings"
)

const (
	Blue  = "\033[34m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func formatPath(path string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return path
	}

	if strings.HasPrefix(path, home) {
		return "~" + strings.TrimPrefix(path, home)
	}

	return path
}

func GetPrompt() string {
	cwd, err := os.Getwd()
	if err != nil {
		return "$ "
	}

	formatted := formatPath(cwd)

	return Blue + formatted + Reset + "" + Green + ">" + Reset + " "
}