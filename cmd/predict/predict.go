package predict

import (
	"github.com/restechnica/anyreleaser/cmd/predict/version"
	"github.com/urfave/cli/v2"
)

const (
	command     = "predict"
	description = "get future information from the cli"
)

var (
	aliases = []string{"p"}
)

// NewCommand a command to get future information from the cli.
// Returns the CLI command.
func NewCommand(app *cli.App) *cli.Command {
	var subcommands = []*cli.Command{
		version.NewCommand(app),
	}

	return &cli.Command{
		Aliases:         aliases,
		HideHelp:        app.HideHelp,
		HideHelpCommand: app.HideHelpCommand,
		Name:            command,
		Subcommands:     subcommands,
		Usage:           description,
	}
}
