package semver

import blangsemver "github.com/blang/semver/v4"

// Major semver version level for major
const Major = "major"

// MajorStrategy implementation of the Strategy interface.
// It makes use of the major level of semver versions.
type MajorStrategy struct{}

// NewMajorStrategy creates a new MajorStrategy.
// Returns the new MajorStrategy.
func NewMajorStrategy() MajorStrategy {
	return MajorStrategy{}
}

// Increment increments a given version using the MajorStrategy.
// Returns the incremented version.
func (majorStrategy MajorStrategy) Increment(targetVersion string) (nextVersion string, err error) {
	var version blangsemver.Version

	if version, err = blangsemver.Parse(targetVersion); err != nil {
		return
	}

	// at point of writing IncrementMajor always returns a nil value error
	_ = version.IncrementMajor()

	return version.FinalizeVersion(), err
}
