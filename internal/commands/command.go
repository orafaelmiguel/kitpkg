package commands

type Command interface {
	Name() string
	Description() string
	Execute(args []string)
}