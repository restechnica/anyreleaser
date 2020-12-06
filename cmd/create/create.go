package create

import (
	"github.com/restechnica/anyreleaser/cmd/create/log"
	"github.com/urfave/cli/v2"
)

const (
	command     = "create"
	description = "create resources with the cli"
)

var (
	aliases = []string{"c"}
)

// NewCommand a command to create resources with the CLI.
// Returns the CLI command.
func NewCommand(app *cli.App) *cli.Command {
	var subcommands = []*cli.Command{
		log.NewCommand(app),
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
