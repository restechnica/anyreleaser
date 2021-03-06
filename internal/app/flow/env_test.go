package flow

import (
	"fmt"
	"os"
	"testing"

	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/restechnica/anyreleaser/internal/app/config"
	"github.com/restechnica/anyreleaser/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoadEnvFiles_Run(t *testing.T) {
	t.Run("LoadSeveralEnvFilesAndCheckEnvVars", func(t *testing.T) {
		var pipe = LoadEnvFiles{}

		var cfg = config.NewRoot()
		cfg.Env.Files = []string{
			"../../../test/env/file1.env",
			"../../../test/env/file2.env",
		}

		var ctx = &app.Context{Config: cfg}
		var err = pipe.Run(ctx)

		assert.NoError(t, err)

		var want = "some_value"
		var got = os.Getenv("some_key")
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)

		want = "another_value"
		got = os.Getenv("another_key")
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})

	t.Run("LoadEmptyEnvFileReturnsNoError", func(t *testing.T) {
		var pipe = LoadEnvFiles{}

		var cfg = config.NewRoot()
		cfg.Env.Files = []string{"../../../test/env/empty.env"}

		var ctx = &app.Context{Config: cfg}
		var err = pipe.Run(ctx)

		assert.NoError(t, err)
	})

	t.Run("LoadInexistentEnvFileReturnsError", func(t *testing.T) {
		var pipe = LoadEnvFiles{}

		var cfg = config.NewRoot()
		cfg.Env.Files = []string{"../../../test/env/does-not-exist.env"}

		var ctx = &app.Context{Config: cfg}
		var err = pipe.Run(ctx)

		assert.Error(t, err)
	})

	t.Run("LoadFaultyEnvFileReturnsError", func(t *testing.T) {
		var pipe = LoadEnvFiles{}

		var cfg = config.NewRoot()
		cfg.Env.Files = []string{"../../../test/env/faulty.env"}

		var ctx = &app.Context{Config: cfg}
		var err = pipe.Run(ctx)

		assert.Error(t, err)
	})
}

func TestLoadEnvScripts_Run(t *testing.T) {
	t.Run("LoadSeveralEnvScriptsAndCheckEnvVars", func(t *testing.T) {
		var pipe = LoadEnvScripts{}
		var cmder = mocks.NewMockCommander()

		cmder.On("Output", "python3", mock.Anything).Return("some_key=some_value", nil)
		cmder.On("Output", "node", mock.Anything).Return("another_key=another_value", nil)

		var cfg = config.NewRoot()
		cfg.Env.Scripts = []config.EnvScript{
			{Bin: "python3", Path: "some-script"},
			{Bin: "node", Path: "some-script"},
		}

		var ctx = &app.Context{Commander: cmder, Config: cfg}
		var err = pipe.Run(ctx)

		assert.NoError(t, err)

		var want = "some_value"
		var got = os.Getenv("some_key")
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)

		want = "another_value"
		got = os.Getenv("another_key")
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})

	t.Run("LoadEmptyEnvScriptAndCheckEnvVars", func(t *testing.T) {
		var pipe = LoadEnvScripts{}
		var cmder = mocks.NewMockCommander()

		cmder.On("Output", "python3", mock.Anything).Return("", nil)

		var cfg = config.NewRoot()
		cfg.Env.Scripts = []config.EnvScript{{Bin: "python3", Path: "some-script"}}

		var ctx = &app.Context{Commander: cmder, Config: cfg}
		var err = pipe.Run(ctx)

		assert.NoError(t, err)
	})

	t.Run("LoadEnvScriptWithCommanderErrorAndReturnError", func(t *testing.T) {
		var pipe = LoadEnvScripts{}
		var cmder = mocks.NewMockCommander()

		var want = fmt.Errorf("failed to run 'python3 some-script'")
		cmder.On("Output", "python3", mock.Anything).Return("", want)

		var cfg = config.NewRoot()
		cfg.Env.Scripts = []config.EnvScript{{Bin: "python3", Path: "some-script"}}

		var ctx = &app.Context{Commander: cmder, Config: cfg}
		var got = pipe.Run(ctx)

		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})

	t.Run("LoadEnvScriptWithScriptOutputErrorAndReturnError", func(t *testing.T) {
		var pipe = LoadEnvScripts{}
		var cmder = mocks.NewMockCommander()

		var want = fmt.Errorf("failed to parse output from 'python3 some-script'")
		cmder.On("Output", "python3", mock.Anything).Return("faulty output", nil)

		var cfg = config.NewRoot()
		cfg.Env.Scripts = []config.EnvScript{{Bin: "python3", Path: "some-script"}}

		var ctx = &app.Context{Commander: cmder, Config: cfg}
		var got = pipe.Run(ctx)

		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})

	t.Run("LoadEnvScriptWithEmptyKeyOutputAndReturnError", func(t *testing.T) {
		var pipe = LoadEnvScripts{}
		var cmder = mocks.NewMockCommander()

		var want = fmt.Errorf("could not set env var '=some_value' from 'python3 some-script'")
		cmder.On("Output", "python3", mock.Anything).Return(" =some_value", nil)

		var cfg = config.NewRoot()
		cfg.Env.Scripts = []config.EnvScript{{Bin: "python3", Path: "some-script"}}

		var ctx = &app.Context{Commander: cmder, Config: cfg}
		var got = pipe.Run(ctx)

		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})
}

func TestLoadEnvVars_Run(t *testing.T) {
	t.Run("LoadSeveralEnvVarsAndCheckEnvVars", func(t *testing.T) {
		var pipe = LoadEnvVars{}

		var cfg = config.NewRoot()
		cfg.Env.Vars = map[string]string{
			"some_key":    "some_value",
			"another_key": "another_value",
		}

		var ctx = &app.Context{Config: cfg}
		var err = pipe.Run(ctx)

		assert.NoError(t, err)

		var want = "some_value"
		var got = os.Getenv("some_key")
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)

		want = "another_value"
		got = os.Getenv("another_key")
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})

	t.Run("LoadEmptyKeyVarAndReturnError", func(t *testing.T) {
		var want = fmt.Errorf("could not set env var '=some_value'")
		var pipe = LoadEnvVars{}

		var cfg = config.NewRoot()
		cfg.Env.Vars = map[string]string{"": "some_value"}

		var ctx = &app.Context{Config: cfg}
		var got = pipe.Run(ctx)

		assert.Error(t, want)
		assert.Equal(t, want, got, "want: '%s', got: '%s'", want, got)
	})
}
