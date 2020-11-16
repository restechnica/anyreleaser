package semver

// Strategy interface to increment a specific semver level.
type Strategy interface {
	Increment(targetVersion string) (nextVersion string, err error)
}
