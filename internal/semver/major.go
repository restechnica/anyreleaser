package semver

// Major semver version level for major
const Major = "major"

type MajorStrategy struct{}

func NewMajorStrategy() MajorStrategy {
	return MajorStrategy{}
}

// GetLevel gets the level to increment using the MajorStrategy.
// Returns the Major level to increment.
func (s MajorStrategy) GetLevel() (level string, err error) {
	return Major, err
}
