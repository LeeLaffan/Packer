package internal

type Config struct {
	ZipFmt    string                `toml:"ZipFmt"`
	Packs     map[string]PackConfig `toml:"Packs"`
}

type PackConfig struct {
	Source   string   `toml:"Source"`
	Excludes []string `toml:"Excludes"`
}
