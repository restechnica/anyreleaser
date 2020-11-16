package semver

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

type gitCommitStrategyGitServiceMock struct {
	mock.Mock
}

func NewGitCommitStrategyGitServiceMock() *gitCommitStrategyGitServiceMock {
	return &gitCommitStrategyGitServiceMock{}
}

func (mock *gitCommitStrategyGitServiceMock) CreateTag(tag string) (err error) {
	args := mock.Called(tag)
	return args.Error(0)
}

func (mock *gitCommitStrategyGitServiceMock) GetLatestCommitMessage() (message string, err error) {
	args := mock.Called()
	return args.String(0), args.Error(1)
}

func (mock *gitCommitStrategyGitServiceMock) GetTag() (output string) {
	args := mock.Called()
	return args.String(0)
}

func TestGitCommitStrategy_GitCommitConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "git-commit"
		var got = GitCommit

		assert.Equal(t, want, got, fmt.Sprintf(`want: "%s", got: "%s"`, want, got))
	})
}

func TestGitCommitStrategy_GetMatchedStrategy(t *testing.T) {
	type GetMatchedStrategyTest struct {
		Message string
		Name    string
		Want    Strategy
	}

	var getMatchedStrategyTests = []GetMatchedStrategyTest{
		{Name: "GetPatchStrategyWithBrackets", Message: "[fix some message", Want: NewPatchStrategy()},
	}

	for _, test := range getMatchedStrategyTests {
		t.Run(test.Name, func(t *testing.T) {
			var want = test.Want

			var service = NewGitCommitStrategyGitServiceMock()
			var gitCommitStrategy = NewGitCommitStrategy(service)
			var got, err = gitCommitStrategy.GetMatchedStrategy(test.Message)

			assert.NoError(t, err)
			assert.IsType(t, want, got, `want: %s, got: %s`, want, got)
		})
	}
}
