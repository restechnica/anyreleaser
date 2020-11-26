package flow

import "github.com/urfave/cli/v2"

type Pipe interface {
	Run(ctx *cli.Context) error
}
