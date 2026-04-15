package tui

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/aaangelmartin/aka/internal/alias"
	"github.com/aaangelmartin/aka/internal/config"
	"github.com/aaangelmartin/aka/internal/i18n"
)

func (m *model) updateList(msg tea.Msg) (tea.Model, tea.Cmd) {
	km, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}
	if m.filterMode {
		return m.handleFilterKey(km)
	}
	switch km.String() {
	case "q", "ctrl+c":
		return m, tea.Quit
	case "?":
		m.screen = screenHelp
		return m, nil
	case "o":
		m.settings = newSettingsModel(m.cfg)
		m.screen = screenSettings
		return m, nil
	case "L":
		return m, m.toggleLanguage()
	case "j", "down":
		f := m.filteredItems()
		if m.cursor < len(f)-1 {
			m.cursor++
		}
	case "k", "up":
		if m.cursor > 0 {
			m.cursor--
		}
	case "g":
		m.cursor = 0
	case "G":
		if f := m.filteredItems(); len(f) > 0 {
			m.cursor = len(f) - 1
		}
	case "/":
		m.filterMode = true
	case "a":
		m.openForm(formAdd)
		return m, m.form.focusFirst()
	case "e":
		f := m.filteredItems()
		if m.cursor < len(f) {
			m.openForm(formEdit)
			m.form.loadFrom(f[m.cursor])
			return m, m.form.focusFirst()
		}
	case "d", "x":
		f := m.filteredItems()
		if m.cursor < len(f) {
			m.confirmTarget = f[m.cursor]
			m.confirmYes = false
			m.screen = screenConfirm
		}
	case "y", "enter":
		f := m.filteredItems()
		if m.cursor < len(f) {
			if err := clipboard.WriteAll(f[m.cursor].Command); err != nil {
				m.setStatus(i18n.Tf("tui.status.copyfail", err.Error()))
			} else {
				m.setStatus(i18n.Tf("tui.status.copied", f[m.cursor].Command))
			}
		}
	case "t":
		f := m.filteredItems()
		if m.tagFilter != "" {
			m.tagFilter = ""
			m.setStatus(i18n.T("tui.status.tag_clear"))
		} else if m.cursor < len(f) && len(f[m.cursor].Tags) > 0 {
			m.tagFilter = f[m.cursor].Tags[0]
			m.cursor = 0
			m.setStatus(i18n.Tf("tui.status.tag_set", m.tagFilter))
		}
	case "esc":
		if m.filter != "" || m.tagFilter != "" {
			m.filter = ""
			m.tagFilter = ""
			m.cursor = 0
		}
	}
	return m, nil
}

func (m *model) handleFilterKey(km tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch km.String() {
	case "esc":
		m.filterMode = false
		m.filter = ""
		m.cursor = 0
	case "enter":
		m.filterMode = false
	case "backspace":
		if len(m.filter) > 0 {
			m.filter = m.filter[:len(m.filter)-1]
			m.cursor = 0
		}
	default:
		if len(km.Runes) > 0 {
			m.filter += string(km.Runes)
			m.cursor = 0
		}
	}
	return m, nil
}

// toggleLanguage cycles EN ↔ ES and persists the choice to config.toml so it
// sticks between sessions (matches GoTo's L key).
func (m *model) toggleLanguage() tea.Cmd {
	next := "en"
	if i18n.Get() == i18n.EN {
		next = "es"
	}
	i18n.Set(i18n.Lang(next))
	m.cfg.Language = next
	if path, err := config.ConfigPath(); err == nil {
		_ = config.Save(path, m.cfg)
	}
	m.setStatus(i18n.Tf("tui.status.lang", next))
	return nil
}

// ---------- list view ----------

