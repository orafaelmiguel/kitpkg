package commands

type Command interface {
	Name() string
	Description() string
	Execute(input string, args []string) string
}