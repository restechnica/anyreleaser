package set

import (
	"github.com/restechnica/anyreleaser/cmd/set/git"
	"github.com/urfave/cli/v2"
)

const (
	command     = "set"
	description = "set information"
)

var aliases = []string{"s"}

// NewCommand a command to set information with the CLI.
// Returns the CLI command.
func NewCommand(app *cli.App) *cli.Command {
	var subcommands = []*cli.Command{
		git.NewCommand(app),
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
