package flow

import (
	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/restechnica/anyreleaser/internal/app/config"
)

type LoadConfig struct {
	ConfigPath string
	Loader     config.Loader
}

func NewLoadConfig(path string) LoadConfig {
	return LoadConfig{
		ConfigPath: path,
		Loader:     config.NewYAMLLoader(),
	}
}

func (pipe LoadConfig) Run(ctx *app.Context) (err error) {
	var cfg = ctx.Config
	var path = pipe.ConfigPath

	if cfg, err = pipe.Loader.Overload(path, cfg); err == nil {
		ctx.Config = cfg
	}

	return
}

type LoadDefaultConfig struct{}

func (pipe LoadDefaultConfig) Run(ctx *app.Context) (err error) {
	var cfg = config.NewRoot()
	ctx.Config = cfg
	return
}
