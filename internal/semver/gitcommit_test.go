package semver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitCommitStrategy_GitCommitConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "git-commit"
		var got = GitCommit

		assert.Equal(t, want, got, fmt.Sprintf(`want: "%s", got: "%s"`, want, got))
	})
}
