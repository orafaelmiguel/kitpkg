package commands

import "fmt"

type HelpCommand struct {
	Commands map[string]Command
}

func (h HelpCommand) Name() string {
	return "help"
}

func (h HelpCommand) Description() string {
	return "Lists all avaliable commands"
}

func (h HelpCommand) Execute(args []string) {
	for name, cmd := range h.Commands {
		fmt.Printf("%s - %s\n", name, cmd.Description())
	}
}