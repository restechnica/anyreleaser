package flow

import (
	"github.com/restechnica/anyreleaser/internal/commands"
	"testing"

	"github.com/restechnica/anyreleaser/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestSetCommander_Run(t *testing.T) {
	t.Run("CheckCommander", func(t *testing.T) {
		var want = commands.NewExecCommander()
		var pipe = SetCommander{}
		var ctx = &app.Context{}

		var err = pipe.Run(ctx)
		var got = ctx.Commander

		assert.NoError(t, err)
		assert.Equal(t, want, got, "want: %s, got: %s", want, got)
	})
}
