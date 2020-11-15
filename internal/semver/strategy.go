package semver

type Strategy interface {
	GetLevel() (level string, err error)
}
