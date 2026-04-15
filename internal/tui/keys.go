package tui

import "github.com/charmbracelet/bubbles/key"

// keyMap is the TUI-wide key binding set. Keys match GoTo's so muscle memory
// transfers between both tools.
type keyMap struct {
	Up       key.Binding
	Down     key.Binding
	Confirm  key.Binding // list: copy command (no "open" concept in aka)
	Add      key.Binding
	Edit     key.Binding
	Delete   key.Binding // d or x
	Filter   key.Binding
	Yank     key.Binding // y — same semantics as Confirm here, matches GoTo
	Tag      key.Binding // t — filter by selected row's first tag
	Help     key.Binding
	Quit     key.Binding
	Escape   key.Binding
	Lang     key.Binding // L — cycle EN ↔ ES, persist
	Settings key.Binding // o — open Settings screen
	Top      key.Binding // g
	Bottom   key.Binding // G
	Submit   key.Binding // ctrl+s on forms
	Next     key.Binding // tab
	Prev     key.Binding // shift+tab
}

func defaultKeys() keyMap {
	return keyMap{
		Up:       key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "up")),
		Down:     key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "down")),
		Confirm:  key.NewBinding(key.WithKeys("enter"), key.WithHelp("↵", "copy")),
		Add:      key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "add")),
		Edit:     key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "edit")),
		Delete:   key.NewBinding(key.WithKeys("d", "x"), key.WithHelp("d", "delete")),
		Filter:   key.NewBinding(key.WithKeys("/"), key.WithHelp("/", "filter")),
		Yank:     key.NewBinding(key.WithKeys("y"), key.WithHelp("y", "copy command")),
		Tag:      key.NewBinding(key.WithKeys("t"), key.WithHelp("t", "by tag")),
		Help:     key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
		Quit:     key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
		Escape:   key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "back")),
		Lang:     key.NewBinding(key.WithKeys("L"), key.WithHelp("L", "language")),
		Settings: key.NewBinding(key.WithKeys("o"), key.WithHelp("o", "settings")),
		Top:      key.NewBinding(key.WithKeys("g"), key.WithHelp("g", "top")),
		Bottom:   key.NewBinding(key.WithKeys("G"), key.WithHelp("G", "bottom")),
		Submit:   key.NewBinding(key.WithKeys("ctrl+s"), key.WithHelp("ctrl+s", "submit")),
		Next:     key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "next field")),
		Prev:     key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "prev field")),
	}
}
