package semver

// Strategy interface to determine semver levels and increment them.
type Strategy interface {
	GetLevel() (level string, err error)
	Increment(targetVersion string) (nextVersion string, err error)
}
