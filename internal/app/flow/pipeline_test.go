package flow

import (
	"fmt"
	"testing"

	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testPipelineFakePipe struct {
	mock.Mock
}

func newTestPipelineFakePipe() *testPipelineFakePipe {
	return &testPipelineFakePipe{}
}

func (mock *testPipelineFakePipe) Run(ctx *app.Context) error {
	var args = mock.Called(ctx)
	return args.Error(0)
}

func TestPipeline_Add(t *testing.T) {
	t.Run("CheckAddedPipe", func(t *testing.T) {
		var pipeline = Pipeline{}
		var want = newTestPipelineFakePipe()

		pipeline.Add(want)
		var got = pipeline.Pipes[0]

		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})

	t.Run("NoErrorOnGoodRun", func(t *testing.T) {
		var want error = nil
		var pipeline = Pipeline{}

		var pipe = newTestPipelineFakePipe()
		pipe.On("Run", mock.Anything).Return(want)

		pipeline.Add(pipe)
		var got = pipeline.Run(&app.Context{})

		assert.NoError(t, got)
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})

	t.Run("ReturnErrorOnPipeError", func(t *testing.T) {
		var want = fmt.Errorf("some-error")
		var pipeline = Pipeline{}

		var pipe = newTestPipelineFakePipe()
		pipe.On("Run", mock.Anything).Return(want)

		pipeline.Add(pipe)
		var got = pipeline.Run(&app.Context{})

		assert.Error(t, got)
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})
}
