package cmd

import (
	"fmt"

	"github.com/restechnica/backbone-cli/internal/commands"
	"github.com/restechnica/backbone-cli/internal/git"
	"github.com/restechnica/backbone-cli/internal/semver"
	"github.com/urfave/cli/v2"
)

// NewVersionCommand a command for semver version management based on git tags.
// Returns the CLI command.
func NewVersionCommand(app *cli.App) *cli.Command {
	var command = "version"
	var description = "semver version operations based on config and/or git tags"
	var aliases = []string{"s"}

	var subcommands = []*cli.Command{
		NewVersionGetCommand(app),
		NewVersionGetNextCommand(app),
		NewVersionUpCommand(app),
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

// NewVersionGetCommand a command to get the current semver version.
// Returns the CLI command.
func NewVersionGetCommand(app *cli.App) *cli.Command {
	var command = "get"
	var description = "gets the current semver version based on a strategy"
	var aliases = []string{"g"}

	var action = func(c *cli.Context) (err error) {
		var commander = commands.NewExecCommander()
		var gitService = git.NewCLIService(commander)

		fmt.Println(gitService.GetTag())

		return
	}

	return &cli.Command{
		Action:          action,
		Aliases:         aliases,
		HideHelp:        app.HideHelp,
		HideHelpCommand: app.HideHelpCommand,
		Name:            command,
		Usage:           description,
	}
}

// NewVersionGetNextCommand a command to get the next semver version.
// Returns the CLI command.
func NewVersionGetNextCommand(app *cli.App) *cli.Command {
	var command = "get-next"
	var description = "gets the next semver version"
	var aliases = []string{"gn"}

	var flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "strategy",
			Aliases: []string{"s"},
			Usage:   "determines the semver level to increment",
			Value:   "auto",
		},
	}

	var action = func(c *cli.Context) (err error) {
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

		fmt.Println(version)

		return
	}

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

// NewVersionUpCommand a command to increment the current semver version.
// The [strategy|s] flag allows you to choose which semver level to increment.
// Returns the CLI command.
func NewVersionUpCommand(app *cli.App) *cli.Command {
	var command = "up"
	var description = "increments the semver version based on a strategy"
	var aliases = []string{"u"}

	var flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "strategy",
			Aliases: []string{"s"},
			Usage:   "determines the semver level to increment",
			Value:   "auto",
		},
	}

	var action = func(c *cli.Context) (err error) {
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
