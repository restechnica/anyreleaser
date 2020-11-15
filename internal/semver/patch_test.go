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

func TestPatchStrategy_Increment(t *testing.T) {
	type IncrementTest struct {
		Name          string
		TargetVersion string
		Want          string
	}

	var incrementTests = []IncrementTest{
		{Name: "HappyPath", TargetVersion: "0.0.1", Want: "0.0.2"},
		{Name: "DiscardPreBuild", TargetVersion: "0.0.8-pre+001", Want: "0.0.9"},
		{Name: "NoResetMajorMinor", TargetVersion: "5.4.3", Want: "5.4.4"},
	}

	for _, test := range incrementTests {
		t.Run(test.Name, func(t *testing.T) {
			var want = test.Want
			var strategy = NewPatchStrategy()
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
			var strategy = NewPatchStrategy()
			var _, got = strategy.Increment(test.TargetVersion)
			assert.Error(t, got)
		})
	}
}
