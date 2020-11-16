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

func TestMajorStrategy_Increment(t *testing.T) {
	type IncrementTest struct {
		Name          string
		TargetVersion string
		Want          string
	}

	var incrementTests = []IncrementTest{
		{Name: "HappyPath", TargetVersion: "1.0.0", Want: "2.0.0"},
		{Name: "ResetPatch", TargetVersion: "7.0.4", Want: "8.0.0"},
		{Name: "ResetMinor", TargetVersion: "6.8.0", Want: "7.0.0"},
		{Name: "DiscardPreBuild", TargetVersion: "2.0.0-pre+001", Want: "3.0.0"},
	}

	for _, test := range incrementTests {
		t.Run(test.Name, func(t *testing.T) {
			var want = test.Want
			var strategy = NewMajorStrategy()
			var got, err = strategy.Increment(test.TargetVersion)

			assert.NoError(t, err)
			assert.Equal(t, want, got, fmt.Sprintf(`want: %s, got: %s`, want, got))
		})
	}

	type IncrementErrorTest struct {
		Name          string
		TargetVersion string
	}

	var incrementErrorTests = []IncrementErrorTest{
		{Name: "ReturnErrorOnInvalidTargetVersion", TargetVersion: "invalid"},
		{Name: "ReturnErrorOnInvalidCharacter", TargetVersion: "v1.2.3"},
	}

	for _, test := range incrementErrorTests {
		t.Run(test.Name, func(t *testing.T) {
			var strategy = NewMajorStrategy()
			var _, got = strategy.Increment(test.TargetVersion)
			assert.Error(t, got)
		})
	}
}
