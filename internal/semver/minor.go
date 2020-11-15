package semver

// Minor semver version level for minor
const Minor = "minor"

type MinorStrategy struct{}

func NewMinorStrategy() MinorStrategy {
	return MinorStrategy{}
}

// GetLevel gets the level to increment using the MinorStrategy.
// Returns the Minor level to increment.
func (s MinorStrategy) GetLevel() (level string, err error) {
	return Minor, err
}
