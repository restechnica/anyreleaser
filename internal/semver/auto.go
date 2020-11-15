package semver

import blangsemver "github.com/blang/semver/v4"

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
func (autoStrategy AutoStrategy) GetLevel() (level string, err error) {
	if level, err = autoStrategy.GitCommitStrategy.GetLevel(); err != nil {
		return autoStrategy.PatchStrategy.GetLevel()
	}
	return
}

// Increment increments a given version using the PatchStrategy.
// Returns the incremented version.
func (autoStrategy AutoStrategy) Increment(targetVersion string) (nextVersion string, err error) {
	var version blangsemver.Version

	if version, err = blangsemver.Parse(targetVersion); err != nil {
		return
	}

	if err = version.IncrementPatch(); err != nil {
		return
	}

	return version.FinalizeVersion(), err
}
