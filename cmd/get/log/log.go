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

	baseTagFlag = &cli.StringFlag{
		Name:    "base-tag",
		Aliases: []string{"b"},
		Usage:   "the oldest version to start the log from",
	}

	topTagFlag = &cli.StringFlag{
		Name:    "top-tag",
		Aliases: []string{"t"},
		Usage:   "the newest version to end the log at",
	}

	flags = []cli.Flag{
		baseTagFlag,
		topTagFlag,
	}
)

// NewCommand a command to get a changelog.
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

func action(clictx *cli.Context) (err error) {
	var appctx, ok = clictx.App.Metadata[flow.AppContext].(*app.Context)

	if !ok {
		return fmt.Errorf("something went wrong with fetching the app context from the cli context")
	}

	var cmder = appctx.Commander

	fmt.Println(cmder)

	var baseTagIsSet = clictx.IsSet("base-tag")
	var topTagIsSet = clictx.IsSet("top-tag")

	var baseTag, topTag string

	if !baseTagIsSet && !topTagIsSet {
		// do tag_before_latest -> latest
		fmt.Println("hoi", baseTag, topTag)
	} else if baseTagIsSet {
		// do base -> latest
		err = fmt.Errorf("tag ranges are not supported yet")
	} else if topTagIsSet {
		// do first commit -> top
		err = fmt.Errorf("tag ranges are not supported yet")
	} else {
		// do base -> top
		err = fmt.Errorf("tag ranges are not supported yet")
	}

	return
}
