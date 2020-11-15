package semver

import (
	"fmt"

	blangsemver "github.com/blang/semver/v4"
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

// IncrementByLevel increments the specified level of a semver version.
// It supports major, minor, patch and discards any prerelease/build information.
// It returns the incremented version or any errors if it failed.
func (Manager) IncrementByLevel(currentVersion string, level string) (nextVersion string, err error) {
	var version blangsemver.Version

	if version, err = blangsemver.Parse(currentVersion); err != nil {
		return
	}

	switch level {
	case Patch:
		err = version.IncrementPatch()
	case Minor:
		err = version.IncrementMinor()
	case Major:
		err = version.IncrementMajor()
	default:
		err = fmt.Errorf("\"%s\" is not a valid semver version level", level)
	}

	nextVersion = version.FinalizeVersion()

	return
}

// IncrementByStrategy increments the specified level of a semver version.
// It uses a strategy to determine the level to increment.
// It returns the incremented version or any errors if it failed.
func (manager Manager) IncrementByStrategy(currentVersion string, strategy Strategy) (nextVersion string, err error) {
	var level string

	if level, err = strategy.GetLevel(); err != nil {
		return
	}

	return manager.IncrementByLevel(currentVersion, level)
}
