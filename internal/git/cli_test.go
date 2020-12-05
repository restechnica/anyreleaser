package git

import (
	"errors"
	"fmt"
	"testing"

	"github.com/restechnica/anyreleaser/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCLIService_CreateTag(t *testing.T) {
	const testTag = "1.2.3"

	t.Run("HappyPath", func(t *testing.T) {
		var want error

		var cmder = mocks.NewMockCommander()
		cmder.On("Run", mock.Anything, mock.Anything).Return(want)

		var service = NewCLIService(cmder)
		var got = service.CreateTag(testTag)

		assert.NoError(t, got)
	})

	t.Run("ReturnErrorOnError", func(t *testing.T) {
		var want = errors.New("some-error")

		var cmder = mocks.NewMockCommander()
		cmder.On("Run", mock.Anything, mock.Anything).Return(want)

		var gitService = NewCLIService(cmder)
		var got = gitService.CreateTag(testTag)

		assert.Error(t, got)
	})
}

func TestCLIService_GetLatestCommitMessage(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		var want = "[fix] hello message"

		var cmder = mocks.NewMockCommander()
		cmder.On("Output", mock.Anything, mock.Anything).Return(want, nil)

		var gitService = NewCLIService(cmder)
		var got, err = gitService.GetLatestCommitMessage()

		assert.NoError(t, err)
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})

	t.Run("ReturnErrorOnError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var cmder = mocks.NewMockCommander()
		cmder.On("Output", mock.Anything, mock.Anything).Return("", want)

		var gitService = NewCLIService(cmder)
		var _, err = gitService.GetLatestCommitMessage()

		assert.Error(t, err)
	})
}

func TestCLIService_GetTag(t *testing.T) {
	type Test struct {
		Err  error
		Name string
		Want string
	}

	var tests = []Test{
		{Name: "HappyPath", Err: nil, Want: "1.0.0"},
		{Name: "ReturnDefaultTagOnError", Err: errors.New("some-error"), Want: DefaultTag},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var cmder = mocks.NewMockCommander()
			cmder.On("Output", mock.Anything, mock.Anything).Return(test.Want, test.Err)

			var gitService = NewCLIService(cmder)

			var want = test.Want
			var got = gitService.GetTag()

			assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
		})
	}
}
