package init

import (
	"os"

	"github.com/urfave/cli/v2"
)

const (
	configFolder   = ".any"
	permissionBits = 0755
)

// NewCommand a command to get the current semver version.
// Returns the CLI command.
func NewCommand(app *cli.App) *cli.Command {
	var command = "init"
	var description = "initialize anyreleaser"
	var aliases = []string{"i"}

	var action = func(c *cli.Context) (err error) {
		if err = os.MkdirAll(configFolder, permissionBits); err != nil {
			return
		}
		// ask some questions
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
