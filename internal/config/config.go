// Package config loads and saves user configuration (TOML) and resolves XDG paths.
package config

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"

	"github.com/aaangelmartin/aka/internal/alias"
)

// Config is the user-facing configuration.
//
// TOML layout:
//
//	language        = "auto"                   # auto | en | es
//	default_shells  = ["zsh", "bash", "fish"]  # shells emitted when an alias
//	                                           # does not specify its own
//	confirm_delete  = true
//	theme           = "default"
type Config struct {
	Language       string   `toml:"language"`
	DefaultShells  []string `toml:"default_shells"`
	ConfirmDelete  bool     `toml:"confirm_delete"`
	Theme          string   `toml:"theme"`
}

// Default returns the zero-value-safe defaults.
func Default() Config {
	shells := make([]string, len(alias.AllShells))
	copy(shells, alias.AllShells)
	return Config{
		Language:      "auto",
		DefaultShells: shells,
		ConfirmDelete: true,
		Theme:         "default",
	}
}

// Load reads the config file, returning defaults if missing.
func Load(path string) (Config, error) {
	cfg := Default()
	b, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return cfg, nil
	}
	if err != nil {
		return cfg, err
	}
	if _, err := toml.Decode(string(b), &cfg); err != nil {
		return cfg, err
	}
	if len(cfg.DefaultShells) == 0 {
		cfg.DefaultShells = Default().DefaultShells
	}
	return cfg, nil
}

// Save writes the config as TOML.
func Save(path string, cfg Config) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return toml.NewEncoder(f).Encode(cfg)
}

// ConfigPath returns the path to config.toml, honoring $AKA_CONFIG and
// $XDG_CONFIG_HOME.
func ConfigPath() (string, error) {
	if x := os.Getenv("AKA_CONFIG"); x != "" {
		return x, nil
	}
	base := os.Getenv("XDG_CONFIG_HOME")
	if base == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		base = filepath.Join(home, ".config")
	}
	return filepath.Join(base, "aka", "config.toml"), nil
}

// AliasesPath returns the path to aliases.json (source of truth), honoring
// $AKA_DATA and $XDG_DATA_HOME.
func AliasesPath() (string, error) {
	if x := os.Getenv("AKA_DATA"); x != "" {
		return x, nil
	}
	base := os.Getenv("XDG_DATA_HOME")
	if base == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		base = filepath.Join(home, ".local", "share")
	}
	return filepath.Join(base, "aka", "aliases.json"), nil
}

// OutDir returns the directory where aliases.{zsh,bash,fish} are emitted,
// honoring $AKA_OUTDIR and $XDG_CONFIG_HOME.
func OutDir() (string, error) {
	if x := os.Getenv("AKA_OUTDIR"); x != "" {
		return x, nil
	}
	base := os.Getenv("XDG_CONFIG_HOME")
	if base == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		base = filepath.Join(home, ".config")
	}
	return filepath.Join(base, "aka"), nil
}
