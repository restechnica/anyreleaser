package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRoot(t *testing.T) {
	t.Run("CheckDefaultValues", func(t *testing.T) {
		var config = NewRoot()
		assert.Equal(t, config.Env, Env{})
		assert.Equal(t, config.Git, Git{Unshallow: true})
		assert.Equal(t, config.Semver, Semver{Strategy: "auto", Matches: map[string]string{}})
	})
}
