package cmd

import (
	"github.com/restechnica/anyreleaser/cmd/create"
	"github.com/restechnica/anyreleaser/cmd/get"
	"github.com/restechnica/anyreleaser/cmd/predict"
	"github.com/restechnica/anyreleaser/cmd/release"
	"github.com/restechnica/anyreleaser/cmd/set"
	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/restechnica/anyreleaser/internal/app/flow"
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
		Value:   ".anyreleaser.yaml",
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
		create.NewCommand(app),
		get.NewCommand(app),
		predict.NewCommand(app),
		release.NewCommand(app),
		set.NewCommand(app),
	}

	return app
}

func before(clictx *cli.Context) (err error) {
	var configPath = clictx.String("config")

	var appctx = app.NewContext()
	var pipeline = flow.Pipeline{}

	pipeline.Add(flow.SetCommander{})
	pipeline.Add(flow.LoadDefaultConfig{})
	pipeline.Add(flow.NewLoadConfig(configPath))
	pipeline.Add(flow.LoadEnvScripts{})
	pipeline.Add(flow.LoadEnvFiles{})
	pipeline.Add(flow.LoadEnvVars{})
	pipeline.Add(flow.PersistAppContext{CLIContext: clictx})

	return pipeline.Run(appctx)
}
