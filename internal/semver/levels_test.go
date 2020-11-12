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
	{"patch_version_level", PATCH, "patch"},
	{"minor_version_level", MINOR, "minor"},
	{"major_version_level", MAJOR, "major"},
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
