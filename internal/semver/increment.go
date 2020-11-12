package semver

import (
	"fmt"

	blangsemver "github.com/blang/semver/v4"
)

// Increment increments the specified level of a semver version.
// It supports major, minor, patch and discards any prerelease/build information.
// It returns the incremented version or any errors if it failed.
func Increment(currentVersion string, level string) (nextVersion string, err error) {
	var version blangsemver.Version

	if version, err = blangsemver.Parse(currentVersion); err != nil {
		return
	}

	switch level {
	case PATCH:
		err = version.IncrementPatch()
	case MINOR:
		err = version.IncrementMinor()
	case MAJOR:
		err = version.IncrementMajor()
	default:
		err = fmt.Errorf("\"%s\" is not a valid semver version level", level)
	}

	nextVersion = version.FinalizeVersion()

	return
}
