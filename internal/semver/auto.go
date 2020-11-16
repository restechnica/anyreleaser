package semver

// Auto strategy name for AutoStrategy.
const Auto = "auto"

type AutoStrategy struct {
	GitCommitStrategy
	PatchStrategy
}

func NewAutoStrategy(gitCommitStrategy GitCommitStrategy) AutoStrategy {
	return AutoStrategy{GitCommitStrategy: gitCommitStrategy, PatchStrategy: PatchStrategy{}}
}

// Increment increments a given version using the AutoStrategy.
// It will attempt to increment the target version with several strategies:
//		1. the GitCommitStrategy
// 		2. the PatchStrategy
// Returns the incremented version or an error if anything went wrong.
func (autoStrategy AutoStrategy) Increment(targetVersion string) (nextVersion string, err error) {
	if nextVersion, err = autoStrategy.GitCommitStrategy.Increment(targetVersion); err == nil {
		return
	}
	return autoStrategy.PatchStrategy.Increment(targetVersion)
}
