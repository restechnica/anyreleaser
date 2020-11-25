package config

type Env struct {
	Files   []string          `yaml:"files,omitempty"`
	Scripts []EnvScript       `yaml:"scripts,omitempty"`
	Vars    map[string]string `yaml:"vars,omitempty"`
}

type EnvScript struct {
	Bin  string `yaml:"bin,omitempty"`
	Path string `yaml:"path,omitempty"`
}

type Git struct {
	Config    map[string]string `yaml:"config,omitempty"`
	Unshallow bool              `yaml:"unshallow,omitempty"`
}

func NewGit() Git {
	return Git{Unshallow: true}
}

type Root struct {
	Env    Env    `yaml:"env,omitempty"`
	Git    Git    `yaml:"git,omitempty"`
	Semver Semver `yaml:"semver,omitempty"`
}

func NewRoot() (root Root) {
	return Root{
		Git:    NewGit(),
		Semver: NewSemver(),
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
		Matches:  map[string]string{
			//`[fix]`:     "patch",
			//`fix/`:      "patch",
			//`[feature]`: "minor",
			//`feature/`:  "minor",
			//`[release]`: "major",
			//`release/`:  "major",
		},
	}
}
