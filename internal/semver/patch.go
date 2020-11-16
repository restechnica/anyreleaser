package semver

import (
	blangsemver "github.com/blang/semver/v4"
)

// Patch semver version level for patch
const Patch = "patch"

type PatchStrategy struct{}

func NewPatchStrategy() PatchStrategy {
	return PatchStrategy{}
}

// Increment increments a given version using the PatchStrategy.
// Returns the incremented version.
func (patchStrategy PatchStrategy) Increment(targetVersion string) (nextVersion string, err error) {
	var version blangsemver.Version

	if version, err = blangsemver.Parse(targetVersion); err != nil {
		return
	}

	if err = version.IncrementPatch(); err != nil {
		return
	}

	return version.FinalizeVersion(), err
}
