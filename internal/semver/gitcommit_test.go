package semver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

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

		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})
}

func TestGitCommitStrategy_GetMatchedStrategy(t *testing.T) {
	type Test struct {
		Message string
		Name    string
		Want    Strategy
	}

	var tests = []Test{
		{Name: "GetPatchStrategyWithBrackets", Message: "[fix] some message", Want: NewPatchStrategy()},
		{Name: "GetPatchStrategyWithTrailingSlash", Message: "Merged: repo/fix/some-error", Want: NewPatchStrategy()},
		{Name: "GetMinorStrategyWithBrackets", Message: "some [feature] message", Want: NewMinorStrategy()},
		{Name: "GetMinorStrategyWithTrailingSlash", Message: "Merged: repo/feature/some-error", Want: NewMinorStrategy()},
		{Name: "GetMajorStrategyWithBrackets", Message: "some message [release]", Want: NewMajorStrategy()},
		{Name: "GetMajorStrategyWithTrailingSlash", Message: "Merged: repo/release/some-error", Want: NewMajorStrategy()},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var want = test.Want

			var service = NewGitCommitStrategyGitServiceMock()
			var gitCommitStrategy = NewGitCommitStrategy(service)
			var got, err = gitCommitStrategy.GetMatchedStrategy(test.Message)

			assert.NoError(t, err)
			assert.IsType(t, want, got, `want: %s, got: %s`, want, got)
		})
	}

	type ErrorTest struct {
		Message string
		Name    string
	}

	var errorTests = []ErrorTest{
		{Name: "ReturnErrorOnUnmatchedStrategy", Message: "[fix some message"},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var want = fmt.Sprintf("could not match a strategy to the commit message %s", test.Message)

			var service = NewGitCommitStrategyGitServiceMock()
			var gitCommitStrategy = NewGitCommitStrategy(service)
			var _, err = gitCommitStrategy.GetMatchedStrategy(test.Message)

			assert.Error(t, err)
			assert.Equal(t, err.Error(), want, `want: "%s", got: "%s"`, want, err.Error())
		})
	}
}
