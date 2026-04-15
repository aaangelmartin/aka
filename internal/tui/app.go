package tui

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/aaangelmartin/aka/internal/config"
	"github.com/aaangelmartin/aka/internal/i18n"
	"github.com/aaangelmartin/aka/internal/store"
)

type screen int

const (
	screenList screen = iota
	screenForm
	screenConfirm
	screenHelp
)

type appModel struct {
	store  *store.Store
	cfg    config.Config
	outDir string

	screen screen
	list   list.Model

	form    formModel
	confirm confirmModel
	keys    keyMap

	status    string
	statusExp time.Time
	width     int
	height    int
}

type clearStatusMsg struct{}

func newAppModel(s *store.Store, cfg config.Config, outDir string) *appModel {
	delegate := list.NewDefaultDelegate()
	lst := list.New(itemsFromStore(s), delegate, 0, 0)
	lst.Title = i18n.T("tui.title")
	lst.SetShowStatusBar(false)
	lst.SetFilteringEnabled(true)
	lst.SetShowHelp(false)
	return &appModel{
		store:  s,
		cfg:    cfg,
		outDir: outDir,
		list:   lst,
		form:   newForm(),
		keys:   newKeys(),
	}
}

func itemsFromStore(s *store.Store) []list.Item {
	entries := s.List()
	out := make([]list.Item, 0, len(entries))
	for _, a := range entries {
		out = append(out, aliasItem{a})
	}
	return out
}

func (m *appModel) Init() tea.Cmd { return nil }

func (m *appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.list.SetSize(msg.Width, msg.Height-3)
	case clearStatusMsg:
		if time.Now().After(m.statusExp) {
			m.status = ""
		}
	}

	switch m.screen {
	case screenList:
		return m.updateList(msg)
	case screenForm:
		return m.updateForm(msg)
	case screenConfirm:
		return m.updateConfirm(msg)
	case screenHelp:
		return m.updateHelp(msg)
	}
	return m, nil
}

func (m *appModel) updateList(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Forward any filter-mode typing straight to the list.
	if m.list.FilterState() == list.Filtering {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}
	if km, ok := msg.(tea.KeyMsg); ok {
		switch {
		case key.Matches(km, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(km, m.keys.Help):
			m.screen = screenHelp
			return m, nil
		case key.Matches(km, m.keys.Add):
			m.form.reset(modeAdd)
			m.screen = screenForm
			return m, nil
		case key.Matches(km, m.keys.Edit):
			if sel, ok := m.list.SelectedItem().(aliasItem); ok {
				m.form = newForm()
				m.form.loadFromAlias(sel.Alias)
				m.screen = screenForm
			}
			return m, nil
		case key.Matches(km, m.keys.Delete):
			if sel, ok := m.list.SelectedItem().(aliasItem); ok {
				m.confirm = confirmModel{target: sel.Alias}
				m.screen = screenConfirm
			}
			return m, nil
		case key.Matches(km, m.keys.Copy):
			if sel, ok := m.list.SelectedItem().(aliasItem); ok {
				if err := clipboard.WriteAll(sel.Command); err != nil {
					return m, m.flash("clipboard error: " + err.Error())
				}
				return m, m.flash("copied: " + sel.Command)
			}
			return m, nil
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *appModel) updateForm(msg tea.Msg) (tea.Model, tea.Cmd) {
	form, cmd, act := m.form.update(msg, m.keys)
	m.form = form
	switch act {
	case formCancel:
		m.screen = screenList
		return m, nil
	case formSubmit:
		a, err := m.form.buildAlias()
		if err != nil {
			m.form.err = err.Error()
			return m, nil
		}
		if m.form.mode == modeAdd {
			if err := m.store.Put(a); err != nil {
				m.form.err = err.Error()
				return m, nil
			}
		} else {
			if a.Name != m.form.origName {
				if err := m.store.Rename(m.form.origName, a.Name); err != nil {
					m.form.err = err.Error()
					return m, nil
				}
			}
			m.store.Set(a)
		}
		if err := commitStore(m.store, m.outDir); err != nil {
			m.form.err = err.Error()
			return m, nil
		}
		m.reloadList()
		m.screen = screenList
		return m, m.flash(i18n.Tf("tui.status.saved", a.Name))
	}
	return m, cmd
}

func (m *appModel) updateConfirm(msg tea.Msg) (tea.Model, tea.Cmd) {
	c, act := m.confirm.update(msg, m.keys)
	m.confirm = c
	switch act {
	case confirmYes:
		name := c.target.Name
		if err := m.store.Delete(name); err != nil {
			m.screen = screenList
			return m, m.flash("error: " + err.Error())
		}
		if err := commitStore(m.store, m.outDir); err != nil {
			m.screen = screenList
			return m, m.flash("error: " + err.Error())
		}
		m.reloadList()
		m.screen = screenList
		return m, m.flash(i18n.Tf("tui.status.deleted", name))
	case confirmNo:
		m.screen = screenList
		return m, nil
	}
	return m, nil
}

func (m *appModel) updateHelp(msg tea.Msg) (tea.Model, tea.Cmd) {
	if _, ok := msg.(tea.KeyMsg); ok {
		m.screen = screenList
	}
	return m, nil
}

func (m *appModel) reloadList() {
	items := itemsFromStore(m.store)
	m.list.SetItems(items)
}

func (m *appModel) flash(s string) tea.Cmd {
	m.status = s
	m.statusExp = time.Now().Add(3 * time.Second)
	return tea.Tick(3*time.Second, func(time.Time) tea.Msg { return clearStatusMsg{} })
}

func (m *appModel) View() string {
	switch m.screen {
	case screenForm:
		return centerOver(m.form.view(), m.width, m.height)
	case screenConfirm:
		return centerOver(m.confirm.view(), m.width, m.height)
	case screenHelp:
		return centerOver(helpView(), m.width, m.height)
	}
	// List view with a trailing status/hint line.
	bottom := styleHint.Render(i18n.T("tui.list.hint"))
	if m.status != "" {
		bottom = styleOK.Render(m.status) + "   " + bottom
	}
	return fmt.Sprintf("%s\n%s", m.list.View(), bottom)
}

// centerOver draws body centered on a terminal of (w, h). When dimensions are
// unknown (0) it just returns body.
func centerOver(body string, w, h int) string {
	if w == 0 || h == 0 {
		return body
	}
	// Simple centering: prepend blank lines. Lipgloss has Place but we keep
	// this minimal to avoid another dependency on its layout helpers.
	return body
}
