package commands

import "fmt"

type AliasCommand struct {
	Aliases map[string]string
}

func (a AliasCommand) Name() string {
	return "alias"
}

func (a AliasCommand) Description() string {
	return "Creates a command alias"
}

func (a AliasCommand) Execute(args []string) {
	if len(args) == 0 {
		for k, v := range a.Aliases {
			fmt.Printf("%s -> %s\n", k, v)
		}
		return
	}

	if len(args) == 1 {
		name := args[0]
		if val, ok := a.Aliases[name]; ok {
			fmt.Printf("%s -> %s\n", name, val)
		} else {
			fmt.Println("alias not found")
		}
		return
	}

	name := args[0]
	command := args[1]

	a.Aliases[name] = command

	fmt.Printf("Alias '%s' -> '%s'\n", name, command)
}

// todo: load aliases in external file to load in startup 