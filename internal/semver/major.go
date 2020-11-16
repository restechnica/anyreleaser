package semver

import blangsemver "github.com/blang/semver/v4"

// Major semver version level for major
const Major = "major"

type MajorStrategy struct{}

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
