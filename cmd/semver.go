package cmd

import (
	"fmt"

	"github.com/restechnica/backbone-cli/internal/commands"
	"github.com/restechnica/backbone-cli/internal/git"
	"github.com/restechnica/backbone-cli/internal/semver"
	"github.com/urfave/cli/v2"
)

// NewSemverCommand a command for semver version management based on git tags.
// Returns the CLI command.
func NewSemverCommand(app *cli.App) *cli.Command {
	var command = "semver"
	var description = "semver version operations based on git tags"
	var aliases = []string{"s"}

	var subcommands = []*cli.Command{
		NewSemverGetCommand(app),
		NewSemverReleaseCommand(app),
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

// NewSemverGetCommand a command to get the latest semver version based on git tags.
// Returns the CLI command.
func NewSemverGetCommand(app *cli.App) *cli.Command {
	var command = "get"
	var description = "gets the semver version based on the latest git tag"
	var aliases = []string{"g"}

	var action = func(c *cli.Context) (err error) {
		var commander = commands.NewExecCommander()
		var tagger = git.NewTagger(commander)

		fmt.Println(tagger.GetTag())

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

// NewSemverReleaseCommand a command to increment the semver version based on git tags and strategies.
// It creates a git tag for the new version.
// The [strategy|s] flag allows you to choose which semver level to increment.
// Returns the CLI command.
func NewSemverReleaseCommand(app *cli.App) *cli.Command {
	var command = "release"
	var description = "increments the semver version based on the latest git tag and a strategy"
	var aliases = []string{"r"}

	var flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "strategy",
			Aliases: []string{"s"},
			Usage:   "determines the semver level increase",
			Value:   "auto",
		},
	}

	var action = func(c *cli.Context) (err error) {
		var strategyName = c.String("strategy")

		var commander = commands.NewExecCommander()
		var tagger = git.NewTagger(commander)
		var strategy = semver.NewStrategy(strategyName)
		var level string

		if level, err = strategy.GetLevel(); err != nil {
			return
		}

		var tag = tagger.GetTag()
		var version string

		if version, err = semver.IncrementByLevel(tag, level); err != nil {
			return
		}

		err = tagger.CreateTag(version)

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
