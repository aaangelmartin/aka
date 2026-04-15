package shellint

import (
	"bufio"
	"regexp"
	"strings"

	"github.com/aaangelmartin/aka/internal/alias"
)

// aliasRe matches `alias NAME=VALUE` on a single line (allowing leading
// whitespace). VALUE keeps its quoting so we can unwrap it deliberately.
var aliasRe = regexp.MustCompile(`^\s*alias\s+([A-Za-z_][A-Za-z0-9_-]*)=(.+?)\s*$`)

// fishAliasRe also accepts the space-separated form used by fish:
// `alias NAME 'VALUE'`.
var fishAliasRe = regexp.MustCompile(`^\s*alias\s+([A-Za-z_][A-Za-z0-9_-]*)\s+(['"])(.*)(['"])\s*$`)

// ParseRC reads a shell rc file and returns every alias definition it can
// recognise. Aliases nested inside the aka-managed block are skipped so that
// importing is safe after an install.
func ParseRC(content string) []alias.Alias {
	lines := bufio.NewScanner(strings.NewReader(content))
	lines.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	var (
		out   []alias.Alias
		inAka bool
	)
	for lines.Scan() {
		ln := lines.Text()
		trimmed := strings.TrimSpace(ln)
		if strings.HasPrefix(trimmed, BeginMarker) {
			inAka = true
			continue
		}
		if strings.HasPrefix(trimmed, EndMarker) {
			inAka = false
			continue
		}
		if inAka {
			continue
		}
		if strings.HasPrefix(trimmed, "#") || trimmed == "" {
			continue
		}
		if m := aliasRe.FindStringSubmatch(ln); m != nil {
			cmd, ok := unquote(m[2])
			if !ok {
				continue
			}
			out = append(out, alias.Alias{Name: m[1], Command: cmd})
			continue
		}
		if m := fishAliasRe.FindStringSubmatch(ln); m != nil && m[2] == m[4] {
			raw := m[3]
			// fish single-quoted strings treat backslash literally.
			out = append(out, alias.Alias{Name: m[1], Command: raw})
			continue
		}
	}
	return out
}

// unquote strips surrounding quotes from a shell-quoted value, expanding the
// POSIX single-quote escape `'\”`. Returns the value and whether we could
// parse the input.
func unquote(s string) (string, bool) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return "", false
	}
	if s[0] == '\'' && s[len(s)-1] == '\'' && len(s) >= 2 {
		inner := s[1 : len(s)-1]
		inner = strings.ReplaceAll(inner, `'\''`, "'")
		return inner, true
	}
	if s[0] == '"' && s[len(s)-1] == '"' && len(s) >= 2 {
		inner := s[1 : len(s)-1]
		inner = strings.ReplaceAll(inner, `\"`, `"`)
		inner = strings.ReplaceAll(inner, `\\`, `\`)
		return inner, true
	}
	// Unquoted — accept as-is.
	return s, true
}
