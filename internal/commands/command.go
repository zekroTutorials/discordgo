package commands

type Command interface {
	Invokes() []string
	Description() string
	AdminRequired() bool
	Exec(ctx *Context) error
}
