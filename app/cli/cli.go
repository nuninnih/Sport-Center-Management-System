package cli

import "github.com/nuninnih/Sport-Center-Management-System/handler"

type CLI struct {
	Handler *handler.Handler
}

func NewCLI(handler *handler.Handler) *CLI {
	return &CLI{Handler: handler}
}

func (c *CLI) Run() {}
