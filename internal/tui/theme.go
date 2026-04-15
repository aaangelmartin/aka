// Package tui implements the interactive Bubble Tea TUI for aka.
//
// The theme is modelled after GoTo's: a single Theme struct owns every
// colour and every lipgloss style the screens use. Switching themes
// rebuilds the struct, so the rest of the TUI only ever reads
// m.theme.* values.
package tui

import "github.com/charmbracelet/lipgloss"

// Theme is a complete colour + style palette for one TUI theme.
type Theme struct {
	Name    string
	Accent  lipgloss.Color
	Accent2 lipgloss.Color
	FG      lipgloss.Color
	FGDim   lipgloss.Color
	BG      lipgloss.Color
	BGDim   lipgloss.Color
	Success lipgloss.Color
	Danger  lipgloss.Color
	Warning lipgloss.Color
	Border  lipgloss.Color

	Title      lipgloss.Style
	Subtitle   lipgloss.Style
	Item       lipgloss.Style
	ItemSel    lipgloss.Style
	Tag        lipgloss.Style
	Key        lipgloss.Style
	Desc       lipgloss.Style
	Box        lipgloss.Style
	BoxFocused lipgloss.Style
	Status     lipgloss.Style
	Danger_    lipgloss.Style
	Cmd        lipgloss.Style
	Help       lipgloss.Style
	Empty      lipgloss.Style

	// ShellStyles maps each shell identifier to a colour-coded badge style.
	ShellStyles map[string]lipgloss.Style
}

// ShellBadge wraps the given label in the style associated with the shell.
// Unknown shells fall back to a muted badge.
func (t Theme) ShellBadge(shell, label string) string {
	st, ok := t.ShellStyles[shell]
	if !ok {
		st = lipgloss.NewStyle().Foreground(t.FGDim)
	}
	return st.Render(label)
}

// ThemeByName returns the Theme struct for a theme identifier. Unknown names
// fall back to the aka-brand default (#00B5E2).
func ThemeByName(name string) Theme {
	switch name {
	case "dracula":
		return build("dracula",
			"#FF79C6", "#8BE9FD", "#F8F8F2", "#6272A4",
			"#282A36", "#44475A", "#50FA7B", "#FF5555", "#F1FA8C", "#44475A")
	case "nord":
		return build("nord",
			"#88C0D0", "#81A1C1", "#ECEFF4", "#4C566A",
			"#2E3440", "#3B4252", "#A3BE8C", "#BF616A", "#EBCB8B", "#434C5E")
	case "gruvbox":
		return build("gruvbox",
			"#FABD2F", "#83A598", "#EBDBB2", "#928374",
			"#282828", "#3C3836", "#B8BB26", "#FB4934", "#FE8019", "#504945")
	case "catppuccin", "catppuccin-mocha":
		return build("catppuccin",
			"#CBA6F7", "#89DCEB", "#CDD6F4", "#6C7086",
			"#1E1E2E", "#313244", "#A6E3A1", "#F38BA8", "#F9E2AF", "#45475A")
	case "tokyonight":
		return build("tokyonight",
			"#BB9AF7", "#7AA2F7", "#C0CAF5", "#565F89",
			"#1A1B26", "#24283B", "#9ECE6A", "#F7768E", "#E0AF68", "#3B4261")
	default:
		return build("default",
			"#00B5E2", "#7DD3FC", "#E5E7EB", "#6B7280",
			"#0B1220", "#1F2937", "#22D3A6", "#F87171", "#FBBF24", "#273043")
	}
}

// AvailableThemes is the canonical ordered list. Used by the settings cycle.
var AvailableThemes = []string{"default", "dracula", "nord", "gruvbox", "catppuccin", "tokyonight"}

func build(name, accent, accent2, fg, fgDim, bg, bgDim, ok, bad, warn, border string) Theme {
	t := Theme{
		Name:    name,
		Accent:  lipgloss.Color(accent),
		Accent2: lipgloss.Color(accent2),
		FG:      lipgloss.Color(fg),
		FGDim:   lipgloss.Color(fgDim),
		BG:      lipgloss.Color(bg),
		BGDim:   lipgloss.Color(bgDim),
		Success: lipgloss.Color(ok),
		Danger:  lipgloss.Color(bad),
		Warning: lipgloss.Color(warn),
		Border:  lipgloss.Color(border),
	}
	t.Title = lipgloss.NewStyle().Bold(true).Foreground(t.Accent)
	t.Subtitle = lipgloss.NewStyle().Foreground(t.Accent2)
	t.Item = lipgloss.NewStyle().Foreground(t.FG)
	t.ItemSel = lipgloss.NewStyle().Bold(true).Foreground(t.Accent).Background(t.BGDim)
	t.Tag = lipgloss.NewStyle().Foreground(t.Success)
	t.Key = lipgloss.NewStyle().Bold(true).Foreground(t.Accent2)
	t.Desc = lipgloss.NewStyle().Foreground(t.FGDim).Italic(true)
	t.Box = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(t.Border).Padding(0, 1)
	t.BoxFocused = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(t.Accent).Padding(0, 1)
	t.Status = lipgloss.NewStyle().Foreground(t.FGDim)
	t.Danger_ = lipgloss.NewStyle().Foreground(t.Danger).Bold(true)
	t.Cmd = lipgloss.NewStyle().Foreground(t.Accent2)
	t.Help = lipgloss.NewStyle().Foreground(t.FGDim)
	t.Empty = lipgloss.NewStyle().Foreground(t.FGDim).Italic(true).Align(lipgloss.Center)

	mk := func(fg string) lipgloss.Style {
		return lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(fg))
	}
	t.ShellStyles = map[string]lipgloss.Style{
		"zsh":  mk(accent),    // primary — aka's home shell
		"bash": mk(warn),      // warm yellow
		"fish": mk("#F472B6"), // pink, stable across themes
	}
	return t
}
