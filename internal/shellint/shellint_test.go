package shellint

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInstallCreatesFile(t *testing.T) {
	dir := t.TempDir()
	rc := filepath.Join(dir, ".zshrc")
	out := filepath.Join(dir, "aka-out")

	action, err := Install(rc, "zsh", out)
	if err != nil {
		t.Fatalf("install: %v", err)
	}
	if action != "installed" {
		t.Fatalf("action = %q; want installed", action)
	}
	b, _ := os.ReadFile(rc)
	if !strings.Contains(string(b), BeginMarker) || !strings.Contains(string(b), EndMarker) {
		t.Fatalf("markers missing:\n%s", b)
	}
	if !strings.Contains(string(b), "aliases.zsh") {
		t.Fatalf("source line missing:\n%s", b)
	}
}

func TestInstallIdempotent(t *testing.T) {
	dir := t.TempDir()
	rc := filepath.Join(dir, ".zshrc")
	out := filepath.Join(dir, "aka-out")

	if _, err := Install(rc, "zsh", out); err != nil {
		t.Fatal(err)
	}
	action, err := Install(rc, "zsh", out)
	if err != nil {
		t.Fatal(err)
	}
	if action != "unchanged" {
		t.Fatalf("second install action = %q; want unchanged", action)
	}
}

func TestInstallPreservesUserContent(t *testing.T) {
	dir := t.TempDir()
	rc := filepath.Join(dir, ".zshrc")
	body := "export FOO=bar\nalias vim='nvim'\n"
	_ = os.WriteFile(rc, []byte(body), 0o644)

	if _, err := Install(rc, "zsh", dir); err != nil {
		t.Fatal(err)
	}
	got, _ := os.ReadFile(rc)
	if !strings.Contains(string(got), "export FOO=bar") || !strings.Contains(string(got), "alias vim='nvim'") {
		t.Fatalf("pre-existing content clobbered:\n%s", got)
	}
	if !strings.Contains(string(got), BeginMarker) {
		t.Fatalf("block missing:\n%s", got)
	}
}

func TestUninstall(t *testing.T) {
	dir := t.TempDir()
	rc := filepath.Join(dir, ".zshrc")
	body := "export FOO=bar\n"
	_ = os.WriteFile(rc, []byte(body), 0o644)
	if _, err := Install(rc, "zsh", dir); err != nil {
		t.Fatal(err)
	}
	action, err := Uninstall(rc)
	if err != nil {
		t.Fatal(err)
	}
	if action != "removed" {
		t.Fatalf("action = %q; want removed", action)
	}
	got, _ := os.ReadFile(rc)
	if strings.Contains(string(got), BeginMarker) {
		t.Fatalf("block not stripped:\n%s", got)
	}
	if !strings.Contains(string(got), "export FOO=bar") {
		t.Fatalf("user content lost:\n%s", got)
	}
}

func TestParseRC(t *testing.T) {
	rc := `
# user stuff
alias vim='nvim'
alias cc="claude"
alias bare=echo

# aka block should be skipped
# >>> aka >>>
alias ghost='should not appear'
# <<< aka <<<

alias gs='git '\''status'\'''
`
	got := ParseRC(rc)
	byName := map[string]string{}
	for _, a := range got {
		byName[a.Name] = a.Command
	}
	if byName["vim"] != "nvim" || byName["cc"] != "claude" || byName["bare"] != "echo" {
		t.Fatalf("basic parses missing: %+v", byName)
	}
	if _, bad := byName["ghost"]; bad {
		t.Fatalf("aka block leaked into parse: %+v", byName)
	}
	if byName["gs"] != "git 'status'" {
		t.Fatalf("escaped quote decode wrong: %q", byName["gs"])
	}
}
