package tui

import "strings"

func helpView() string {
	lines := []struct{ key, desc string }{
		{"↑/k   ↓/j", "move"},
		{"/", "filter"},
		{"a", "add"},
		{"e", "edit"},
		{"d", "delete"},
		{"enter", "copy command to clipboard"},
		{"?", "toggle help"},
		{"q / ctrl+c", "quit"},
		{"", ""},
		{"Form:", ""},
		{"tab / shift+tab", "next/prev field"},
		{"enter", "next field (last = submit)"},
		{"ctrl+s", "submit"},
		{"esc", "cancel"},
	}
	var b strings.Builder
	b.WriteString(styleTitle.Render("aka — keys"))
	b.WriteString("\n\n")
	for _, l := range lines {
		if l.key == "" && l.desc == "" {
			b.WriteString("\n")
			continue
		}
		if l.desc == "" {
			b.WriteString(styleTitle.Render(l.key))
			b.WriteString("\n")
			continue
		}
		b.WriteString(styleInput.Render(l.key))
		b.WriteString("  ")
		b.WriteString(styleHint.Render(l.desc))
		b.WriteString("\n")
	}
	b.WriteString("\n")
	b.WriteString(styleHint.Render("press any key to return"))
	return styleFrame.Render(b.String())
}
