package semver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevels_Constants(t *testing.T) {
	type LevelTest struct {
		Name string
		Got  string
		Want string
	}

	var LevelsTests = []LevelTest{
		{"CheckPatchLevelConstant", Patch, "patch"},
		{"CheckMinorLevelConstant", Minor, "minor"},
		{"CheckMajorLevelConstant", Major, "major"},
	}

	for _, test := range LevelsTests {
		t.Run(test.Name, func(t *testing.T) {
			var got = test.Got
			var want = test.Want

			assert.Equal(t, got, want, fmt.Sprintf(`want: "%s", got: "%s"`, want, got))
		})
	}
}
