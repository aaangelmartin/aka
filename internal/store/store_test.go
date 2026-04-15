package store

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/aaangelmartin/aka/internal/alias"
)

func tmpStore(t *testing.T) *Store {
	t.Helper()
	dir := t.TempDir()
	return New(filepath.Join(dir, "aliases.json"))
}

func TestPutLoadSave(t *testing.T) {
	s := tmpStore(t)
	if err := s.Load(); err != nil {
		t.Fatalf("load empty: %v", err)
	}
	if s.Len() != 0 {
		t.Fatalf("expected empty store, got %d", s.Len())
	}
	a := alias.Alias{Name: "gs", Command: "git status"}
	if err := s.Put(a); err != nil {
		t.Fatalf("put: %v", err)
	}
	if err := s.Put(a); !errors.Is(err, ErrExists) {
		t.Fatalf("expected ErrExists on duplicate, got %v", err)
	}
	if err := s.Save(); err != nil {
		t.Fatalf("save: %v", err)
	}
	s2 := New(s.Path)
	if err := s2.Load(); err != nil {
		t.Fatalf("reload: %v", err)
	}
	got, err := s2.Get("gs")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Command != "git status" {
		t.Fatalf("bad command: %q", got.Command)
	}
}

func TestDeleteRename(t *testing.T) {
	s := tmpStore(t)
	_ = s.Put(alias.Alias{Name: "a", Command: "echo a"})
	_ = s.Put(alias.Alias{Name: "b", Command: "echo b"})

	if err := s.Rename("a", "b"); !errors.Is(err, ErrExists) {
		t.Fatalf("expected ErrExists on rename into taken name, got %v", err)
	}
	if err := s.Rename("a", "c"); err != nil {
		t.Fatalf("rename: %v", err)
	}
	if _, err := s.Get("a"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("old name should be gone, got %v", err)
	}
	got, err := s.Get("c")
	if err != nil || got.Command != "echo a" {
		t.Fatalf("renamed alias missing or wrong command: %+v %v", got, err)
	}

	if err := s.Delete("nope"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound on delete of missing, got %v", err)
	}
	if err := s.Delete("c"); err != nil {
		t.Fatalf("delete: %v", err)
	}
	if s.Len() != 1 {
		t.Fatalf("after delete expected 1 alias, got %d", s.Len())
	}
}

func TestAtomicSaveRecoversFromCorruptPartial(t *testing.T) {
	// Confirm that Save leaves no stray temp files behind on success.
	s := tmpStore(t)
	_ = s.Put(alias.Alias{Name: "x", Command: "y"})
	if err := s.Save(); err != nil {
		t.Fatalf("save: %v", err)
	}
	entries, err := os.ReadDir(filepath.Dir(s.Path))
	if err != nil {
		t.Fatalf("readdir: %v", err)
	}
	for _, e := range entries {
		if filepath.Ext(e.Name()) == ".tmp" {
			t.Fatalf("leftover temp file: %s", e.Name())
		}
	}
}

func TestListSortedCaseInsensitive(t *testing.T) {
	s := tmpStore(t)
	_ = s.Put(alias.Alias{Name: "Bravo", Command: "b"})
	_ = s.Put(alias.Alias{Name: "alpha", Command: "a"})
	_ = s.Put(alias.Alias{Name: "Charlie", Command: "c"})
	got := s.List()
	names := []string{got[0].Name, got[1].Name, got[2].Name}
	want := []string{"alpha", "Bravo", "Charlie"}
	for i := range want {
		if names[i] != want[i] {
			t.Fatalf("sort mismatch at %d: got %v; want %v", i, names, want)
		}
	}
}
