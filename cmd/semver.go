package cmd

import "github.com/urfave/cli/v2"

// NewSemverCommand a command for semver version management based on git tags.
// Returns the CLI command.
func NewSemverCommand(app *cli.App) *cli.Command {
	var command = "semver"
	var description = "semver version operations based on git tags"
	var aliases = []string{"s"}

	var commands = []*cli.Command{}

	return &cli.Command{
		Aliases:         aliases,
		HideHelp:        app.HideHelp,
		HideHelpCommand: app.HideHelpCommand,
		Name:            command,
		Subcommands:     commands,
		Usage:           description,
	}
}
