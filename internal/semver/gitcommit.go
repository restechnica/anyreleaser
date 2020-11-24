package semver

import (
	"fmt"
	"strings"

	"github.com/restechnica/anyreleaser/internal/git"
)

// GitCommit strategy name for GitCommitStrategy.
const GitCommit = "git-commit"

// GitCommitStrategy implementation of the Strategy interface.
// It makes use of several matching strategies based on git commit messages.
type GitCommitStrategy struct {
	matches    map[string]string
	gitService git.Service
}

// NewGitCommitStrategy creates a new GitCommitStrategy.
// Returns the new GitCommitStrategy.
func NewGitCommitStrategy(matches map[string]string, gitService git.Service) GitCommitStrategy {
	return GitCommitStrategy{matches: matches, gitService: gitService}
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

// GetMatchedStrategy gets the strategy that testMatches specific tokens within the git commit message.
// It returns the matched strategy.
func (strategy GitCommitStrategy) GetMatchedStrategy(message string) (matched Strategy, err error) {
	for match, strategy := range strategy.matches {
		if strings.Contains(message, match) {
			switch strategy {
			case Patch:
				matched = NewPatchStrategy()
			case Minor:
				matched = NewMinorStrategy()
			case Major:
				matched = NewMajorStrategy()
			}
			return
		}
	}

	return matched, fmt.Errorf(`could not match a strategy to the commit message "%s"`, message)
}
