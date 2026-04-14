package commands

import (
	"fmt"
	"os"

	"kitpkg/internal/parser"
)

type RmCommand struct{}

func (r RmCommand) Name() string {
	return "rm"
}

func (r RmCommand) Description() string {
	return "Removes files or directories"
}

func (r RmCommand) Execute(input string, args []string) string {
	flags, params := parser.ParseFlags(args)

	recursive := flags["r"]
	force := flags["f"]

	if len(params) == 0 {
		fmt.Println("usage: rm [-r] [-f] <target>")
		return ""
	}

	target := params[0]

	info, err := os.Stat(target)
	if err != nil {
		if !force {
			fmt.Println("error:", err)
		}
		return ""
	}

	if info.IsDir() {
		if !recursive {
			fmt.Println("cannot remove directory (use -r)")
			return ""
		}

		err = os.RemoveAll(target)
		if err != nil && !force {
			fmt.Println("error removing directory:", err)
		}
		return ""
	}

	err = os.Remove(target)
	if err != nil && !force {
		fmt.Println("error removing file:", err)
	}

	return ""
}