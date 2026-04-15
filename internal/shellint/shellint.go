// Package shellint handles rc-file modifications and parsing.
//
// aka only ever touches the user's rc file to insert or remove a single
// block delimited by magic markers. Outside of install/uninstall, alias
// changes propagate by regenerating the sourced aliases.<shell> file.
package shellint

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/aaangelmartin/aka/internal/alias"
)

// BeginMarker / EndMarker delimit the aka-managed block inside rc files.
const (
	BeginMarker = "# >>> aka >>>"
	EndMarker   = "# <<< aka <<<"
)

// RCPath returns the default rc file for the given shell.
func RCPath(shell string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	switch shell {
	case alias.ShellZsh:
		return filepath.Join(home, ".zshrc"), nil
	case alias.ShellBash:
		return filepath.Join(home, ".bashrc"), nil
	case alias.ShellFish:
		return filepath.Join(home, ".config", "fish", "config.fish"), nil
	}
	return "", fmt.Errorf("unknown shell %q", shell)
}

// DetectShell returns the user's current shell based on $SHELL.
func DetectShell() string {
	s := strings.ToLower(filepath.Base(os.Getenv("SHELL")))
	switch {
	case strings.Contains(s, "zsh"):
		return alias.ShellZsh
	case strings.Contains(s, "bash"):
		return alias.ShellBash
	case strings.Contains(s, "fish"):
		return alias.ShellFish
	}
	return alias.ShellZsh // sensible default for most macOS / modern Linux users
}

// Block returns the rc snippet that sources aliases.<shell> from outDir.
func Block(shell, outDir string) string {
	var b strings.Builder
	fmt.Fprintln(&b, BeginMarker)
	fmt.Fprintln(&b, "# Added by aka. Do not edit manually — use `aka install` / `aka uninstall`.")
	fmt.Fprintf(&b, "[ -f %q ] && source %q\n", filepath.Join(outDir, "aliases."+shell), filepath.Join(outDir, "aliases."+shell))
	fmt.Fprintln(&b, EndMarker)
	return b.String()
}

// Install inserts (or replaces) the aka block in rcPath. It writes a
// timestamped backup next to the file on first install, then performs an
// atomic write (temp + rename). Returns the action performed: "installed",
// "updated", or "unchanged".
func Install(rcPath, shell, outDir string) (string, error) {
	existing, exists, err := readRC(rcPath)
	if err != nil {
		return "", err
	}
	block := Block(shell, outDir)
	var next string
	var action string
	if !exists {
		next = block
		action = "installed"
	} else if hasBlock(existing) {
		replaced := replaceBlock(existing, block)
		if replaced == existing {
			return "unchanged", nil
		}
		next = replaced
		action = "updated"
	} else {
		// Append, separated by a blank line for readability.
		sep := "\n"
		if !strings.HasSuffix(existing, "\n") {
			sep = "\n\n"
		}
		next = existing + sep + block
		action = "installed"
	}
	if exists && action != "unchanged" {
		if err := backup(rcPath); err != nil {
			return "", fmt.Errorf("backup rc: %w", err)
		}
	}
	if err := writeAtomic(rcPath, next); err != nil {
		return "", err
	}
	return action, nil
}

// Uninstall removes the aka block from rcPath. Returns:
//   - "removed"   — block was present and has been stripped
//   - "absent"    — rc file exists but had no aka block
//   - "no-rc"     — rc file does not exist
func Uninstall(rcPath string) (string, error) {
	existing, exists, err := readRC(rcPath)
	if err != nil {
		return "", err
	}
	if !exists {
		return "no-rc", nil
	}
	if !hasBlock(existing) {
		return "absent", nil
	}
	if err := backup(rcPath); err != nil {
		return "", fmt.Errorf("backup rc: %w", err)
	}
	next := stripBlock(existing)
	if err := writeAtomic(rcPath, next); err != nil {
		return "", err
	}
	return "removed", nil
}

func readRC(path string) (string, bool, error) {
	b, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	return string(b), true, nil
}

func hasBlock(s string) bool {
	return strings.Contains(s, BeginMarker) && strings.Contains(s, EndMarker)
}

// replaceBlock swaps the first BeginMarker…EndMarker span with newBlock.
// Callers must verify hasBlock() first.
func replaceBlock(s, newBlock string) string {
	start := strings.Index(s, BeginMarker)
	endMarkerIdx := strings.Index(s, EndMarker)
	if start < 0 || endMarkerIdx < 0 || endMarkerIdx <= start {
		return s
	}
	// end of EndMarker line:
	end := endMarkerIdx + len(EndMarker)
	if end < len(s) && s[end] == '\n' {
		end++
	}
	trimmed := strings.TrimRight(newBlock, "\n") + "\n"
	return s[:start] + trimmed + s[end:]
}

func stripBlock(s string) string {
	out := replaceBlock(s, "")
	// Collapse the double blank line that may remain after removal.
	out = strings.ReplaceAll(out, "\n\n\n", "\n\n")
	return out
}

func backup(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	dst := path + ".aka.bak"
	// Don't clobber an existing backup silently — suffix with .1, .2, …
	for i := 1; ; i++ {
		if _, err := os.Stat(dst); errors.Is(err, fs.ErrNotExist) {
			break
		}
		dst = fmt.Sprintf("%s.aka.bak.%d", path, i)
	}
	return os.WriteFile(dst, b, 0o644)
}

func writeAtomic(path, content string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), ".aka-rc-*.tmp")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)
	if _, err := tmp.WriteString(content); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmpPath, path)
}
