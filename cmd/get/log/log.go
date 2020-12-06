package log

import (
	"fmt"
	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/restechnica/anyreleaser/internal/app/flow"
	"github.com/urfave/cli/v2"
)

const (
	command     = "log"
	description = "gets a changelog"
)

var (
	aliases = []string{"l"}

	rangeFlag = &cli.StringFlag{
		Name:    "range",
		Aliases: []string{"r"},
		Usage:   "determines the range of commits to add to the changelog",
	}

	flags = []cli.Flag{
		rangeFlag,
	}
)

// NewCommand a command to get a changelog.
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
