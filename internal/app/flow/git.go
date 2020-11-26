package flow

import (
	"github.com/restechnica/anyreleaser/internal/app/config"
	"github.com/restechnica/anyreleaser/internal/commands"
	"github.com/urfave/cli/v2"
)

type SetGit struct{}

func (pipe SetGit) Run(ctx *cli.Context) (err error) {
	var cfg = ctx.App.Metadata["config"].(config.Root)
	var cmder = ctx.App.Metadata["commander"].(commands.Commander)

	for key, value := range cfg.Git.Config {
		_ = cmder.Run("git", "config", key, value)
	}

	if cfg.Git.Unshallow {
		_ = cmder.Run("git", "fetch", "--unshallow")
	}

	return
}
