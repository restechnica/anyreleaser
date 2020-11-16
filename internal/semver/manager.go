package semver

import (
	"github.com/restechnica/backbone-cli/internal/git"
)

type Manager struct {
	GitService git.Service
}

func NewManager(gitService git.Service) Manager {
	return Manager{GitService: gitService}
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
		var gitCommitStrategy = NewGitCommitStrategy(manager.GitService)
		return NewAutoStrategy(gitCommitStrategy)
	case GitCommit:
		return NewGitCommitStrategy(manager.GitService)
	default:
		return NewPatchStrategy()
	}
}
