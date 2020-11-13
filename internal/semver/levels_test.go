package semver

import (
	"testing"
)

type LevelsTest struct {
	Name     string
	Constant string
	Value    string
}

var LevelsTests = []LevelsTest{
	{"patch_version_level", Patch, "patch"},
	{"minor_version_level", Minor, "minor"},
	{"major_version_level", Major, "major"},
}

func TestSemverNumbers(t *testing.T) {
	for _, test := range LevelsTests {
		t.Run(test.Name, func(t *testing.T) {
			if test.Constant != test.Value {
				t.Errorf("got \"%s\", expected \"%s\"", test.Constant, test.Value)
			}
		})
	}
}
