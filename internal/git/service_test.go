package git

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceCommanderMock struct {
	mock.Mock
}

func NewServiceCommanderMock() *serviceCommanderMock {
	return &serviceCommanderMock{}
}

func (m *serviceCommanderMock) Output(name string, arg ...string) (string, error) {
	args := m.Called(name, arg)
	return args.String(0), args.Error(1)
}

func (m *serviceCommanderMock) Run(name string, arg ...string) error {
	args := m.Called(name, arg)
	return args.Error(0)
}

func TestService_CreateTag(t *testing.T) {
	const testTag = "1.2.3"

	t.Run("HappyPath", func(t *testing.T) {
		var want error

		var commander = NewServiceCommanderMock()
		commander.On("Run", mock.Anything, mock.Anything).Return(want)

		var service = NewService(commander)
		var got = service.CreateTag(testTag)

		assert.NoError(t, got)
	})

	t.Run("ReturnErrorOnError", func(t *testing.T) {
		var want = errors.New("some-error")

		var commander = new(serviceCommanderMock)
		commander.On("Run", mock.Anything, mock.Anything).Return(want)

		var gitService = NewService(commander)
		var got = gitService.CreateTag(testTag)

		assert.Error(t, got)
	})
}

func TestService_GetLatestCommitMessage(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		var want = "[fix] hello message"

		var commander = NewServiceCommanderMock()
		commander.On("Output", mock.Anything, mock.Anything).Return(want, nil)

		var gitService = NewService(commander)
		var got, err = gitService.GetLatestCommitMessage()

		assert.NoError(t, err)
		assert.Equal(t, want, got, fmt.Sprintf("want: %s, got: %s", want, got))
	})

	t.Run("ReturnErrorOnError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")

		var commander = NewServiceCommanderMock()
		commander.On("Output", mock.Anything, mock.Anything).Return("", want)

		var gitService = NewService(commander)
		var _, err = gitService.GetLatestCommitMessage()

		assert.Error(t, err)
	})
}

func TestService_GetTag(t *testing.T) {
	type GetTagTest struct {
		Err  error
		Name string
		Want string
	}

	var getTagTests = []GetTagTest{
		{Name: "HappyPath", Err: nil, Want: "1.0.0"},
		{Name: "ReturnDefaultTagOnError", Err: errors.New("some-error"), Want: DefaultTag},
	}

	for _, test := range getTagTests {
		t.Run(test.Name, func(t *testing.T) {
			var commander = NewServiceCommanderMock()
			commander.On("Output", mock.Anything, mock.Anything).Return(test.Want, test.Err)

			var gitService = NewService(commander)

			var got = gitService.GetTag()
			var want = test.Want

			assert.Equal(t, want, got, fmt.Sprintf("Wanted %s and got %s", want, got))
		})
	}
}
