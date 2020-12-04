package flow

import (
	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/urfave/cli/v2"
)

const AppContext = "app.context"

type PersistAppContext struct {
	CLIContext *cli.Context
}

func (pipe PersistAppContext) Run(ctx *app.Context) (err error) {
	pipe.CLIContext.App.Metadata[AppContext] = ctx
	return
}
