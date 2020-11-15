package semver

import "github.com/restechnica/backbone-cli/internal/git"

type Manager struct {
	GitService git.Service
}

func NewManager(gitService git.Service) Manager {
	return Manager{GitService: gitService}
}

func (manager Manager) GetStrategy(name string) Strategy {
	switch name {
	case Patch:
		return PatchStrategy{}
	case Minor:
		return MinorStrategy{}
	case Major:
		return MajorStrategy{}
	case "auto":
		var gitCommitStrategy = NewGitCommitStrategy(manager.GitService)
		return NewAutoStrategy(gitCommitStrategy)
	case "commit":
		return NewGitCommitStrategy(manager.GitService)
	default:
		return PatchStrategy{}
	}
}
