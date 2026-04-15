// Package alias defines the Alias model used across aka.
package alias

import "time"

// ShellZsh / ShellBash / ShellFish are the supported shells.
const (
	ShellZsh  = "zsh"
	ShellBash = "bash"
	ShellFish = "fish"
)

// AllShells lists every shell aka can emit aliases for.
var AllShells = []string{ShellZsh, ShellBash, ShellFish}

// Alias is a single user-defined shell alias.
//
// Name is the token the user types at the prompt. Command is what the shell
// substitutes in. Shells restricts which shells this alias is emitted for;
// an empty slice means "all shells" (normalized on load).
type Alias struct {
	Name        string    `json:"name"`
	Command     string    `json:"command"`
	Shells      []string  `json:"shells,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	LastUsed    time.Time `json:"last_used,omitempty"`
	HitCount    int       `json:"hit_count,omitempty"`
}

// ShellsOrAll returns the shells this alias targets, defaulting to every
// supported shell when the slice is empty.
func (a Alias) ShellsOrAll() []string {
	if len(a.Shells) == 0 {
		out := make([]string, len(AllShells))
		copy(out, AllShells)
		return out
	}
	out := make([]string, len(a.Shells))
	copy(out, a.Shells)
	return out
}

// TargetsShell reports whether the alias should be emitted for the given shell.
func (a Alias) TargetsShell(shell string) bool {
	for _, s := range a.ShellsOrAll() {
		if s == shell {
			return true
		}
	}
	return false
}

// IsValidShell returns true if s is one of the supported shell identifiers.
func IsValidShell(s string) bool {
	for _, v := range AllShells {
		if v == s {
			return true
		}
	}
	return false
}
