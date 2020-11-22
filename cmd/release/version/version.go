package version

import (
	"github.com/urfave/cli/v2"

	"github.com/restechnica/anyreleaser/internal/commands"
	"github.com/restechnica/anyreleaser/internal/git"
	"github.com/restechnica/anyreleaser/internal/semver"
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

func action(c *cli.Context) (err error) {
	var version string

	var strategyName = c.String("strategy")

	var commander = commands.NewExecCommander()
	var gitService = git.NewCLIService(commander)
	var semverManager = semver.NewManager(gitService)

	var strategy = semverManager.GetStrategy(strategyName)
	var tag = gitService.GetTag()

	if version, err = strategy.Increment(tag); err != nil {
		return
	}

	return gitService.CreateTag(version)
}
