package git

import (
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

func action(context *cli.Context) (err error) {
	var pipeline = flow.Pipeline{}
	pipeline.Add(flow.SetGit{})
	return pipeline.Run(context)
}
