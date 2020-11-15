package semver

// Patch semver version level for patch
const Patch = "patch"

type PatchStrategy struct{}

func NewPatchStrategy() PatchStrategy {
	return PatchStrategy{}
}

// GetLevel gets the level to increment using the PatchStrategy.
// Returns the Patch level to increment.
func (s PatchStrategy) GetLevel() (level string, err error) {
	return Patch, err
}
