package semver

import (
	blangsemver "github.com/blang/semver/v4"
)

// Patch semver version level for patch
const Patch = "patch"

// PatchStrategy implementation of the Strategy interface.
// It makes use of the patch level of semver versions.
type PatchStrategy struct{}

// NewPatchStrategy creates a new PatchStrategy.
// Returns the new PatchStrategy.
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

	// at point of writing IncrementPatch always returns a nil value error
	_ = version.IncrementPatch()

	return version.FinalizeVersion(), err
}
