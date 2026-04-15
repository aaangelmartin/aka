package tui

import "github.com/charmbracelet/lipgloss"

// Accent is the aka brand color (matches the shields.io badges).
var Accent = lipgloss.Color("#00B5E2")

// Muted is used for hints / secondary text.
var Muted = lipgloss.AdaptiveColor{Light: "#6c6c6c", Dark: "#9a9a9a"}

// Danger is used for destructive-action prompts.
var Danger = lipgloss.Color("#D22128")

// Good is used for success messages.
var Good = lipgloss.Color("#22D3A6")

var (
	styleTitle  = lipgloss.NewStyle().Bold(true).Foreground(Accent)
	styleHint   = lipgloss.NewStyle().Foreground(Muted)
	styleDanger = lipgloss.NewStyle().Foreground(Danger).Bold(true)
	styleOK     = lipgloss.NewStyle().Foreground(Good)
	styleInput  = lipgloss.NewStyle().Foreground(Accent)
	styleFrame  = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Accent).
			Padding(1, 2)
)
