package alias

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	cases := []struct {
		name    string
		a       Alias
		wantErr error
	}{
		{"ok", Alias{Name: "gs", Command: "git status"}, nil},
		{"ok-with-dash", Alias{Name: "my-tool", Command: "do"}, nil},
		{"ok-underscore", Alias{Name: "_hidden", Command: "do"}, nil},
		{"empty-name", Alias{Name: "", Command: "do"}, ErrEmptyName},
		{"empty-command", Alias{Name: "x", Command: ""}, ErrEmptyCommand},
		{"whitespace-command", Alias{Name: "x", Command: "   "}, ErrEmptyCommand},
		{"leading-digit", Alias{Name: "1bad", Command: "do"}, ErrInvalidName},
		{"space-in-name", Alias{Name: "a b", Command: "do"}, ErrInvalidName},
		{"unknown-shell", Alias{Name: "x", Command: "do", Shells: []string{"ksh"}}, ErrInvalidShell},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := Validate(tc.a)
			if tc.wantErr == nil {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				return
			}
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("got %v; want wrap of %v", err, tc.wantErr)
			}
		})
	}
}

func TestTargetsShell(t *testing.T) {
	all := Alias{Name: "x", Command: "y"}
	if !all.TargetsShell(ShellZsh) || !all.TargetsShell(ShellBash) || !all.TargetsShell(ShellFish) {
		t.Fatalf("empty Shells should target every shell")
	}
	only := Alias{Name: "x", Command: "y", Shells: []string{ShellFish}}
	if only.TargetsShell(ShellZsh) || !only.TargetsShell(ShellFish) {
		t.Fatalf("unexpected shell targeting: %+v", only.ShellsOrAll())
	}
}
