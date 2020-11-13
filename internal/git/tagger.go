package git

import (
	"github.com/restechnica/backbone-cli/internal/commands"
)

const defaultTag = "0.0.0"

type Tagger struct {
	commander commands.Commander
}

func NewTagger(commander commands.Commander) Tagger {
	return Tagger{commander: commander}
}

// CreateTag creates an annotated git tag.
// Returns an error if the command failed.
func (tagger Tagger) CreateTag(tag string) (err error) {
	return tagger.commander.Run("git", "tag", "-a", tag, "-m", tag)
}

// GetTag gets the latest annotated git tag.
// It returns the latest annotated git tag or the default "0.0.0" tag if the git command fails.
func (tagger Tagger) GetTag() (output string) {
	var err error

	if output, err = tagger.commander.Output("git", "describe"); err != nil {
		return defaultTag
	}

	return
}
