package cli

import (
	"fmt"

	"github.com/aaangelmartin/aka/internal/config"
	"github.com/aaangelmartin/aka/internal/emit"
	"github.com/aaangelmartin/aka/internal/store"
)

// session bundles the loaded store + config + resolved paths for a single
// command invocation. All CLI subcommands go through openSession.
type session struct {
	store  *store.Store
	cfg    config.Config
	outDir string
}

func openSession() (*session, error) {
	cfgPath, err := config.ConfigPath()
	if err != nil {
		return nil, fmt.Errorf("resolve config path: %w", err)
	}
	cfg, err := config.Load(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}
	aliasPath, err := config.AliasesPath()
	if err != nil {
		return nil, fmt.Errorf("resolve aliases path: %w", err)
	}
	s := store.New(aliasPath)
	if err := s.Load(); err != nil {
		return nil, fmt.Errorf("load aliases: %w", err)
	}
	outDir, err := config.OutDir()
	if err != nil {
		return nil, fmt.Errorf("resolve out dir: %w", err)
	}
	return &session{store: s, cfg: cfg, outDir: outDir}, nil
}

// commit persists the store and regenerates the shell alias files. Call it
// after any mutation (add/rm/edit/import).
func (s *session) commit() error {
	if err := s.store.Save(); err != nil {
		return fmt.Errorf("save aliases: %w", err)
	}
	if err := emit.Regenerate(s.outDir, s.store.List()); err != nil {
		return fmt.Errorf("regenerate shell files: %w", err)
	}
	return nil
}
