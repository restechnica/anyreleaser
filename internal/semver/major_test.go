package semver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorStrategy_MajorConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "major"
		var got = Major

		assert.Equal(t, want, got, fmt.Sprintf(`want: "%s", got: "%s"`, want, got))
	})
}

func TestMajorStrategy_GetLevel(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		var strategy = NewMajorStrategy()

		var want = Major
		var got, err = strategy.GetLevel()

		assert.NoError(t, err)
		assert.Equal(t, want, got, fmt.Sprintf(`want: "%s", got: "%s"`, want, got))
	})
}
