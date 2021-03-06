package version

import (
	"fmt"

	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/restechnica/anyreleaser/internal/app/flow"
	"github.com/restechnica/anyreleaser/internal/git"
	"github.com/restechnica/anyreleaser/internal/semver"
	"github.com/urfave/cli/v2"
)

const (
	command     = "version"
	description = "increments the semver version based on a strategy"
)

var (
	aliases = []string{"v"}

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "strategy",
			Aliases: []string{"s"},
			Usage:   "determines the semver level to increment",
			Value:   "auto",
		},
	}
)

// NewCommand a command to increment the current semver version.
// The [strategy|s] flag allows you to choose which semver level to increment.
// Returns the CLI command.
func NewCommand(app *cli.App) *cli.Command {
	return &cli.Command{
		Action:          action,
		Aliases:         aliases,
		Flags:           flags,
		HideHelp:        app.HideHelp,
		HideHelpCommand: app.HideHelpCommand,
		Name:            command,
		Usage:           description,
	}
}

func action(clictx *cli.Context) (err error) {
	var version string

	var appctx, ok = clictx.App.Metadata[flow.AppContext].(*app.Context)

	if !ok {
		return fmt.Errorf("something went wrong with fetching the app context from the cli context")
	}

	if clictx.IsSet("strategy") {
		appctx.Config.Semver.Strategy = clictx.String("strategy")
	}

	var gitService = git.NewCLIService(appctx.Commander)
	var semverManager = semver.NewManager(appctx.Config, gitService)

	var strategy = semverManager.GetStrategy(appctx.Config.Semver.Strategy)
	var tag = gitService.GetTag()

	if version, err = strategy.Increment(tag); err != nil {
		return
	}

	return gitService.CreateTag(version)
}
