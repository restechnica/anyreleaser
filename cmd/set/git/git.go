package git

import (
	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/restechnica/anyreleaser/internal/app/flow"
	"github.com/urfave/cli/v2"
)

const (
	command     = "git"
	description = "sets up git"
)

var aliases = []string{"g"}

// NewCommand a command to set up git.
// Returns the CLI command.
func NewCommand(app *cli.App) *cli.Command {
	return &cli.Command{
		Action:          action,
		Aliases:         aliases,
		HideHelp:        app.HideHelp,
		HideHelpCommand: app.HideHelpCommand,
		Name:            command,
		Usage:           description,
	}
}

func action(clictx *cli.Context) (err error) {
	var appctx = clictx.App.Metadata[flow.AppContext].(*app.Context)
	var pipeline = flow.Pipeline{}
	pipeline.Add(flow.SetGit{})
	return pipeline.Run(appctx)
}
