package flow

import (
	"github.com/restechnica/anyreleaser/internal/app/config"
	"github.com/urfave/cli/v2"
)

type LoadConfig struct{}

func (pipe LoadConfig) Run(ctx *cli.Context) (err error) {
	var cfg = ctx.App.Metadata["config"].(config.Root)
	var path = ctx.String("config")

	if cfg, err = config.Overload(path, cfg); err == nil {
		ctx.App.Metadata["config"] = cfg
	}

	return
}

type LoadDefaultConfig struct{}

func (pipe LoadDefaultConfig) Run(ctx *cli.Context) (err error) {
	ctx.App.Metadata["config"] = config.NewRoot()
	return
}
