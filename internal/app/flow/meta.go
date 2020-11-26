package flow

import (
	"github.com/restechnica/anyreleaser/internal/commands"
	"github.com/urfave/cli/v2"
)

type SetCommander struct{}

func (pipe SetCommander) Run(ctx *cli.Context) (err error) {
	ctx.App.Metadata["commander"] = commands.NewExecCommander()
	return
}
