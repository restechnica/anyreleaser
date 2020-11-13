package git

import (
	"github.com/restechnica/backbone-cli/internal/commands"
)

type Committer struct {
	commander commands.Commander
}

func NewCommmitter(commander commands.Commander) Committer {
	return Committer{commander: commander}
}

// GetLatestCommitMessage gets the latest commit message from git.
// Returns the commit message.
func (committer Committer) GetLatestCommitMessage() (message string, err error) {
	return committer.commander.Output("git", "show", "-s", "--format='%s'")
}
