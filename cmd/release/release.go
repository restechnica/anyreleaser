package release

import (
	"github.com/restechnica/anyreleaser/cmd/release/version"
	"github.com/urfave/cli/v2"
)

const (
	command     = "release"
	description = "makes releases"
)

var (
	aliases = []string{"r"}
)

// NewCommand a command to create releases.
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
