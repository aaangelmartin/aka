// Package tui implements the Bubble Tea interactive interface for aka.
package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/aaangelmartin/aka/internal/config"
	"github.com/aaangelmartin/aka/internal/store"
)

// Run boots the TUI. It loads the store + config, then enters the Bubble Tea
// event loop.
func Run() error {
	cfgPath, err := config.ConfigPath()
	if err != nil {
		return err
	}
	cfg, err := config.Load(cfgPath)
	if err != nil {
		return err
	}
	aliasPath, err := config.AliasesPath()
	if err != nil {
		return err
	}
	s := store.New(aliasPath)
	if err := s.Load(); err != nil {
		return err
	}
	outDir, err := config.OutDir()
	if err != nil {
		return err
	}

	m := newModel(s, cfg, outDir)
	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())
	_, err = p.Run()
	return err
}
