package flow

import (
	"testing"

	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/restechnica/anyreleaser/internal/app/config"
	"github.com/restechnica/anyreleaser/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSetGit_Run(t *testing.T) {
	t.Run("SetGitNoErrors", func(t *testing.T) {
		var pipe = SetGit{}

		var cmder = mocks.NewMockCommander()
		cmder.On("Run", mock.Anything, mock.Anything).Return(nil)

		var cfg = config.NewRoot()
		cfg.Git.Config = map[string]string{
			"user.name":  "test",
			"user.email": "test@test.test",
		}

		var ctx = &app.Context{Commander: cmder, Config: cfg}
		var err = pipe.Run(ctx)

		assert.NoError(t, err)
		cmder.AssertNumberOfCalls(t, "Run", 3)
	})

	// TODO currently this code does not return errors

	//t.Run("ReturnErrorOnGitConfigError", func(t *testing.T) {
	//	var want = fmt.Errorf("some-error")
	//	var pipe = SetGit{}
	//
	//	var cmder = NewTestGitCommanderMock()
	//	cmder.On("Run", "git", []string{"config", "user.name", "test"}).Return(want)
	//	cmder.On("Run", "git", mock.Anything, mock.Anything).Return(nil)
	//
	//	var cfg = config.NewRoot()
	//	cfg.Git.Config = map[string]string{
	//		"user.name": "test",
	//	}
	//
	//	var ctx = &app.Context{Commander: cmder, Config: cfg}
	//	var got = pipe.Run(ctx)
	//
	//	assert.Error(t, got)
	//	assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	//})
}
