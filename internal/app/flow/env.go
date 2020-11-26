package flow

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/restechnica/anyreleaser/internal/app/config"
	"github.com/restechnica/anyreleaser/internal/commands"
	"github.com/urfave/cli/v2"
)

type LoadEnvFiles struct{}

func (pipe LoadEnvFiles) Run(ctx *cli.Context) (err error) {
	var cfg = ctx.App.Metadata["config"].(config.Root)

	if len(cfg.Env.Files) > 0 {
		if err = godotenv.Overload(cfg.Env.Files...); err != nil {
			return fmt.Errorf("could not set env file variables: %s", err)
		}
	}

	return
}

type LoadEnvScripts struct{}

func (pipe LoadEnvScripts) Run(ctx *cli.Context) (err error) {
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
				return fmt.Errorf("could not set env var '%s=%s' from '%s %s'",
					key, value, script.Bin, script.Path)
			}
		}
	}

	return
}

type LoadEnvVars struct{}

func (pipe LoadEnvVars) Run(ctx *cli.Context) (err error) {
	var cfg = ctx.App.Metadata["config"].(config.Root)

	for key, value := range cfg.Env.Vars {
		if err = os.Setenv(key, value); err != nil {
			return fmt.Errorf("could not set env var '%s=%s'", key, value)
		}
	}

	return
}
