package semver

import (
	"testing"
)

type IncrementIn struct {
	Version string
	Number  string
}

type IncrementOut struct {
	Version string
}

type IncrementTest struct {
	Name string
	IncrementIn
	IncrementOut
}

var incrementTests = []IncrementTest{
	{"increment_patch", IncrementIn{"0.0.1", PATCH}, IncrementOut{"0.0.2"}},
	{"increment_minor_and_reset_patch", IncrementIn{"0.0.2", MINOR}, IncrementOut{"0.1.0"}},
	{"increment_major_and_reset_minor_patch", IncrementIn{"0.3.0", MAJOR}, IncrementOut{"1.0.0"}},
	{"increment_patch_discard_pre_build", IncrementIn{"0.0.1-pre+001", PATCH}, IncrementOut{"0.0.2"}},
	{"increment_minor_discard_pre_build", IncrementIn{"0.0.2-pre+001", MINOR}, IncrementOut{"0.1.0"}},
	{"increment_major_discard_pre_build", IncrementIn{"0.3.0-pre+001", MAJOR}, IncrementOut{"1.0.0"}},
	{"increment_patch_no_discard_minor_major", IncrementIn{"3.2.0", PATCH}, IncrementOut{"3.2.1"}},
	{"increment_minor_no_discard_major", IncrementIn{"3.0.0", MINOR}, IncrementOut{"3.1.0"}},
}

func TestSemverIncrement(t *testing.T) {
	for _, test := range incrementTests {
		t.Run(test.Name, func(t *testing.T) {
			var version string
			var err error

			version, err = Increment(test.IncrementIn.Version, test.IncrementIn.Number)

			if err != nil {
				t.Errorf("%+v\ngot error %s", test, err)
			} else if version != test.IncrementOut.Version {
				t.Errorf("%+v\nexpected %+v", test.IncrementIn, test.IncrementOut)
			}
		})
	}
}

type IncrementErrorTest struct {
	Name string
	IncrementIn
	Err string
}

var incrementErrorTests = []IncrementErrorTest{
	{"non_semver_input", IncrementIn{"help", PATCH}, `No Major.Minor.Patch elements found`},
	{"invalid_version_level", IncrementIn{"0.0.1", "invalid"}, `"invalid" is not a valid semver version level`},
	{"invalid_version_level", IncrementIn{"v0.0.1", "patch"}, `Invalid character(s) found in major number "v0"`},
}

func TestSemverIncrementErrors(t *testing.T) {
	for _, test := range incrementErrorTests {
		t.Run(test.Name, func(t *testing.T) {
			var err error

			_, err = Increment(test.IncrementIn.Version, test.IncrementIn.Number)

			if err == nil {
				t.Errorf("%+v\nthis should give an error \"%s\"", test, err)
			} else if err.Error() != test.Err {
				t.Errorf("%s\nexpected \"%+v\"", err, test.Err)
			}
		})
	}
}
