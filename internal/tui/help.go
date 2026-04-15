package tui

import (
	"strings"

	"github.com/aaangelmartin/aka/internal/i18n"
)

func (m *model) helpView() string {
	rows := [][2]string{
		{"↑/k, ↓/j", i18n.T("help.move")},
		{"g, G", i18n.T("help.jump")},
		{"/", i18n.T("help.filter")},
		{"esc", i18n.T("help.esc")},
		{"a", i18n.T("help.add")},
		{"e", i18n.T("help.edit")},
		{"d, x", i18n.T("help.delete")},
		{"enter, y", i18n.T("help.copy")},
		{"t", i18n.T("help.tag")},
		{"o", i18n.T("help.settings")},
		{"L", i18n.T("help.lang")},
		{"?", i18n.T("help.toggle")},
		{"q, ctrl+c", i18n.T("help.quit")},
	}
	var b strings.Builder
	b.WriteString(m.theme.Title.Render(i18n.T("tui.help.title")))
	b.WriteString("\n\n")
	for _, r := range rows {
		b.WriteString(m.theme.Key.Render(padRight(r[0], 14)))
		b.WriteString("  ")
		b.WriteString(m.theme.Item.Render(r[1]))
		b.WriteString("\n")
	}
	b.WriteString("\n")
	b.WriteString(m.theme.Subtitle.Render(i18n.T("help.theme")))
	b.WriteString(" ")
	b.WriteString(m.theme.Status.Render(m.theme.Name))
	b.WriteString("   ")
	b.WriteString(m.theme.Subtitle.Render(i18n.T("help.lang_cur")))
	b.WriteString(" ")
	b.WriteString(m.theme.Status.Render(string(i18n.Get())))
	return m.theme.BoxFocused.Width(m.innerWidth() - 2).Height(m.innerHeight()).Render(b.String())
}

func padRight(s string, n int) string {
	if len(s) >= n {
		return s
	}
	return s + strings.Repeat(" ", n-len(s))
}
