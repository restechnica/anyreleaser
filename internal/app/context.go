package app

import (
	"github.com/restechnica/anyreleaser/internal/app/config"
	"github.com/restechnica/anyreleaser/internal/commands"
)

type Context struct {
	Commander commands.Commander
	Config    config.Root
}

func NewContext() *Context {
	return &Context{}
}
