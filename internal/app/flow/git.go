package flow

import (
	"github.com/restechnica/anyreleaser/internal/app"
)

type SetGit struct{}

func (pipe SetGit) Run(ctx *app.Context) (err error) {
	var cmder = ctx.Commander

	for key, value := range ctx.Config.Git.Config {
		_ = cmder.Run("git", "config", key, value)
	}

	if ctx.Config.Git.Unshallow {
		_ = cmder.Run("git", "fetch", "--unshallow")
	}

	return
}
