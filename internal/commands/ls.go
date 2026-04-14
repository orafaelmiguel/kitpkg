package commands

import (
	"fmt"
	"os"
	"time"

	"kitpkg/internal/parser"
)

type LsCommand struct{}

func (l LsCommand) Name() string {
	return "ls"
}

func (l LsCommand) Description() string {
	return "Lists files in directory"
}

func (l LsCommand) Execute(input string, args []string) string {
	flags, params := parser.ParseFlags(args)

	showAll := flags["a"]
	longFormat := flags["l"]

	dir := "."

	if len(params) > 0 {
		dir = params[0]
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("error reading directory:", err)
		return ""
	}

	for _, file := range files {
		name := file.Name()

		if !showAll && len(name) > 0 && name[0] == '.' {
			continue
		}

		if longFormat {
			info, err := file.Info()
			if err != nil {
				continue
			}

			size := info.Size()
			modTime := info.ModTime().Format(time.RFC822)

			fmt.Printf("%-20s %10d %s\n", name, size, modTime)
		} else {
			fmt.Println(name)
		}
	}

	return ""
}