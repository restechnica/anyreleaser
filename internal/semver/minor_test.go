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

func TestMinorStrategy_Increment(t *testing.T) {
	type IncrementTest struct {
		Name          string
		TargetVersion string
		Want          string
	}

	var incrementTests = []IncrementTest{
		{Name: "HappyPath", TargetVersion: "0.1.0", Want: "0.2.0"},
		{Name: "ResetPatch", TargetVersion: "0.2.3", Want: "0.3.0"},
		{Name: "NoResetMajor", TargetVersion: "6.7.0", Want: "6.8.0"},
		{Name: "DiscardPreBuild", TargetVersion: "0.6.0-pre+001", Want: "0.7.0"},
	}

	for _, test := range incrementTests {
		t.Run(test.Name, func(t *testing.T) {
			var want = test.Want
			var strategy = NewMinorStrategy()
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
			var strategy = NewMinorStrategy()
			var _, got = strategy.Increment(test.TargetVersion)
			assert.Error(t, got)
		})
	}
}
