package semver

import blangsemver "github.com/blang/semver/v4"

// Minor semver version level for minor
const Minor = "minor"

type MinorStrategy struct{}

func NewMinorStrategy() MinorStrategy {
	return MinorStrategy{}
}

// GetLevel gets the level to increment using the MinorStrategy.
// Returns the Minor level to increment.
func (MinorStrategy) GetLevel() (level string, err error) {
	return Minor, err
}

// Increment increments a given version using the MinorStrategy.
// Returns the incremented version.
func (minorStrategy MinorStrategy) Increment(targetVersion string) (nextVersion string, err error) {
	var version blangsemver.Version

	if version, err = blangsemver.Parse(targetVersion); err != nil {
		return
	}

	if err = version.IncrementMinor(); err != nil {
		return
	}

	return version.FinalizeVersion(), err
}
