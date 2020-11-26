package semver

import (
	"github.com/restechnica/anyreleaser/internal/app/config"
	"github.com/restechnica/anyreleaser/internal/git"
)

type Manager struct {
	Config     config.Root
	GitService git.Service
}

func NewManager(config config.Root, gitService git.Service) Manager {
	return Manager{Config: config, GitService: gitService}
}

func (manager Manager) GetStrategy(name string) Strategy {
	switch name {
	case Patch:
		return NewPatchStrategy()
	case Minor:
		return NewMinorStrategy()
	case Major:
		return NewMajorStrategy()
	case Auto:
		var gitCommitStrategy = NewGitCommitStrategy(manager.Config.Semver.Matches, manager.GitService)
		return NewAutoStrategy(gitCommitStrategy)
	case GitCommit:
		return NewGitCommitStrategy(manager.Config.Semver.Matches, manager.GitService)
	default:
		return NewPatchStrategy()
	}
}
