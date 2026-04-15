package tui

import "github.com/charmbracelet/lipgloss"

// palette is the set of colors used by every TUI screen. Themes are built
// by swapping palettes; the lipgloss styles are rebuilt from the active
// palette.
type palette struct {
	Accent lipgloss.TerminalColor
	Good   lipgloss.TerminalColor
	Danger lipgloss.TerminalColor
	Muted  lipgloss.TerminalColor
}

func paletteFor(theme string) palette {
	switch theme {
	case "dracula":
		return palette{
			Accent: lipgloss.Color("#FF79C6"),
			Good:   lipgloss.Color("#50FA7B"),
			Danger: lipgloss.Color("#FF5555"),
			Muted:  lipgloss.AdaptiveColor{Light: "#6272A4", Dark: "#6272A4"},
		}
	case "nord":
		return palette{
			Accent: lipgloss.Color("#88C0D0"),
			Good:   lipgloss.Color("#A3BE8C"),
			Danger: lipgloss.Color("#BF616A"),
			Muted:  lipgloss.AdaptiveColor{Light: "#4C566A", Dark: "#D8DEE9"},
		}
	case "gruvbox":
		return palette{
			Accent: lipgloss.Color("#FABD2F"),
			Good:   lipgloss.Color("#B8BB26"),
			Danger: lipgloss.Color("#FB4934"),
			Muted:  lipgloss.AdaptiveColor{Light: "#928374", Dark: "#A89984"},
		}
	}
	// default — aka brand.
	return palette{
		Accent: lipgloss.Color("#00B5E2"),
		Good:   lipgloss.Color("#22D3A6"),
		Danger: lipgloss.Color("#D22128"),
		Muted:  lipgloss.AdaptiveColor{Light: "#6c6c6c", Dark: "#9a9a9a"},
	}
}

var (
	styleTitle  lipgloss.Style
	styleHint   lipgloss.Style
	styleDanger lipgloss.Style
	styleOK     lipgloss.Style
	styleInput  lipgloss.Style
	styleFrame  lipgloss.Style
)

// applyTheme rebuilds the package-level lipgloss styles from the named theme.
// It is called once during TUI bootstrap.
func applyTheme(theme string) {
	p := paletteFor(theme)
	styleTitle = lipgloss.NewStyle().Bold(true).Foreground(p.Accent)
	styleHint = lipgloss.NewStyle().Foreground(p.Muted)
	styleDanger = lipgloss.NewStyle().Foreground(p.Danger).Bold(true)
	styleOK = lipgloss.NewStyle().Foreground(p.Good)
	styleInput = lipgloss.NewStyle().Foreground(p.Accent)
	styleFrame = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(p.Accent).
		Padding(1, 2)
}

// AvailableThemes is the canonical list of theme names. Exposed so future
// commands (e.g. `aka config theme <name>`) can validate input.
var AvailableThemes = []string{"default", "dracula", "nord", "gruvbox"}

func init() {
	// Sensible default so callers that forget applyTheme don't crash on nil
	// lipgloss styles.
	applyTheme("default")
}
