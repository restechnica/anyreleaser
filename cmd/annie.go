package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/restechnica/anyreleaser/cmd/get"
	_init "github.com/restechnica/anyreleaser/cmd/init"
	"github.com/restechnica/anyreleaser/cmd/predict"
	"github.com/restechnica/anyreleaser/cmd/release"
)

// NewApp creates a new CLI app.
// Returns the cli app.
func NewApp() (app *cli.App) {
	var command = "annie"
	var description = "a CLI which serves as a backbone for your projects"

	app = &cli.App{
		HideHelp:        false,
		HideHelpCommand: true,
		Name:            command,
		Usage:           description,
	}

	app.Commands = []*cli.Command{
		get.NewCommand(app),
		_init.NewCommand(app),
		predict.NewCommand(app),
		release.NewCommand(app),
	}

	return
}
