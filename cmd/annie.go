package cmd

import (
	"github.com/restechnica/anyreleaser/cmd/get"
	"github.com/restechnica/anyreleaser/cmd/predict"
	"github.com/restechnica/anyreleaser/cmd/release"
	"github.com/restechnica/anyreleaser/cmd/set"
	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/urfave/cli/v2"
)

const (
	command     = "annie"
	description = "a CLI which serves as a backbone for your projects"
)

var (
	configFlag = &cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Usage:   "load configuration from a file",
		Value:   "./example/.anyreleaser.yaml",
	}

	flags = []cli.Flag{
		configFlag,
	}
)

// NewApp creates a new CLI app.
// Returns the cli app.
func NewApp() (app *cli.App) {
	app = &cli.App{
		Before:          before,
		Flags:           flags,
		HideHelp:        false,
		HideHelpCommand: true,
		Name:            command,
		Usage:           description,
	}

	app.Commands = []*cli.Command{
		get.NewCommand(app),
		predict.NewCommand(app),
		release.NewCommand(app),
		set.NewCommand(app),
	}

	return app
}

func before(context *cli.Context) (err error) {
	var pipeline = app.Pipeline{}

	// populate config
	pipeline.Add(app.DefaultConfigPipe{})
	pipeline.Add(app.ConfigPipe{})

	// populate commander
	pipeline.Add(app.CommanderPipe{})

	// set up env variables
	pipeline.Add(app.EnvScriptsPipe{})
	pipeline.Add(app.EnvFilesPipe{})
	pipeline.Add(app.EnvVarsPipe{})

	return pipeline.Run(context)
}
