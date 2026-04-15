package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/aaangelmartin/aka/internal/alias"
	"github.com/aaangelmartin/aka/internal/i18n"
)

type confirmModel struct {
	target alias.Alias
}

type confirmAction int

const (
	confirmNone confirmAction = iota
	confirmYes
	confirmNo
)

func (c confirmModel) update(msg tea.Msg, k keyMap) (confirmModel, confirmAction) {
	if km, ok := msg.(tea.KeyMsg); ok {
		switch {
		case key.Matches(km, k.Yes):
			return c, confirmYes
		case km.String() == "s": // Spanish "sí"
			return c, confirmYes
		case key.Matches(km, k.No):
			return c, confirmNo
		}
	}
	return c, confirmNone
}

func (c confirmModel) view() string {
	q := i18n.Tf("tui.confirm.title", styleDanger.Render(c.target.Name)) +
		"\n\n  " + styleHint.Render(c.target.Command) + "\n"
	q += "\n" + styleHint.Render(i18n.T("tui.confirm.hint"))
	return styleFrame.Render(q)
}
