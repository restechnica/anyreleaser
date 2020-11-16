package semver

import blangsemver "github.com/blang/semver/v4"

// Minor semver version level for minor
const Minor = "minor"

// MinorStrategy implementation of the Strategy interface.
// It makes use of the minor level of semver versions.
type MinorStrategy struct{}

// NewMinorStrategy creates a new MinorStrategy.
// Returns the new MinorStrategy.
func NewMinorStrategy() MinorStrategy {
	return MinorStrategy{}
}

// Increment increments a given version using the MinorStrategy.
// Returns the incremented version.
func (minorStrategy MinorStrategy) Increment(targetVersion string) (nextVersion string, err error) {
	var version blangsemver.Version

	if version, err = blangsemver.Parse(targetVersion); err != nil {
		return
	}

	// at point of writing IncrementMinor always returns a nil value error
	_ = version.IncrementMinor()

	return version.FinalizeVersion(), err
}
