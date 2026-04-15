// Package tui implements the Bubble Tea interactive interface for aka.
package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/aaangelmartin/aka/internal/config"
	"github.com/aaangelmartin/aka/internal/emit"
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

	m := newAppModel(s, cfg, outDir)
	p := tea.NewProgram(m, tea.WithAltScreen())
	_, err = p.Run()
	return err
}

// commit persists the store and regenerates the shell files. Shared by TUI
// mutations.
func commitStore(s *store.Store, outDir string) error {
	if err := s.Save(); err != nil {
		return fmt.Errorf("save aliases: %w", err)
	}
	return emit.Regenerate(outDir, s.List())
}
