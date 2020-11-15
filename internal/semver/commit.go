package semver

import (
	"fmt"
	"regexp"

	blangsemver "github.com/blang/semver/v4"

	"github.com/restechnica/backbone-cli/internal/commands"
	"github.com/restechnica/backbone-cli/internal/git"
)

var DefaultFeatureRegex = regexp.MustCompile(`(\[feature]|feature/)`)
var DefaultFixRegex = regexp.MustCompile(`(\[fix]|fix/)`)
var DefaultReleaseRegex = regexp.MustCompile(`(\[release]|release/)`)

type GitCommitStrategy struct {
	gitService git.Service
}

func NewGitCommitStrategy(gitService git.Service) GitCommitStrategy {
	return GitCommitStrategy{gitService: gitService}
}

// GetLevel gets the level to increment using the GitCommitStrategy.
// It tries to determine which level to increment based on the latest git commit message.
// Returns the level to increment.
func (s GitCommitStrategy) GetLevel() (level string, err error) {
	var commander = commands.NewExecCommander()
	var gitService = git.NewCLIService(commander)

	var message string

	if message, err = gitService.GetLatestCommitMessage(); err != nil {
		return
	}

	var strategy Strategy

	if DefaultFixRegex.MatchString(message) {
		strategy = PatchStrategy{}
	} else if DefaultFeatureRegex.MatchString(message) {
		strategy = MinorStrategy{}
	} else if DefaultReleaseRegex.MatchString(message) {
		strategy = MajorStrategy{}
	} else {
		return level, fmt.Errorf("unable to determine semver level to increment")
	}

	return strategy.GetLevel()
}

// Increment increments a given version using the PatchStrategy.
// Returns the incremented version.
func (s GitCommitStrategy) Increment(targetVersion string) (nextVersion string, err error) {
	var version blangsemver.Version

	if version, err = blangsemver.Parse(targetVersion); err != nil {
		return
	}

	if err = version.IncrementPatch(); err != nil {
		return
	}

	return version.FinalizeVersion(), err
}
