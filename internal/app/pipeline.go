package app

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/restechnica/anyreleaser/internal/config"
	"github.com/urfave/cli/v2"
)

type Pipe interface {
	Run(ctx *cli.Context) error
}

type Pipeline struct {
	Pipes []Pipe
}

func (pipeline *Pipeline) Add(pipe Pipe) {
	pipeline.Pipes = append(pipeline.Pipes, pipe)
}

func (pipeline Pipeline) Run(ctx *cli.Context) (err error) {
	for _, pipe := range pipeline.Pipes {
		if err = pipe.Run(ctx); err != nil {
			return
		}
	}
	return
}

type DefaultConfigPipe struct{}

func (pipe DefaultConfigPipe) Run(ctx *cli.Context) (err error) {
	ctx.App.Metadata["config"] = config.NewRoot()
	return
}

type ConfigPipe struct{}

func (pipe ConfigPipe) Run(ctx *cli.Context) (err error) {
	var cfg = ctx.App.Metadata["config"].(config.Root)
	var path = ctx.String("config")

	if cfg, err = config.Overload(path, cfg); err == nil {
		ctx.App.Metadata["config"] = cfg
	}

	return
}

type EnvFilesPipe struct{}

func (pipe EnvFilesPipe) Run(ctx *cli.Context) (err error) {
	var cfg = ctx.App.Metadata["config"].(config.Root)

	for _, path := range cfg.Env.Files {
		if err = godotenv.Load(path); err != nil {
			return fmt.Errorf("could not set env file variables for '%s'", path)
		}
	}

	return
}

type EnvVarsPipe struct{}

func (pipe EnvVarsPipe) Run(ctx *cli.Context) (err error) {
	var cfg = ctx.App.Metadata["config"].(config.Root)

	for key, value := range cfg.Env.Vars {
		if err = os.Setenv(key, value); err != nil {
			return fmt.Errorf("could not set env var '%s=%s'", key, value)
		}
	}

	return
}

type GitUnshallowPipe struct{}

func (pipe GitUnshallowPipe) Run(ctx *cli.Context) (err error) {
	var cfg = ctx.App.Metadata["config"].(config.Root)

	if cfg.Versioning.Git.Unshallow {
		fmt.Println("on")
	}

	return
}
