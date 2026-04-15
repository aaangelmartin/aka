package emit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aaangelmartin/aka/internal/alias"
)

func TestRenderZsh(t *testing.T) {
	got := Render(alias.ShellZsh, []alias.Alias{
		{Name: "cc", Command: "claude"},
		{Name: "ccauto", Command: "claude --dangerously-skip-permissions"},
	})
	want := "alias cc='claude'"
	if !strings.Contains(got, want) {
		t.Fatalf("missing %q in output:\n%s", want, got)
	}
	if !strings.Contains(got, "alias ccauto='claude --dangerously-skip-permissions'") {
		t.Fatalf("missing ccauto line:\n%s", got)
	}
}

func TestRenderEscapesSingleQuotes(t *testing.T) {
	got := Render(alias.ShellBash, []alias.Alias{
		{Name: "say", Command: `echo 'hello'`},
	})
	want := `alias say='echo '\''hello'\'''`
	if !strings.Contains(got, want) {
		t.Fatalf("quote escape failed.\ngot:\n%s\nwant line: %s", got, want)
	}
}

func TestRenderFishSyntax(t *testing.T) {
	got := Render(alias.ShellFish, []alias.Alias{
		{Name: "gs", Command: "git status"},
	})
	if !strings.Contains(got, "alias gs 'git status'") {
		t.Fatalf("fish syntax wrong:\n%s", got)
	}
}

func TestRenderSkipsShellsNotTargeted(t *testing.T) {
	got := Render(alias.ShellFish, []alias.Alias{
		{Name: "only-zsh", Command: "do", Shells: []string{alias.ShellZsh}},
	})
	if strings.Contains(got, "only-zsh") {
		t.Fatalf("alias leaked into wrong shell:\n%s", got)
	}
}

func TestRenderIncludesDescription(t *testing.T) {
	got := Render(alias.ShellZsh, []alias.Alias{
		{Name: "x", Command: "y", Description: "multi\nline"},
	})
	if !strings.Contains(got, "# multi line") {
		t.Fatalf("description newline not sanitized:\n%s", got)
	}
}

func TestRegenerateWritesAllThreeFiles(t *testing.T) {
	dir := t.TempDir()
	err := Regenerate(dir, []alias.Alias{{Name: "hi", Command: "echo hi"}})
	if err != nil {
		t.Fatalf("regenerate: %v", err)
	}
	for _, sh := range alias.AllShells {
		path := filepath.Join(dir, "aliases."+sh)
		b, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read %s: %v", path, err)
		}
		if !strings.Contains(string(b), "hi") {
			t.Fatalf("%s missing alias:\n%s", path, b)
		}
	}
}
