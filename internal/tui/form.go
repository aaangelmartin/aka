package tui

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/aaangelmartin/aka/internal/alias"
)

type formMode int

const (
	modeAdd formMode = iota
	modeEdit
)

const (
	fldName = iota
	fldCommand
	fldShells
	fldTags
	fldDesc
	fldCount
)

type formModel struct {
	mode       formMode
	inputs     [fldCount]textinput.Model
	focused    int
	origName   string
	origCreate time.Time
	err        string
}

func newForm() formModel {
	f := formModel{}
	labels := [fldCount]string{
		"name",
		"command",
		"shells (zsh,bash,fish — blank = all)",
		"tags (comma-separated)",
		"description",
	}
	for i := 0; i < fldCount; i++ {
		ti := textinput.New()
		ti.Prompt = labels[i] + ": "
		ti.CharLimit = 512
		ti.Width = 60
		f.inputs[i] = ti
	}
	f.focus(fldName)
	return f
}

func (f *formModel) loadFromAlias(a alias.Alias) {
	f.mode = modeEdit
	f.origName = a.Name
	f.origCreate = a.CreatedAt
	f.inputs[fldName].SetValue(a.Name)
	f.inputs[fldCommand].SetValue(a.Command)
	f.inputs[fldShells].SetValue(strings.Join(a.Shells, ","))
	f.inputs[fldTags].SetValue(strings.Join(a.Tags, ","))
	f.inputs[fldDesc].SetValue(a.Description)
	f.focus(fldName)
}

func (f *formModel) reset(mode formMode) {
	next := newForm()
	next.mode = mode
	*f = next
}

func (f *formModel) focus(i int) {
	for k := range f.inputs {
		if k == i {
			f.inputs[k].Focus()
		} else {
			f.inputs[k].Blur()
		}
	}
	f.focused = i
}

func (f formModel) buildAlias() (alias.Alias, error) {
	name := strings.TrimSpace(f.inputs[fldName].Value())
	cmd := f.inputs[fldCommand].Value()
	var shells []string
	if raw := strings.TrimSpace(f.inputs[fldShells].Value()); raw != "" {
		for _, s := range strings.Split(raw, ",") {
			s = strings.TrimSpace(s)
			if s != "" {
				shells = append(shells, s)
			}
		}
	}
	var tags []string
	if raw := strings.TrimSpace(f.inputs[fldTags].Value()); raw != "" {
		for _, t := range strings.Split(raw, ",") {
			t = strings.TrimSpace(t)
			if t != "" {
				tags = append(tags, t)
			}
		}
	}
	created := f.origCreate
	if created.IsZero() {
		created = time.Now().UTC()
	}
	a := alias.Alias{
		Name:        name,
		Command:     cmd,
		Shells:      shells,
		Tags:        tags,
		Description: strings.TrimSpace(f.inputs[fldDesc].Value()),
		CreatedAt:   created,
	}
	if err := alias.Validate(a); err != nil {
		return a, err
	}
	return a, nil
}

type formAction int

const (
	formNone formAction = iota
	formSubmit
	formCancel
)

func (f formModel) update(msg tea.Msg, k keyMap) (formModel, tea.Cmd, formAction) {
	if km, ok := msg.(tea.KeyMsg); ok {
		switch {
		case key.Matches(km, k.Cancel):
			return f, nil, formCancel
		case key.Matches(km, k.NextFld):
			f.focus((f.focused + 1) % fldCount)
			return f, nil, formNone
		case key.Matches(km, k.PrevFld):
			f.focus((f.focused - 1 + fldCount) % fldCount)
			return f, nil, formNone
		case key.Matches(km, k.Save):
			return f, nil, formSubmit
		case km.Type == tea.KeyEnter:
			if f.focused < fldCount-1 {
				f.focus(f.focused + 1)
				return f, nil, formNone
			}
			return f, nil, formSubmit
		}
	}
	var cmd tea.Cmd
	f.inputs[f.focused], cmd = f.inputs[f.focused].Update(msg)
	return f, cmd, formNone
}

func (f formModel) view() string {
	var b strings.Builder
	title := "Add alias"
	if f.mode == modeEdit {
		title = "Edit alias"
	}
	b.WriteString(styleTitle.Render(title))
	b.WriteString("\n\n")
	for i, in := range f.inputs {
		if i == f.focused {
			b.WriteString(styleInput.Render("▸ "))
		} else {
			b.WriteString("  ")
		}
		b.WriteString(in.View())
		b.WriteString("\n")
	}
	b.WriteString("\n")
	if f.err != "" {
		b.WriteString(styleDanger.Render("✗ " + f.err))
		b.WriteString("\n\n")
	}
	b.WriteString(styleHint.Render("tab/shift+tab move · enter next/submit · ctrl+s submit · esc cancel"))
	return styleFrame.Render(b.String())
}
