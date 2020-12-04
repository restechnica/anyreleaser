package flow

import (
	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/restechnica/anyreleaser/internal/app/config"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"testing"
)

func TestPersistAppContext_Run(t *testing.T) {
	t.Run("CheckPersistedAppContext", func(t *testing.T) {
		var clictx = cli.Context{App: &cli.App{Metadata: map[string]interface{}{}}}
		var pipe = PersistAppContext{CLIContext: &clictx}

		var cfg = config.NewRoot()
		cfg.Env.Vars["some_var"] = "some_value"
		var want = &app.Context{Config: cfg}

		var err = pipe.Run(want)
		var got = clictx.App.Metadata[AppContext]

		assert.NoError(t, err)
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})
}
