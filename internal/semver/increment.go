package semver

import (
	"fmt"

	blangsemver "github.com/blang/semver/v4"
)

// IncrementByLevel increments the specified level of a semver version.
// It supports major, minor, patch and discards any prerelease/build information.
// It returns the incremented version or any errors if it failed.
func IncrementByLevel(currentVersion string, level string) (nextVersion string, err error) {
	var version blangsemver.Version

	if version, err = blangsemver.Parse(currentVersion); err != nil {
		return
	}

	switch level {
	case Patch:
		err = version.IncrementPatch()
	case Minor:
		err = version.IncrementMinor()
	case Major:
		err = version.IncrementMajor()
	default:
		err = fmt.Errorf("\"%s\" is not a valid semver version level", level)
	}

	nextVersion = version.FinalizeVersion()

	return
}

// IncrementByStrategy increments the specified level of a semver version.
// It uses a strategy to determine the level to increment.
// It returns the incremented version or any errors if it failed.
func IncrementByStrategy(currentVersion string, strategy Strategy) (nextVersion string, err error) {
	var level string

	if level, err = strategy.GetLevel(); err != nil {
		return
	}

	return IncrementByLevel(currentVersion, level)
}
