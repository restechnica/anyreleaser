package semver

import (
	"fmt"
	"testing"

	"github.com/restechnica/anyreleaser/internal/mocks"
	"github.com/stretchr/testify/assert"
)

var testMatches = map[string]string{
	"[fix]":     Patch,
	"fix/":      Patch,
	"[feature]": Minor,
	"feature/":  Minor,
	"[release]": Major,
	"release/":  Major,
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

			var service = mocks.NewMockGitService()
			var gitCommitStrategy = NewGitCommitStrategy(testMatches, service)
			var got, err = gitCommitStrategy.GetMatchedStrategy(test.Message)

			assert.NoError(t, err)
			assert.IsType(t, want, got, `want: "%s", got: "%s"`, want, got)
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
			var want = fmt.Sprintf(`could not match a strategy to the commit message "%s"`, test.Message)

			var service = mocks.NewMockGitService()
			var gitCommitStrategy = NewGitCommitStrategy(testMatches, service)
			var _, err = gitCommitStrategy.GetMatchedStrategy(test.Message)

			assert.Error(t, err)
			assert.Equal(t, err.Error(), want, `want: "%s", got: "%s"`, want, err.Error())
		})
	}
}

func TestGitCommitStrategy_Increment(t *testing.T) {
	type Test struct {
		Message string
		Name    string
		Version string
		Want    string
	}

	var tests = []Test{
		{Name: "IncrementPatchWithBrackets", Message: "[fix] some message", Version: "0.0.0", Want: "0.0.1"},
		{Name: "IncrementPatchWithTrailingSlash", Message: "Merged: repo/fix/some-error", Version: "0.0.1", Want: "0.0.2"},
		{Name: "IncrementMinorWithBrackets", Message: "some [feature] message", Version: "0.0.0", Want: "0.1.0"},
		{Name: "IncrementMinorWithTrailingSlash", Message: "Merged: repo/feature/some-error", Version: "0.1.0", Want: "0.2.0"},
		{Name: "IncrementMajorWithBrackets", Message: "some message [release]", Version: "0.0.0", Want: "1.0.0"},
		{Name: "IncrementMajorWithTrailingSlash", Message: "Merged: repo/release/some-error", Version: "1.0.0", Want: "2.0.0"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var want = test.Want

			var service = mocks.NewMockGitService()
			service.On("GetLatestCommitMessage").Return(test.Message, nil)

			var gitCommitStrategy = NewGitCommitStrategy(testMatches, service)
			var got, err = gitCommitStrategy.Increment(test.Version)

			assert.NoError(t, err)
			assert.Equal(t, want, got, `want: "%s, got: "%s"`, want, got)
		})
	}

	type ErrorTest struct {
		GitError error
		Message  string
		Name     string
		Version  string
	}

	var errorTests = []ErrorTest{
		{Name: "ReturnErrorOnUnmatchedStrategy", Message: "[fix some message", Version: "0.0.0", GitError: nil},
		{Name: "ReturnErrorOnInvalidVersion", Message: "[fix] some message", Version: "invalid", GitError: nil},
		{Name: "ReturnErrorOnInvalidCharacter", Message: "[fix] some message", Version: "v1.0.0", GitError: nil},
		{Name: "ReturnErrorOnGitError", Message: "[fix] some message", Version: "1.0.0",
			GitError: fmt.Errorf("some-error")},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var service = mocks.NewMockGitService()
			service.On("GetLatestCommitMessage").Return(test.Message, test.GitError)

			var gitCommitStrategy = NewGitCommitStrategy(testMatches, service)
			var _, err = gitCommitStrategy.Increment(test.Version)

			assert.Error(t, err)
		})
	}
}
