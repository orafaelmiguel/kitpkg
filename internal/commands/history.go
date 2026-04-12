package commands

import "fmt"

type HistoryCommand struct {
	History *[]string
}

func (h HistoryCommand) Name() string {
	return "history"
}

func (h HistoryCommand) Description() string {
	return "Shows command history"
}

func (h HistoryCommand) Execute(args []string) {
	for i, cmd := range *h.History {
		fmt.Printf("%d: %s\n", i+1, cmd)
	}
}