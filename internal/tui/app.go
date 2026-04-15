package tui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/aaangelmartin/aka/internal/alias"
	"github.com/aaangelmartin/aka/internal/config"
	"github.com/aaangelmartin/aka/internal/emit"
	"github.com/aaangelmartin/aka/internal/i18n"
	"github.com/aaangelmartin/aka/internal/store"
)

type screen int

const (
	screenList screen = iota
	screenForm
	screenConfirm
	screenHelp
	screenSettings
)

type formMode int

const (
	formAdd formMode = iota
	formEdit
)

// model is the TUI root model. It owns the store handle, the current theme,
// every screen's transient state, and the shared dimensions / status line.
type model struct {
	store  *store.Store
	cfg    config.Config
	outDir string
	theme  Theme
	keys   keyMap

	screen screen

	// list state
	items      []alias.Alias
	filter     string
	filterMode bool
	cursor     int
	offset     int
	tagFilter  string

	// form state
	form     formModel
	formKind formMode

	// confirm state
	confirmTarget alias.Alias
	confirmYes    bool

	// settings state
	settings settingsModel

	// dimensions
	width  int
	height int

	// status line (auto-fades after 3s)
	status    string
	statusExp time.Time
}

func newModel(st *store.Store, cfg config.Config, outDir string) *model {
	return &model{
		store:    st,
		cfg:      cfg,
		outDir:   outDir,
		theme:    ThemeByName(cfg.Theme),
		keys:     defaultKeys(),
		screen:   screenList,
		items:    st.List(),
		settings: newSettingsModel(cfg),
	}
}

func (m *model) Init() tea.Cmd { return nil }

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if w, ok := msg.(tea.WindowSizeMsg); ok {
		m.width, m.height = w.Width, w.Height
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
	case screenSettings:
		return m.updateSettings(msg)
	}
	return m, nil
}

func (m *model) View() string {
	header := m.headerView()
	var body string
	switch m.screen {
	case screenList:
		body = m.listView()
	case screenForm:
		body = m.formView()
	case screenConfirm:
		body = m.confirmView()
	case screenHelp:
		body = m.helpView()
	case screenSettings:
		body = m.settingsView()
	}
	footer := m.footerView()
	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

// ---------- shared helpers ----------

func (m *model) setStatus(s string) {
	m.status = s
	m.statusExp = time.Now().Add(3 * time.Second)
}

// commit saves the store and regenerates shell files. All mutating handlers
// call this; the return error is surfaced in the status line.
func (m *model) commit() error {
	if err := m.store.Save(); err != nil {
		return err
	}
	return emit.Regenerate(m.outDir, m.store.List())
}

func (m *model) refresh() {
	m.items = m.store.List()
	if f := m.filteredItems(); m.cursor >= len(f) {
		m.cursor = len(f) - 1
	}
	if m.cursor < 0 {
		m.cursor = 0
	}
}

func (m *model) filteredItems() []alias.Alias {
	q := strings.ToLower(strings.TrimSpace(m.filter))
	out := make([]alias.Alias, 0, len(m.items))
	for _, a := range m.items {
		if m.tagFilter != "" && !hasTag(a, m.tagFilter) {
			continue
		}
		if q == "" {
			out = append(out, a)
			continue
		}
		if strings.Contains(strings.ToLower(a.Name), q) ||
			strings.Contains(strings.ToLower(a.Command), q) ||
			strings.Contains(strings.ToLower(a.Description), q) {
			out = append(out, a)
			continue
		}
		for _, t := range a.Tags {
			if strings.Contains(strings.ToLower(t), q) {
				out = append(out, a)
				break
			}
		}
	}
	return out
}

func hasTag(a alias.Alias, tag string) bool {
	for _, t := range a.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

// ---------- dimensions ----------

func (m *model) innerWidth() int {
	if m.width == 0 {
		return 100
	}
	return m.width
}

func (m *model) innerHeight() int {
	if m.height == 0 {
		return 24
	}
	return max(6, m.height-3) // header + footer
}

func (m *model) leftWidth() int  { return max(28, m.innerWidth()*2/5) }
func (m *model) rightWidth() int { return max(28, m.innerWidth()-m.leftWidth()) }

// ---------- header / footer ----------

func (m *model) headerView() string {
	title := m.theme.Title.Render(" aka ")
	sub := m.theme.Subtitle.Render(i18n.Tf("tui.header.count", len(m.items)))
	extras := ""
	if m.filterMode || m.filter != "" {
		extras += m.theme.Status.Render(fmt.Sprintf("  /%s", m.filter))
	}
	if m.tagFilter != "" {
		extras += "  " + m.theme.Tag.Render(fmt.Sprintf("[#%s]", m.tagFilter))
	}
	return title + "  " + sub + extras
}

func (m *model) footerView() string {
	if m.status != "" && time.Now().Before(m.statusExp) {
		return m.theme.Status.Render(" " + m.status)
	}
	switch m.screen {
	case screenList:
		if m.filterMode {
			return m.theme.Help.Render(i18n.T("tui.footer.filter"))
		}
		return m.theme.Help.Render(i18n.T("tui.footer.list"))
	case screenForm:
		return m.theme.Help.Render(i18n.T("tui.footer.form"))
	case screenConfirm:
		return m.theme.Help.Render(i18n.T("tui.footer.confirm"))
	case screenHelp:
		return m.theme.Help.Render(i18n.T("tui.footer.back"))
	case screenSettings:
		return m.theme.Help.Render(i18n.T("tui.footer.settings"))
	}
	return ""
}

// ---------- help screen update (simple) ----------

func (m *model) updateHelp(msg tea.Msg) (tea.Model, tea.Cmd) {
	if k, ok := msg.(tea.KeyMsg); ok {
		switch k.String() {
		case "esc", "q", "?":
			m.screen = screenList
		}
	}
	return m, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
