package cmd

import (
	"github.com/urfave/cli/v2"
)

// NewApp creates a new CLI app.
// Returns the cli app.
func NewApp() (app *cli.App) {
	var command = "bone"
	var description = "a CLI which serves as a backbone for your projects"

	app = &cli.App{
		HideHelp:        false,
		HideHelpCommand: true,
		Name:            command,
		Usage:           description,
	}

	app.Commands = []*cli.Command{}

	return
}
