package tui

import (
	"strings"

	"github.com/aaangelmartin/aka/internal/alias"
)

// aliasItem adapts alias.Alias to the bubbles/list.Item interface.
type aliasItem struct{ alias.Alias }

// Title is the left column shown for each row.
func (i aliasItem) Title() string { return i.Name }

// Description is the right column shown for each row.
func (i aliasItem) Description() string {
	if len(i.Tags) == 0 {
		return i.Command
	}
	return i.Command + "  · " + strings.Join(i.Tags, ",")
}

// FilterValue is the text matched against when the user types in the list's
// filter prompt.
func (i aliasItem) FilterValue() string {
	return i.Name + " " + i.Command + " " + strings.Join(i.Tags, " ")
}
