package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

type MvCommand struct{}

func (m MvCommand) Name() string {
	return "mv"
}

func (m MvCommand) Description() string {
	return "Moves or renames files and directories"
}

func (m MvCommand) Execute(input string, args []string) string {
	if len(args) < 2 {
		fmt.Println("usage: mv <source> <destination>")
		return ""
	}

	src := args[0]
	dst := args[1]

	destInfo, err := os.Stat(dst)
	if err == nil && destInfo.IsDir() {
		dst = filepath.Join(dst, filepath.Base(src))
	}

	err = os.Rename(src, dst)
	if err != nil {
		fmt.Println("error moving file:", err)
		return ""
	}

	return ""
}