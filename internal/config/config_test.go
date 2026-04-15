package config

import (
	"path/filepath"
	"testing"
)

func TestLoadMissingReturnsDefaults(t *testing.T) {
	cfg, err := Load(filepath.Join(t.TempDir(), "nope.toml"))
	if err != nil {
		t.Fatalf("load missing: %v", err)
	}
	if cfg.Language != "auto" {
		t.Fatalf("want Language=auto, got %q", cfg.Language)
	}
	if len(cfg.DefaultShells) != 3 {
		t.Fatalf("want 3 default shells, got %v", cfg.DefaultShells)
	}
}

func TestSaveLoadRoundTrip(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.toml")
	in := Config{
		Language:      "es",
		DefaultShells: []string{"zsh"},
		ConfirmDelete: false,
		Theme:         "dracula",
	}
	if err := Save(path, in); err != nil {
		t.Fatalf("save: %v", err)
	}
	out, err := Load(path)
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if out.Language != "es" || out.Theme != "dracula" || out.ConfirmDelete {
		t.Fatalf("round-trip mismatch: %+v", out)
	}
	if len(out.DefaultShells) != 1 || out.DefaultShells[0] != "zsh" {
		t.Fatalf("round-trip shells mismatch: %v", out.DefaultShells)
	}
}

func TestPathsHonorEnv(t *testing.T) {
	t.Setenv("AKA_CONFIG", "/x/config.toml")
	t.Setenv("AKA_DATA", "/x/aliases.json")
	t.Setenv("AKA_OUTDIR", "/x/out")
	cp, _ := ConfigPath()
	ap, _ := AliasesPath()
	od, _ := OutDir()
	if cp != "/x/config.toml" || ap != "/x/aliases.json" || od != "/x/out" {
		t.Fatalf("env not honored: %q %q %q", cp, ap, od)
	}
}
