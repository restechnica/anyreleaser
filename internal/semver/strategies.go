package semver

import (
	"fmt"
	"regexp"

	"github.com/restechnica/backbone-cli/internal/commands"
	"github.com/restechnica/backbone-cli/internal/git"
)

var DefaultFeatureRegex = regexp.MustCompile(`(\[feature\]|feature\/)`)
var DefaultFixRegex = regexp.MustCompile(`(\[fix\]|fix\/)`)
var DefaultReleaseRegex = regexp.MustCompile(`(\[release\]|release\/)`)

type Strategy interface {
	GetLevel() (level string, err error)
}

func NewStrategy(strategy string) Strategy {
	switch strategy {
	case Patch:
		return PatchStrategy{}
	case Minor:
		return MinorStrategy{}
	case Major:
		return MajorStrategy{}
	case "auto":
		var gitCommitStrategy = NewGitCommitStrategy(git.Committer{})
		return NewAutoStrategy(gitCommitStrategy)
	case "commit":
		return NewGitCommitStrategy(git.Committer{})
	default:
		return PatchStrategy{}
	}
}

type AutoStrategy struct {
	GitCommitStrategy
	PatchStrategy
}

func NewAutoStrategy(gitCommitStrategy GitCommitStrategy) AutoStrategy {
	return AutoStrategy{GitCommitStrategy: gitCommitStrategy, PatchStrategy: PatchStrategy{}}
}

// GetLevel gets the level to increment using the AutoStrategy.
// It will attempt to determine the level with several strategies:
//		1. the GitCommitStrategy
// 		2. the PatchStrategy
// Returns the determined level or an error if anything went wrong.
func (s AutoStrategy) GetLevel() (level string, err error) {
	if level, err = s.GitCommitStrategy.GetLevel(); err != nil {
		return s.PatchStrategy.GetLevel()
	}

	return
}

type GitCommitStrategy struct {
	Committer git.Committer
}

func NewGitCommitStrategy(committer git.Committer) GitCommitStrategy {
	return GitCommitStrategy{Committer: committer}
}

// GetLevel gets the level to increment using the GitCommitStrategy.
// It tries to determine which level to increment based on the latest git commit message.
// Returns the level to increment.
func (s GitCommitStrategy) GetLevel() (level string, err error) {
	var commander = commands.NewExecCommander()
	var committer = git.NewCommmitter(commander)

	var message string

	if message, err = committer.GetLatestCommitMessage(); err != nil {
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

type MajorStrategy struct{}

// GetLevel gets the level to increment using the MajorStrategy.
// Returns the Major level to increment.
func (s MajorStrategy) GetLevel() (level string, err error) {
	return Major, err
}

type MinorStrategy struct{}

// GetLevel gets the level to increment using the MinorStrategy.
// Returns the Minor level to increment.
func (s MinorStrategy) GetLevel() (level string, err error) {
	return Minor, err
}

type PatchStrategy struct{}

// GetLevel gets the level to increment using the PatchStrategy.
// Returns the Patch level to increment.
func (s PatchStrategy) GetLevel() (level string, err error) {
	return Patch, err
}
