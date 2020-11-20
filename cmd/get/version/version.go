package version

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/restechnica/anyreleaser/internal/commands"
	"github.com/restechnica/anyreleaser/internal/git"
)

// NewCommand a command to get the current semver version.
//// Returns the CLI command.
func NewCommand(app *cli.App) *cli.Command {
	var command = "version"
	var description = "gets the current semver version"
	var aliases = []string{"v"}

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
