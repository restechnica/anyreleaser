package git

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCommander struct {
	mock.Mock
}

func (m mockCommander) Output(name string, arg ...string) (string, error) {
	args := m.Called(name, arg)
	return args.String(0), args.Error(1)
}

func (m mockCommander) Run(name string, arg ...string) error {
	args := m.Called(name, arg)
	return args.Error(0)
}

func TestTagger_CreateTag(t *testing.T) {
	const methodName = "Run"
	const testTag = "1.2.3"

	t.Run("HappyPath", func(t *testing.T) {
		var want error

		var commander = new(mockCommander)
		commander.On(methodName, mock.Anything, mock.Anything).Return(want)

		var tagger = NewTagger(commander)
		var got = tagger.CreateTag(testTag)

		assert.NoError(t, got)
	})

	t.Run("GitError", func(t *testing.T) {
		var want = errors.New("some-error")

		var commander = new(mockCommander)
		commander.On("Run", mock.Anything, mock.Anything).Return(want)

		var tagger = NewTagger(commander)
		var got = tagger.CreateTag(testTag)

		assert.Error(t, got)
	})
}

func TestTagger_GetTag(t *testing.T) {
	type GetTagTest struct {
		CommanderErr error
		Name         string
		Expected     string
	}

	var getTagTests = []GetTagTest{
		{Name: "HappyPath", CommanderErr: nil, Expected: "1.0.0"},
		{Name: "GitError", CommanderErr: errors.New("some-error"), Expected: DefaultTag},
	}

	for _, test := range getTagTests {
		const methodName = "Output"

		t.Run(test.Name, func(t *testing.T) {
			var commander = new(mockCommander)
			commander.On(methodName, mock.Anything, mock.Anything).Return(test.Expected, test.CommanderErr)

			var tagger = NewTagger(commander)

			var got = tagger.GetTag()
			var want = test.Expected

			assert.Equal(t, want, got, fmt.Sprintf("We wanted %s and got %s", want, got))
		})
	}
}
