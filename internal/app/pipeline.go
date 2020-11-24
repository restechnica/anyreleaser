package app

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/restechnica/anyreleaser/internal/commands"
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

type CommanderPipe struct{}

func (pipe CommanderPipe) Run(ctx *cli.Context) (err error) {
	ctx.App.Metadata["commander"] = commands.NewExecCommander()
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

	if err = godotenv.Overload(cfg.Env.Files...); err != nil {
		return fmt.Errorf("could not set env file variables: %s", err)
	}

	return
}

type EnvScriptsPipe struct{}

func (pipe EnvScriptsPipe) Run(ctx *cli.Context) (err error) {
	var cfg = ctx.App.Metadata["config"].(config.Root)
	var cmder = ctx.App.Metadata["commander"].(commands.Commander)

	for _, script := range cfg.Env.Scripts {
		var output string

		if output, err = cmder.Output(script.Bin, script.Path); err != nil {
			return fmt.Errorf("failed to run '%s %s'", script.Bin, script.Path)
		}

		var vars map[string]string

		if vars, err = godotenv.Unmarshal(output); err != nil {
			return fmt.Errorf("failed to parse output from '%s %s'", script.Bin, script.Path)
		}

		for key, value := range vars {
			if err = os.Setenv(key, value); err != nil {
				return fmt.Errorf("could not set env var '%s=%s' from '%s=%s'",
					key, value, script.Bin, script.Path)
			}
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
