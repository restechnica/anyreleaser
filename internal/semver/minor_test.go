package semver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinorStrategy_MinorConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "minor"
		var got = Minor

		assert.Equal(t, want, got, fmt.Sprintf(`want: "%s", got: "%s"`, want, got))
	})
}

func TestMinorStrategy_GetLevel(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		var strategy = NewMinorStrategy()

		var want = Minor
		var got, err = strategy.GetLevel()

		assert.NoError(t, err)
		assert.Equal(t, want, got, fmt.Sprintf(`want: "%s", got: "%s"`, want, got))
	})
}
