package flow

import (
	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/restechnica/anyreleaser/internal/app/config"
)

type LoadConfig struct {
	ConfigPath string
}

func (pipe LoadConfig) Run(ctx *app.Context) (err error) {
	var cfg = ctx.Config
	var path = pipe.ConfigPath

	var loader = config.NewYAMLLoader()

	if cfg, err = loader.Overload(path, cfg); err == nil {
		ctx.Config = cfg
	}

	return
}

type LoadDefaultConfig struct {
	ConfigPath string
}

func (pipe LoadDefaultConfig) Run(ctx *app.Context) (err error) {
	var cfg = config.NewRoot()
	ctx.Config = cfg
	return
}