func (m *model) listView() string {
	if len(m.items) == 0 {
		return m.theme.Box.
			Width(m.innerWidth()-2).
			Height(m.innerHeight()).
			Align(lipgloss.Center, lipgloss.Center).
			Render(m.theme.Empty.Render(i18n.T("tui.empty")))
	}

	f := m.filteredItems()
	if len(f) == 0 {
		return m.theme.Box.
			Width(m.innerWidth()-2).
			Height(m.innerHeight()).
			Align(lipgloss.Center, lipgloss.Center).
			Render(m.theme.Empty.Render(i18n.T("tui.no_matches")))
	}

	if m.cursor >= len(f) {
		m.cursor = len(f) - 1
	}

	leftW := m.leftWidth()
	rightW := m.rightWidth()
	h := m.innerHeight()

	left := m.renderList(f, leftW, h)
	right := m.renderPreview(f[m.cursor], rightW, h)
	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

func (m *model) renderList(items []alias.Alias, width, height int) string {
	innerH := height - 2
	if innerH < 1 {
		innerH = 1
	}
	if m.cursor < m.offset {
		m.offset = m.cursor
	}
	if m.cursor >= m.offset+innerH {
		m.offset = m.cursor - innerH + 1
	}
	end := m.offset + innerH
	if end > len(items) {
		end = len(items)
	}

	var b strings.Builder
	for i := m.offset; i < end; i++ {
		a := items[i]
		badge := m.shellsBadge(a)
		nameCol := truncate(a.Name, 16)
		cmdCol := truncate(a.Command, width-26)
		line := fmt.Sprintf("%s %-16s %s", badge, nameCol, cmdCol)
		if i == m.cursor {
			b.WriteString(m.theme.ItemSel.Width(width - 2).Render("▶ " + line))
		} else {
			b.WriteString(m.theme.Item.Render("  " + line))
		}
		b.WriteString("\n")
	}
	return m.theme.BoxFocused.Width(width - 2).Height(height).Render(b.String())
}

// shellsBadge returns a compact coloured badge describing which shell(s) the
// alias targets. All-shells (the common case) renders as a small diamond to
// avoid cluttering every row; single-shell entries show the shell name.
func (m *model) shellsBadge(a alias.Alias) string {
	shells := a.ShellsOrAll()
	if len(shells) == len(alias.AllShells) {
		return m.theme.Status.Render("◆ ")
	}
	if len(shells) == 1 {
		return m.theme.ShellBadge(shells[0], strings.ToUpper(shells[0])[:1]) + " "
	}
	// Two shells: show first letter of each, colour of first.
	initials := ""
	for _, s := range shells {
		initials += strings.ToUpper(s)[:1]
	}
	return m.theme.ShellBadge(shells[0], initials) + " "
}

func (m *model) renderPreview(a alias.Alias, width, height int) string {
	var b strings.Builder
	// Header: name + shells
	b.WriteString(m.theme.Title.Render(a.Name))
	b.WriteString("\n")
	shells := strings.Join(a.ShellsOrAll(), " · ")
	b.WriteString(m.theme.Subtitle.Render(shells))
	b.WriteString("\n\n")

	// Command block
	b.WriteString(m.theme.Cmd.Render("$ " + wrap(a.Command, width-4)))
	b.WriteString("\n\n")

	// Tags
	if len(a.Tags) > 0 {
		parts := make([]string, 0, len(a.Tags))
		for _, t := range a.Tags {
			parts = append(parts, m.theme.Tag.Render("#"+t))
		}
		b.WriteString(strings.Join(parts, " "))
		b.WriteString("\n\n")
	}

	// Description
	if a.Description != "" {
		b.WriteString(m.theme.Desc.Render(wrap(a.Description, width-4)))
		b.WriteString("\n\n")
	}

	// Status line: created + last used
	if !a.CreatedAt.IsZero() {
		b.WriteString(m.theme.Status.Render(i18n.Tf("tui.preview.created", a.CreatedAt.Format("2006-01-02"))))
	}
	if !a.LastUsed.IsZero() {
		b.WriteString("  ·  ")
		b.WriteString(m.theme.Status.Render(i18n.Tf("tui.preview.last", a.LastUsed.Format("2006-01-02 15:04"))))
	}

	return m.theme.Box.Width(width - 2).Height(height).Render(b.String())
}

// ---------- string helpers ----------

func truncate(s string, n int) string {
	if n <= 0 {
		return ""
	}
	if len(s) <= n {
		return s
	}
	if n < 3 {
		return s[:n]
	}
	return s[:n-1] + "…"
}

func wrap(s string, width int) string {
	if width <= 0 {
		return s
	}
	var b strings.Builder
	line := 0
	for i, w := range strings.Fields(s) {
		if line+len(w)+1 > width {
			b.WriteString("\n")
			line = 0
		} else if i > 0 {
			b.WriteString(" ")
			line++
		}
		b.WriteString(w)
		line += len(w)
	}
	return b.String()
}
