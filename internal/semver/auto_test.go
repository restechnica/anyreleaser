package semver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAutoStrategy_AutoConstant(t *testing.T) {
	t.Run("CheckConstant", func(t *testing.T) {
		var want = "auto"
		var got = Auto

		assert.Equal(t, want, got, fmt.Sprintf(`want: "%s", got: "%s"`, want, got))
	})
}
