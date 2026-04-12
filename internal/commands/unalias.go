package commands

import "fmt"

type UnaliasCommand struct {
	Aliases map[string]string
}

func (u UnaliasCommand) Name() string {
	return "unalias"
}

func (u UnaliasCommand) Description() string {
	return "Removes an alias"
}

func (u UnaliasCommand) Execute(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: unalias <name>")
		return
	}

	name := args[0]

	if _, exists := u.Aliases[name]; exists {
		delete(u.Aliases, name)
		fmt.Printf("Alias '%s' removed\n", name)
	} else {
		fmt.Println("Alias not found")
	}
}