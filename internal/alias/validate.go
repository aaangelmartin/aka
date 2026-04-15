package alias

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// ErrEmptyName / ErrEmptyCommand / ErrInvalidName / ErrInvalidShell describe
// validation failures.
var (
	ErrEmptyName    = errors.New("alias name is empty")
	ErrEmptyCommand = errors.New("alias command is empty")
	ErrInvalidName  = errors.New("alias name has invalid characters")
	ErrInvalidShell = errors.New("unknown shell")
)

// nameRe accepts the conservative intersection of what zsh, bash, and fish
// accept for alias names: start with a letter or underscore, then
// letters / digits / underscores / dashes.
var nameRe = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_-]*$`)

// Validate returns a non-nil error describing why the alias is malformed,
// or nil if it is valid. Trailing/leading whitespace in Name is rejected.
func Validate(a Alias) error {
	if a.Name == "" {
		return ErrEmptyName
	}
	if strings.TrimSpace(a.Command) == "" {
		return ErrEmptyCommand
	}
	if !nameRe.MatchString(a.Name) {
		return fmt.Errorf("%w: %q", ErrInvalidName, a.Name)
	}
	for _, s := range a.Shells {
		if !IsValidShell(s) {
			return fmt.Errorf("%w: %q", ErrInvalidShell, s)
		}
	}
	return nil
}
