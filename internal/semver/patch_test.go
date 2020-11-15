package semver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatchStrategy_PatchConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "patch"
		var got = Patch

		assert.Equal(t, want, got, fmt.Sprintf(`want: "%s", got: "%s"`, want, got))
	})
}

func TestPatchStrategy_GetLevel(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		var strategy = NewPatchStrategy()

		var want = Patch
		var got, err = strategy.GetLevel()

		assert.NoError(t, err)
		assert.Equal(t, want, got, fmt.Sprintf(`want: "%s", got: "%s"`, want, got))
	})
}
