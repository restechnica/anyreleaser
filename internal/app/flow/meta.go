package flow

import (
	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/restechnica/anyreleaser/internal/commands"
)

type SetCommander struct{}

func (pipe SetCommander) Run(ctx *app.Context) (err error) {
	ctx.Commander = commands.NewExecCommander()
	return
}
