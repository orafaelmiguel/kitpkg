package commands

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"kitpkg/internal/parser"
)

type CpCommand struct{}

func (c CpCommand) Name() string {
	return "cp"
}

func (c CpCommand) Description() string {
	return "Copies files and directories"
}

func (c CpCommand) Execute(input string, args []string) string {
	flags, params := parser.ParseFlags(args)

	recursive := flags["r"]

	if len(params) < 2 {
		fmt.Println("usage: cp [-r] <source> <destination>")
		return ""
	}

	src := params[0]
	dst := params[1]

	info, err := os.Stat(src)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	if info.IsDir() {
		if !recursive {
			fmt.Println("cp: omitting directory (use -r)")
			return ""
		}
		return copyDir(src, dst)
	}

	return copyFile(src, dst)
}

func copyFile(src, dst string) string {
	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Println("error opening source:", err)
		return ""
	}
	defer sourceFile.Close()

	destInfo, err := os.Stat(dst)
	if err == nil && destInfo.IsDir() {
		dst = filepath.Join(dst, filepath.Base(src))
	}

	destFile, err := os.Create(dst)
	if err != nil {
		fmt.Println("error creating destination:", err)
		return ""
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Println("error copying:", err)
		return ""
	}

	return ""
}

func copyDir(src, dst string) string {
	err := os.MkdirAll(dst, 0755)
	if err != nil {
		fmt.Println("error creating directory:", err)
		return ""
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		fmt.Println("error reading directory:", err)
		return ""
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		info, err := entry.Info()
		if err != nil {
			continue
		}

		if info.IsDir() {
			copyDir(srcPath, dstPath)
		} else {
			copyFile(srcPath, dstPath)
		}
	}

	return ""
}