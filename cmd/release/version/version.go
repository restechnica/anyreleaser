package version

import (
	"github.com/restechnica/anyreleaser/internal/commands"
	"github.com/restechnica/anyreleaser/internal/config"
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

func action(context *cli.Context) (err error) {
	var version string

	var cfg = context.App.Metadata["cfg"].(config.Root)

	if context.IsSet("strategy") {
		cfg.Semver.Strategy = context.String("strategy")
	}

	var commander = commands.NewExecCommander()
	var gitService = git.NewCLIService(commander)
	var semverManager = semver.NewManager(cfg, gitService)

	var strategy = semverManager.GetStrategy(cfg.Semver.Strategy)
	var tag = gitService.GetTag()

	if version, err = strategy.Increment(tag); err != nil {
		return
	}

	return gitService.CreateTag(version)
}
