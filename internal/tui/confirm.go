package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/aaangelmartin/aka/internal/alias"
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
		case key.Matches(km, k.No):
			return c, confirmNo
		}
	}
	return c, confirmNone
}

func (c confirmModel) view() string {
	q := fmt.Sprintf("Delete %s?\n\n  %s\n",
		styleDanger.Render(c.target.Name),
		styleHint.Render(c.target.Command),
	)
	q += "\n" + styleHint.Render("y = yes   ·   n/esc = no")
	return styleFrame.Render(q)
}
