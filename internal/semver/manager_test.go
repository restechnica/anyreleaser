package semver

import (
	"testing"

	"github.com/restechnica/anyreleaser/internal/app/config"
	"github.com/restechnica/anyreleaser/internal/git"
	"github.com/stretchr/testify/assert"
)

func TestManager_GetStrategy(t *testing.T) {
	type Test struct {
		Name         string
		StrategyName string
		Want         Strategy
	}

	var tests = []Test{
		{Name: "GetPatchStrategy", StrategyName: "patch", Want: NewPatchStrategy()},
		{Name: "GetMinorStrategy", StrategyName: "minor", Want: NewMinorStrategy()},
		{Name: "GetMajorStrategy", StrategyName: "major", Want: NewMajorStrategy()},
		{Name: "GetGitCommitStrategy", StrategyName: "git-commit", Want: GitCommitStrategy{}},
		{Name: "GetAutoStrategy", StrategyName: "auto", Want: AutoStrategy{}},
		{Name: "GetDefaultStrategy", StrategyName: "wrong-name", Want: NewPatchStrategy()},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var fakeGitService = git.CLIService{}
			var manager = NewManager(config.Root{}, fakeGitService)
			var got = manager.GetStrategy(test.StrategyName)
			assert.IsType(t, test.Want, got, `want: "%s", got: "%s"`, test.Want, got)
		})
	}
}
