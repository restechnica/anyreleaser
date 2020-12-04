package version

import (
	"fmt"

	"github.com/restechnica/anyreleaser/internal/commands"
	"github.com/restechnica/anyreleaser/internal/git"
	"github.com/urfave/cli/v2"
)

const (
	command     = "version"
	description = "gets the current semver version"
)

var (
	aliases = []string{"v"}
)

// NewCommand a command to get the current semver version.
// Returns the CLI command.
func NewCommand(app *cli.App) *cli.Command {
	return &cli.Command{
		Action:          action,
		Aliases:         aliases,
		HideHelp:        app.HideHelp,
		HideHelpCommand: app.HideHelpCommand,
		Name:            command,
		Usage:           description,
	}
}

func action(clictx *cli.Context) (err error) {
	var commander = commands.NewExecCommander()
	var _git = git.NewCLIService(commander)

	fmt.Println(_git.GetTag())

	return
}
