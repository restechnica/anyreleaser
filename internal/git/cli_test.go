package git

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type cliServiceCommanderMock struct {
	mock.Mock
}

func NewCLIServiceCommanderMock() *cliServiceCommanderMock {
	return &cliServiceCommanderMock{}
}

func (m *cliServiceCommanderMock) Output(name string, arg ...string) (string, error) {
	args := m.Called(name, arg)
	return args.String(0), args.Error(1)
}

func (m *cliServiceCommanderMock) Run(name string, arg ...string) error {
	args := m.Called(name, arg)
	return args.Error(0)
}

func TestCLIService_CreateTag(t *testing.T) {
	const testTag = "1.2.3"

	t.Run("HappyPath", func(t *testing.T) {
		var want error

		var commander = NewCLIServiceCommanderMock()
		commander.On("Run", mock.Anything, mock.Anything).Return(want)

		var service = NewCLIService(commander)
		var got = service.CreateTag(testTag)

		assert.NoError(t, got)
	})

	t.Run("ReturnErrorOnError", func(t *testing.T) {
		var want = errors.New("some-error")

		var commander = NewCLIServiceCommanderMock()
		commander.On("Run", mock.Anything, mock.Anything).Return(want)

		var gitService = NewCLIService(commander)
		var got = gitService.CreateTag(testTag)

		assert.Error(t, got)
	})
}

func TestCLIService_GetLatestCommitMessage(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		var want = "[fix] hello message"

		var commander = NewCLIServiceCommanderMock()
		commander.On("Output", mock.Anything, mock.Anything).Return(want, nil)

		var gitService = NewCLIService(commander)
		var got, err = gitService.GetLatestCommitMessage()

		assert.NoError(t, err)
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})

	t.Run("ReturnErrorOnError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var commander = NewCLIServiceCommanderMock()
		commander.On("Output", mock.Anything, mock.Anything).Return("", want)

		var gitService = NewCLIService(commander)
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
			var commander = NewCLIServiceCommanderMock()
			commander.On("Output", mock.Anything, mock.Anything).Return(test.Want, test.Err)

			var gitService = NewCLIService(commander)

			var want = test.Want
			var got = gitService.GetTag()

			assert.Equal(t, want, got, `want: "%s", got: "%s"`, want, got)
		})
	}
}
