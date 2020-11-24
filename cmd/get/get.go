package get

import (
	"github.com/restechnica/anyreleaser/cmd/get/version"
	"github.com/urfave/cli/v2"
)

const (
	command     = "get"
	description = "get information from the cli"
)

var (
	aliases = []string{"g"}
)

// NewCommand a command to get the current semver version.
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
