package tui

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/aaangelmartin/aka/internal/alias"
	"github.com/aaangelmartin/aka/internal/i18n"
)

const (
	fldName = iota
	fldCommand
	fldShells
	fldTags
	fldDesc
	fldCount
)

// formModel mirrors GoTo's: an array of textinputs plus a focused index,
// with inline error messages translated via i18n.
type formModel struct {
	inputs     [fldCount]textinput.Model
	focused    int
	errMsg     string // i18n key
	editing    string
	origCreate time.Time
	theme      Theme
}

func (m *model) openForm(kind formMode) {
	m.formKind = kind
	m.form = newFormModel(m.theme)
	m.screen = screenForm
}

func newFormModel(th Theme) formModel {
	mk := func(placeholder string, width int) textinput.Model {
		ti := textinput.New()
		ti.Placeholder = placeholder
		ti.Width = width
		ti.CharLimit = 512
		ti.Prompt = "› "
		ti.PromptStyle = lipgloss.NewStyle().Foreground(th.Accent)
		ti.TextStyle = lipgloss.NewStyle().Foreground(th.FG)
		return ti
	}
	f := formModel{theme: th}
	f.inputs[fldName] = mk(i18n.T("tui.placeholder.name"), 40)
	f.inputs[fldCommand] = mk(i18n.T("tui.placeholder.command"), 60)
	f.inputs[fldShells] = mk(i18n.T("tui.placeholder.shells"), 40)
	f.inputs[fldTags] = mk(i18n.T("tui.placeholder.tags"), 40)
	f.inputs[fldDesc] = mk(i18n.T("tui.placeholder.desc"), 60)
	return f
}

func (f *formModel) focusFirst() tea.Cmd {
	f.focused = 0
	f.inputs[0].Focus()
	return textinput.Blink
}

func (f *formModel) loadFrom(a alias.Alias) {
	f.editing = a.Name
	f.origCreate = a.CreatedAt
	f.inputs[fldName].SetValue(a.Name)
	f.inputs[fldCommand].SetValue(a.Command)
	f.inputs[fldShells].SetValue(strings.Join(a.Shells, ", "))
	f.inputs[fldTags].SetValue(strings.Join(a.Tags, ", "))
	f.inputs[fldDesc].SetValue(a.Description)
}

func (f *formModel) nextField() {
	f.inputs[f.focused].Blur()
	f.focused = (f.focused + 1) % fldCount
	f.inputs[f.focused].Focus()
}

func (f *formModel) prevField() {
	f.inputs[f.focused].Blur()
	f.focused = (f.focused - 1 + fldCount) % fldCount
	f.inputs[f.focused].Focus()
}

func (m *model) updateForm(msg tea.Msg) (tea.Model, tea.Cmd) {
	if km, ok := msg.(tea.KeyMsg); ok {
		switch km.String() {
		case "esc":
			m.screen = screenList
			return m, nil
		case "tab", "down":
			m.form.nextField()
			return m, nil
		case "shift+tab", "up":
			m.form.prevField()
			return m, nil
		case "ctrl+s":
			return m, m.submitForm()
		case "enter":
			if m.form.focused == fldCount-1 {
				return m, m.submitForm()
			}
			m.form.nextField()
			return m, nil
		}
	}
	var cmd tea.Cmd
	m.form.inputs[m.form.focused], cmd = m.form.inputs[m.form.focused].Update(msg)
	return m, cmd
}

func (m *model) submitForm() tea.Cmd {
	a, errKey := m.form.build()
	if errKey != "" {
		m.form.errMsg = errKey
		return nil
	}
	if err := alias.Validate(a); err != nil {
		m.form.errMsg = "err.validation"
		m.setStatus(err.Error())
		return nil
	}
	if m.formKind == formEdit {
		if a.Name != m.form.editing {
			if err := m.store.Rename(m.form.editing, a.Name); err != nil {
				m.form.errMsg = "err.rename"
				m.setStatus(err.Error())
				return nil
			}
		}
		m.store.Set(a)
	} else {
		if err := m.store.Put(a); err != nil {
			m.form.errMsg = "err.exists"
			m.setStatus(err.Error())
			return nil
		}
	}
	if err := m.commit(); err != nil {
		m.setStatus(err.Error())
		return nil
	}
	m.refresh()
	m.setStatus(i18n.Tf("tui.status.saved", a.Name))
	m.screen = screenList
	return nil
}

// build returns the Alias from the current form values, or an i18n key when a
// field is empty. Validation (charset, shell names) is deferred to
// alias.Validate in the caller.
func (f *formModel) build() (alias.Alias, string) {
	name := strings.TrimSpace(f.inputs[fldName].Value())
	cmd := f.inputs[fldCommand].Value()
	if name == "" {
		return alias.Alias{}, "err.empty_name"
	}
	if strings.TrimSpace(cmd) == "" {
		return alias.Alias{}, "err.empty_command"
	}
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
	for _, t := range strings.Split(f.inputs[fldTags].Value(), ",") {
		t = strings.TrimSpace(t)
		if t != "" {
			tags = append(tags, t)
		}
	}
	created := f.origCreate
	if created.IsZero() {
		created = time.Now().UTC()
	}
	return alias.Alias{
		Name:        name,
		Command:     cmd,
		Shells:      shells,
		Tags:        tags,
		Description: strings.TrimSpace(f.inputs[fldDesc].Value()),
		CreatedAt:   created,
	}, ""
}

func (m *model) formView() string {
	labels := [fldCount]string{
		i18n.T("tui.field.name"),
		i18n.T("tui.field.command"),
		i18n.T("tui.field.shells"),
		i18n.T("tui.field.tags"),
		i18n.T("tui.field.desc"),
	}
	var b strings.Builder
	title := i18n.T("tui.form.add")
	if m.formKind == formEdit {
		title = i18n.T("tui.form.edit")
	}
	b.WriteString(m.theme.Title.Render(title))
	b.WriteString("\n\n")
	for i, l := range labels {
		var label string
		if i == m.form.focused {
			label = m.theme.Key.Render("› " + l + ":")
		} else {
			label = "  " + m.theme.Subtitle.Render(l+":")
		}
		b.WriteString(label)
		b.WriteString("\n  ")
		b.WriteString(m.form.inputs[i].View())
		b.WriteString("\n\n")
	}
	if m.form.errMsg != "" {
		b.WriteString(m.theme.Danger_.Render("✗ " + i18n.T(m.form.errMsg)))
	}
	return m.theme.BoxFocused.Width(m.innerWidth() - 2).Height(m.innerHeight()).Render(b.String())
}
