package log

import (
	"fmt"
	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/restechnica/anyreleaser/internal/app/flow"
	"github.com/urfave/cli/v2"
)

const (
	command     = "log"
	description = "creates a changelog"
)

var (
	aliases = []string{"l"}
)

// NewCommand a command to create a changelog.
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
	var appctx, ok = clictx.App.Metadata[flow.AppContext].(*app.Context)

	if !ok {
		return fmt.Errorf("something went wrong with fetching the app context from the cli context")
	}

	var cmder = appctx.Commander

	fmt.Println(cmder)

	return
}
