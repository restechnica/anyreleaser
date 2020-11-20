package get

import (
	"github.com/urfave/cli/v2"

	"github.com/restechnica/anyreleaser/cmd/get/version"
)

// NewCommand a command to get the current semver version.
// Returns the CLI command.
func NewCommand(app *cli.App) *cli.Command {
	var command = "get"
	var description = "get information from the cli"
	var aliases = []string{"g"}

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
