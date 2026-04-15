package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// runCLI invokes the aka Cobra tree in-process with the given args and
// returns (stdout, stderr, err). It sets every XDG override to t.TempDir()
// so tests are fully isolated.
func runCLI(t *testing.T, dirs struct{ cfg, data, out string }, args ...string) (string, string, error) {
	t.Helper()
	t.Setenv("AKA_CONFIG", dirs.cfg)
	t.Setenv("AKA_DATA", dirs.data)
	t.Setenv("AKA_OUTDIR", dirs.out)

	root := newRoot()
	var stdout, stderr bytes.Buffer
	root.SetOut(&stdout)
	root.SetErr(&stderr)
	root.SetArgs(args)
	err := root.Execute()
	return stdout.String(), stderr.String(), err
}

func TestEndToEnd(t *testing.T) {
	base := t.TempDir()
	dirs := struct{ cfg, data, out string }{
		cfg:  filepath.Join(base, "config.toml"),
		data: filepath.Join(base, "aliases.json"),
		out:  filepath.Join(base, "out"),
	}

	// 1. add
	if out, _, err := runCLI(t, dirs, "add", "cc", "claude"); err != nil {
		t.Fatalf("add: %v (%s)", err, out)
	}
	if out, _, err := runCLI(t, dirs, "add", "-d", "auto", "ccauto", "claude --dangerously-skip-permissions"); err != nil {
		t.Fatalf("add ccauto: %v (%s)", err, out)
	}

	// 2. ls shows both
	stdout, _, err := runCLI(t, dirs, "ls")
	if err != nil {
		t.Fatalf("ls: %v", err)
	}
	if !strings.Contains(stdout, "cc") || !strings.Contains(stdout, "ccauto") {
		t.Fatalf("ls output missing aliases:\n%s", stdout)
	}

	// 3. generated aliases.zsh contains both
	zshPath := filepath.Join(dirs.out, "aliases.zsh")
	b, err := os.ReadFile(zshPath)
	if err != nil {
		t.Fatalf("read aliases.zsh: %v", err)
	}
	if !strings.Contains(string(b), "alias cc='claude'") {
		t.Fatalf("aliases.zsh missing cc:\n%s", b)
	}
	if !strings.Contains(string(b), "alias ccauto='claude --dangerously-skip-permissions'") {
		t.Fatalf("aliases.zsh missing ccauto:\n%s", b)
	}

	// 4. edit
	if out, _, err := runCLI(t, dirs, "edit", "-c", "git status --short", "cc"); err == nil {
		// cc's command should change but its name stays; we didn't change
		// the name, so this must succeed.
		_ = out
	} else {
		t.Fatalf("edit: %v", err)
	}

	// 5. rm (with -y to skip confirm)
	if out, _, err := runCLI(t, dirs, "rm", "-y", "ccauto"); err != nil {
		t.Fatalf("rm: %v (%s)", err, out)
	}

	// 6. aliases.zsh no longer has ccauto
	b, err = os.ReadFile(zshPath)
	if err != nil {
		t.Fatalf("reread aliases.zsh: %v", err)
	}
	if strings.Contains(string(b), "ccauto") {
		t.Fatalf("ccauto should have been removed:\n%s", b)
	}
	if !strings.Contains(string(b), "git status --short") {
		t.Fatalf("edit didn't propagate:\n%s", b)
	}

	// 7. export + import round-trip
	export, _, err := runCLI(t, dirs, "export")
	if err != nil {
		t.Fatalf("export: %v", err)
	}
	dump := filepath.Join(base, "dump.json")
	if err := os.WriteFile(dump, []byte(export), 0o644); err != nil {
		t.Fatal(err)
	}
	// Wipe and reimport.
	if err := os.Remove(dirs.data); err != nil {
		t.Fatal(err)
	}
	if out, _, err := runCLI(t, dirs, "import", dump); err != nil {
		t.Fatalf("import: %v (%s)", err, out)
	}
	stdout, _, err = runCLI(t, dirs, "ls")
	if err != nil || !strings.Contains(stdout, "cc") {
		t.Fatalf("post-import ls missing cc: %v / %s", err, stdout)
	}
}

func TestInstallUninstallCycle(t *testing.T) {
	base := t.TempDir()
	dirs := struct{ cfg, data, out string }{
		cfg:  filepath.Join(base, "config.toml"),
		data: filepath.Join(base, "aliases.json"),
		out:  filepath.Join(base, "out"),
	}
	rc := filepath.Join(base, ".zshrc")
	if err := os.WriteFile(rc, []byte("export FOO=bar\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if out, _, err := runCLI(t, dirs, "install", "zsh", "--rc", rc); err != nil {
		t.Fatalf("install: %v (%s)", err, out)
	}
	content, _ := os.ReadFile(rc)
	if !strings.Contains(string(content), "# >>> aka >>>") {
		t.Fatalf("install did not add marker block:\n%s", content)
	}
	if !strings.Contains(string(content), "export FOO=bar") {
		t.Fatalf("install clobbered user content:\n%s", content)
	}

	// Second install is a no-op.
	stdout, _, err := runCLI(t, dirs, "install", "zsh", "--rc", rc)
	if err != nil {
		t.Fatalf("second install: %v", err)
	}
	if !strings.Contains(stdout, "unchanged") {
		t.Fatalf("second install should be unchanged:\n%s", stdout)
	}

	// Uninstall restores the original content.
	if out, _, err := runCLI(t, dirs, "uninstall", "zsh", "--rc", rc); err != nil {
		t.Fatalf("uninstall: %v (%s)", err, out)
	}
	content, _ = os.ReadFile(rc)
	if strings.Contains(string(content), "# >>> aka >>>") {
		t.Fatalf("uninstall left markers behind:\n%s", content)
	}
	if !strings.Contains(string(content), "export FOO=bar") {
		t.Fatalf("uninstall dropped user content:\n%s", content)
	}
}

func TestImportFromRC(t *testing.T) {
	base := t.TempDir()
	dirs := struct{ cfg, data, out string }{
		cfg:  filepath.Join(base, "config.toml"),
		data: filepath.Join(base, "aliases.json"),
		out:  filepath.Join(base, "out"),
	}
	rc := filepath.Join(base, ".zshrc")
	rcBody := "alias vim='nvim'\nalias gs=\"git status\"\n"
	if err := os.WriteFile(rc, []byte(rcBody), 0o644); err != nil {
		t.Fatal(err)
	}
	if out, _, err := runCLI(t, dirs, "import", "--from-rc", rc); err != nil {
		t.Fatalf("import --from-rc: %v (%s)", err, out)
	}
	stdout, _, err := runCLI(t, dirs, "ls")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(stdout, "vim") || !strings.Contains(stdout, "gs") {
		t.Fatalf("rc parses missing from ls:\n%s", stdout)
	}
}
