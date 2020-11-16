package semver

import (
	"fmt"
	"regexp"

	"github.com/restechnica/backbone-cli/internal/git"
)

var DefaultFeatureRegex = regexp.MustCompile(`(\[feature]|feature/)`)
var DefaultFixRegex = regexp.MustCompile(`(\[fix]|fix/)`)
var DefaultReleaseRegex = regexp.MustCompile(`(\[release]|release/)`)

// GitCommit strategy name for GitCommitStrategy.
const GitCommit = "git-commit"

// GitCommitStrategy implementation of the Strategy interface.
// It makes use of several matching strategies based on git commit messages.
type GitCommitStrategy struct {
	gitService git.Service
}

// NewGitCommitStrategy creates a new GitCommitStrategy.
// Returns the new GitCommitStrategy.
func NewGitCommitStrategy(gitService git.Service) GitCommitStrategy {
	return GitCommitStrategy{gitService: gitService}
}

// Increment increments a given version using the GitCommitStrategy.
// Returns the incremented version.
func (strategy GitCommitStrategy) Increment(targetVersion string) (nextVersion string, err error) {
	var message string
	var matchedStrategy Strategy

	if message, err = strategy.gitService.GetLatestCommitMessage(); err != nil {
		return
	}

	if matchedStrategy, err = strategy.GetMatchedStrategy(message); err != nil {
		return
	}

	return matchedStrategy.Increment(targetVersion)
}

// GetMatchedStrategy gets the strategy that matches specific tokens within the git commit message.
// It returns the matched strategy.
func (strategy GitCommitStrategy) GetMatchedStrategy(message string) (matched Strategy, err error) {
	if DefaultFixRegex.MatchString(message) {
		matched = NewPatchStrategy()
		return
	}

	if DefaultFeatureRegex.MatchString(message) {
		matched = NewMinorStrategy()
		return
	}

	if DefaultReleaseRegex.MatchString(message) {
		matched = NewMajorStrategy()
		return
	}

	return matched, fmt.Errorf(`could not match a strategy to the commit message "%s"`, message)
}
