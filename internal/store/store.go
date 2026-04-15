// Package store persists aliases to a JSON file atomically.
package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/aaangelmartin/aka/internal/alias"
)

// ErrNotFound is returned when a requested alias does not exist.
var ErrNotFound = errors.New("alias not found")

// ErrExists is returned when adding an alias that already exists.
var ErrExists = errors.New("alias already exists")

// Store manages a persistent collection of aliases.
type Store struct {
	Path    string
	aliases map[string]alias.Alias
}

// New returns a Store backed by path (not yet loaded).
func New(path string) *Store {
	return &Store{Path: path, aliases: map[string]alias.Alias{}}
}

// Load reads the file into memory. A missing file is treated as an empty store.
func (s *Store) Load() error {
	b, err := os.ReadFile(s.Path)
	if errors.Is(err, fs.ErrNotExist) {
		s.aliases = map[string]alias.Alias{}
		return nil
	}
	if err != nil {
		return err
	}
	if len(b) == 0 {
		s.aliases = map[string]alias.Alias{}
		return nil
	}
	var list []alias.Alias
	if err := json.Unmarshal(b, &list); err != nil {
		return fmt.Errorf("parse %s: %w", s.Path, err)
	}
	s.aliases = make(map[string]alias.Alias, len(list))
	for _, a := range list {
		s.aliases[key(a.Name)] = a
	}
	return nil
}

// Save writes all aliases atomically (temp file + rename).
func (s *Store) Save() error {
	if err := os.MkdirAll(filepath.Dir(s.Path), 0o755); err != nil {
		return err
	}
	list := s.List()
	b, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(s.Path), ".aka-*.tmp")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)
	if _, err := tmp.Write(b); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmpPath, s.Path)
}

// Get returns the alias by (case-sensitive) name.
func (s *Store) Get(name string) (alias.Alias, error) {
	a, ok := s.aliases[key(name)]
	if !ok {
		return alias.Alias{}, ErrNotFound
	}
	return a, nil
}

// Put inserts a new alias. Returns ErrExists if name is taken.
func (s *Store) Put(a alias.Alias) error {
	k := key(a.Name)
	if _, ok := s.aliases[k]; ok {
		return ErrExists
	}
	s.aliases[k] = a
	return nil
}

// Set inserts or replaces an alias (upsert).
func (s *Store) Set(a alias.Alias) {
	s.aliases[key(a.Name)] = a
}

// Delete removes an alias by name.
func (s *Store) Delete(name string) error {
	k := key(name)
	if _, ok := s.aliases[k]; !ok {
		return ErrNotFound
	}
	delete(s.aliases, k)
	return nil
}

// Rename moves an alias from old to new, preserving its fields.
func (s *Store) Rename(oldName, newName string) error {
	oldKey := key(oldName)
	newKey := key(newName)
	a, ok := s.aliases[oldKey]
	if !ok {
		return ErrNotFound
	}
	if oldKey == newKey {
		a.Name = newName
		s.aliases[oldKey] = a
		return nil
	}
	if _, exists := s.aliases[newKey]; exists {
		return ErrExists
	}
	delete(s.aliases, oldKey)
	a.Name = newName
	s.aliases[newKey] = a
	return nil
}

// List returns all aliases sorted by name (ASCII case-insensitive).
func (s *Store) List() []alias.Alias {
	out := make([]alias.Alias, 0, len(s.aliases))
	for _, a := range s.aliases {
		out = append(out, a)
	}
	sort.Slice(out, func(i, j int) bool {
		return strings.ToLower(out[i].Name) < strings.ToLower(out[j].Name)
	})
	return out
}

// Len returns the number of aliases.
func (s *Store) Len() int { return len(s.aliases) }

// key normalizes names for lookup. Shell alias names are case-sensitive in
// zsh/bash, but treating them that way here lets `foo` and `FOO` coexist;
// that usually causes more confusion than value, so we lower-case the key.
func key(name string) string { return strings.ToLower(name) }
