package semver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

type gitCommitStrategyMock struct {
	mock.Mock
}

func NewGitCommitStrategyMock() *gitCommitStrategyMock {
	return &gitCommitStrategyMock{}
}

func (mock *gitCommitStrategyMock) Increment(targetVersion string) (nextVersion string, err error) {
	args := mock.Called(targetVersion)
	return args.String(0), args.Error(1)
}

func TestAutoStrategy_AutoConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "auto"
		var got = Auto

		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})
}

func TestAutoStrategy_Increment(t *testing.T) {
	t.Run("IncrementWithGitCommitStrategy", func(t *testing.T) {
		const target = "0.0.0"
		const want = "1.0.0"

		var gitCommitStrategy = NewGitCommitStrategyMock()
		gitCommitStrategy.On("Increment", target).Return(want, nil)

		var autoStrategy = NewAutoStrategy(gitCommitStrategy)
		var got, err = autoStrategy.Increment(target)

		assert.NoError(t, err)
		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})

	t.Run("IncrementWithPatchStrategy", func(t *testing.T) {
		const target = "0.0.0"
		const want = "0.0.1"

		var gitCommitStrategy = NewGitCommitStrategyMock()
		gitCommitStrategy.On("Increment", target).Return("", fmt.Errorf("some-error"))

		var autoStrategy = NewAutoStrategy(gitCommitStrategy)
		var got, err = autoStrategy.Increment(target)

		assert.NoError(t, err)
		assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
	})

	type ErrorTest struct {
		Name    string
		Version string
	}

	var errorTests = []ErrorTest{
		{Name: "ReturnErrorOnInvalidVersion", Version: "invalid"},
		{Name: "ReturnErrorOnInvalidCharacter", Version: "v1.0.0"},
	}

	for _, test := range errorTests {
		t.Run(test.Name, func(t *testing.T) {
			var gitCommitStrategy = NewGitCommitStrategyMock()
			gitCommitStrategy.On("Increment", test.Version).Return("", fmt.Errorf("some-error"))

			var autoStrategy = NewAutoStrategy(gitCommitStrategy)
			var _, err = autoStrategy.Increment(test.Version)

			assert.Error(t, err)
		})
	}
}
