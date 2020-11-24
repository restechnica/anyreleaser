package version

import (
	"fmt"

	"github.com/restechnica/anyreleaser/internal/commands"
	"github.com/restechnica/anyreleaser/internal/config"
	"github.com/restechnica/anyreleaser/internal/git"
	"github.com/restechnica/anyreleaser/internal/semver"
	"github.com/urfave/cli/v2"
)

const (
	command     = "version"
	description = "predicts the next semver version"
)

var (
	aliases = []string{"v"}

	strategyFlag = &cli.StringFlag{
		Name:    "strategy",
		Aliases: []string{"s"},
		Usage:   "determines the semver level to increment",
		Value:   "auto",
	}

	flags = []cli.Flag{
		strategyFlag,
	}
)

// NewCommand a command to predict the next semver version.
//// Returns the CLI command.
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

func action(context *cli.Context) (err error) {
	var version string

	var config = context.App.Metadata["config"].(config.Root)

	if context.IsSet("strategy") {
		config.Semver.Strategy = context.String("strategy")
	}

	var commander = commands.NewExecCommander()
	var gitService = git.NewCLIService(commander)
	var semverManager = semver.NewManager(config, gitService)

	var strategy = semverManager.GetStrategy(config.Semver.Strategy)
	var tag = gitService.GetTag()

	if version, err = strategy.Increment(tag); err != nil {
		return
	}

	fmt.Println(version)

	return
}
