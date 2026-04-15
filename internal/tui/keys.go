package tui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Up      key.Binding
	Down    key.Binding
	Add     key.Binding
	Edit    key.Binding
	Delete  key.Binding
	Copy    key.Binding
	Filter  key.Binding
	Help    key.Binding
	Quit    key.Binding
	Save    key.Binding
	Cancel  key.Binding
	NextFld key.Binding
	PrevFld key.Binding
	Yes     key.Binding
	No      key.Binding
}

func newKeys() keyMap {
	return keyMap{
		Up:      key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "up")),
		Down:    key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "down")),
		Add:     key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "add")),
		Edit:    key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "edit")),
		Delete:  key.NewBinding(key.WithKeys("d"), key.WithHelp("d", "delete")),
		Copy:    key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "copy command")),
		Filter:  key.NewBinding(key.WithKeys("/"), key.WithHelp("/", "filter")),
		Help:    key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
		Quit:    key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
		Save:    key.NewBinding(key.WithKeys("ctrl+s"), key.WithHelp("ctrl+s", "save")),
		Cancel:  key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "cancel")),
		NextFld: key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next field")),
		PrevFld: key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "prev field")),
		Yes:     key.NewBinding(key.WithKeys("y"), key.WithHelp("y", "yes")),
		No:      key.NewBinding(key.WithKeys("n", "esc"), key.WithHelp("n/esc", "no")),
	}
}
