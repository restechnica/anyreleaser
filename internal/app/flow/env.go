package flow

import (
	"fmt"
	"os"

	"github.com/restechnica/anyreleaser/internal/app"

	"github.com/joho/godotenv"
)

type LoadEnvFiles struct{}

func (pipe LoadEnvFiles) Run(ctx *app.Context) (err error) {
	if len(ctx.Config.Env.Files) > 0 {
		if err = godotenv.Overload(ctx.Config.Env.Files...); err != nil {
			return fmt.Errorf("could not set env file variables: %s", err)
		}
	}

	return
}

type LoadEnvScripts struct{}

func (pipe LoadEnvScripts) Run(ctx *app.Context) (err error) {
	var cmder = ctx.Commander

	for _, script := range ctx.Config.Env.Scripts {
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

func (pipe LoadEnvVars) Run(ctx *app.Context) (err error) {
	for key, value := range ctx.Config.Env.Vars {
		if err = os.Setenv(key, value); err != nil {
			return fmt.Errorf("could not set env var '%s=%s'", key, value)
		}
	}

	return
}
