package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/aaangelmartin/aka/internal/config"
	"github.com/aaangelmartin/aka/internal/i18n"
)

type settingsRow int

const (
	rowLanguage settingsRow = iota
	rowTheme
	rowConfirmDelete
	rowCount
)

// settingsModel mirrors GoTo's simpler case: a small list of cycle-able rows
// backed by config.Config. Persisting happens on every change.
type settingsModel struct {
	cursor int
	langs  []string
	themes []string
}

func newSettingsModel(_ config.Config) settingsModel {
	return settingsModel{
		langs:  []string{"auto", "en", "es"},
		themes: AvailableThemes,
	}
}

func (m *model) updateSettings(msg tea.Msg) (tea.Model, tea.Cmd) {
	km, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}
	switch km.String() {
	case "esc", "q":
		m.screen = screenList
		return m, nil
	case "j", "down":
		if m.settings.cursor < int(rowCount)-1 {
			m.settings.cursor++
		}
	case "k", "up":
		if m.settings.cursor > 0 {
			m.settings.cursor--
		}
	case "left", "h":
		m.cycleSetting(-1)
	case "right", "l", "enter":
		m.cycleSetting(+1)
	}
	return m, nil
}

func (m *model) cycleSetting(dir int) {
	switch settingsRow(m.settings.cursor) {
	case rowLanguage:
		i := indexOf(m.settings.langs, m.cfg.Language)
		if i < 0 {
			i = 0
		}
		i = (i + dir + len(m.settings.langs)) % len(m.settings.langs)
		m.cfg.Language = m.settings.langs[i]
		// Apply immediately so the UI re-renders in the chosen language.
		if m.cfg.Language == "auto" {
			i18n.Set(i18n.Detect("auto", ""))
		} else {
			i18n.Set(i18n.Lang(m.cfg.Language))
		}
	case rowTheme:
		i := indexOf(m.settings.themes, m.cfg.Theme)
		if i < 0 {
			i = 0
		}
		i = (i + dir + len(m.settings.themes)) % len(m.settings.themes)
		m.cfg.Theme = m.settings.themes[i]
		m.theme = ThemeByName(m.cfg.Theme)
	case rowConfirmDelete:
		m.cfg.ConfirmDelete = !m.cfg.ConfirmDelete
	}
	if path, err := config.ConfigPath(); err == nil {
		_ = config.Save(path, m.cfg)
	}
}

func indexOf(slice []string, v string) int {
	for i, s := range slice {
		if s == v {
			return i
		}
	}
	return -1
}

func (m *model) settingsView() string {
	rows := []struct {
		label string
		value string
	}{
		{i18n.T("settings.language"), m.cfg.Language},
		{i18n.T("settings.theme"), m.cfg.Theme},
		{i18n.T("settings.confirm_delete"), boolLabel(m.cfg.ConfirmDelete)},
	}
	var b strings.Builder
	b.WriteString(m.theme.Title.Render(i18n.T("tui.settings.title")))
	b.WriteString("\n")
	b.WriteString(m.theme.Desc.Render(i18n.T("tui.settings.desc")))
	b.WriteString("\n\n")
	for i, r := range rows {
		label := m.theme.Subtitle.Render(r.label)
		line := label + "  " + m.theme.Status.Render("←") + " " +
			m.theme.Item.Render(r.value) + " " + m.theme.Status.Render("→")
		if i == m.settings.cursor {
			b.WriteString(m.theme.ItemSel.Render("▶ " + label + "  ← " + r.value + " →"))
		} else {
			b.WriteString("  " + line)
		}
		b.WriteString("\n\n")
	}
	return m.theme.BoxFocused.Width(m.innerWidth() - 2).Height(m.innerHeight()).Render(b.String())
}

func boolLabel(v bool) string {
	if v {
		return i18n.T("settings.on")
	}
	return i18n.T("settings.off")
}
