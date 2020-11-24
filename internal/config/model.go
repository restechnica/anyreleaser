package config

type Env struct {
	Files   []string          `yaml:"files,omitempty"`
	Scripts []EnvScript       `yaml:"scripts,omitempty"`
	Vars    map[string]string `yaml:"vars,omitempty"`
}

type EnvScript struct {
	Bin  string `yaml:"bin,omitempty"`
	Mode string `yaml:"mode,omitempty"`
	Path string `yaml:"path,omitempty"`
}

type Root struct {
	Env        Env        `yaml:"env,omitempty"`
	Semver     Semver     `yaml:"semver,omitempty"`
	Versioning Versioning `yaml:"version,omitempty"`
}

func NewRoot() (root Root) {
	return Root{
		Semver:     NewSemver(),
		Versioning: NewVersioning(),
	}
}

type Semver struct {
	Bin      string            `yaml:"bin,omitempty"`
	Strategy string            `yaml:"strategy,omitempty"`
	Matches  map[string]string `yaml:"matches,omitempty"`
	Path     string            `yaml:"path,omitempty"`
}

func NewSemver() Semver {
	return Semver{
		Strategy: "auto",
		Matches: map[string]string{
			`[fix]`:     "patch",
			`fix/`:      "patch",
			`[feature]`: "minor",
			`feature/`:  "minor",
			`[release]`: "major",
			`release/`:  "major",
		},
	}
}

type Versioning struct {
	Files []VersioningFile `yaml:"files,omitempty"`
	Git   VersioningGit    `yaml:"git,omitempty"`
}

func NewVersioning() Versioning {
	return Versioning{
		Git: NewVersioningGit(),
	}
}

type VersioningFile struct {
	Mode       string `yaml:"mode,omitempty"`
	Enabled    bool   `yaml:"enabled,omitempty"`
	Path       string `yaml:"path,omitempty"`
	Expression string `yaml:"expression,omitempty"`
}

type VersioningGit struct {
	Tags      bool `yaml:"tags,omitempty"`
	Unshallow bool `yaml:"unshallow,omitempty"`
}

func NewVersioningGit() VersioningGit {
	return VersioningGit{Tags: true, Unshallow: true}
}
