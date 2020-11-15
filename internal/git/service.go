package git

// Service interface to run git operations.
type Service interface {
	CreateTag(tag string) (err error)
	GetLatestCommitMessage() (message string, err error)
	GetTag() (output string)
}
