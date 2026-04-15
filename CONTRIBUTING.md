# Contributing to aka · Contribuir a aka

**English** · [Español](#español)

Thanks for your interest in `aka`! This project manages shell aliases
through a CLI + TUI, with a strong bias for reversibility and safety.

## Development setup

Requirements: Go 1.22+, `make`, `golangci-lint` (optional), `goreleaser`
(optional for local snapshot builds), `vhs` (optional for demo gif).

```sh
git clone https://github.com/aaangelmartin/aka.git
cd aka
make build         # builds ./bin/aka
./bin/aka version
```

Run tests / lint:

```sh
make test          # go test -race -cover ./...
make lint          # golangci-lint run
```

## Workflow

- Conventional Commits: `feat:`, `fix:`, `docs:`, `chore:`, `refactor:`,
  `test:`, `ci:`, `build:`. Breaking changes use `!`: `feat!: drop X`.
- Work on a feature branch (`feat/…`, `fix/…`), open a PR against `main`.
- Never push directly to `main`.
- Commit messages in English.

## Bilingual rule

Every user-facing string must exist in **both English and Spanish**. CLI
and TUI strings live in `internal/i18n/catalog.go` with an `EN` and `ES`
map. A parity test guards this — if you add a key to one, add it to the
other.

README has an English (`README.md`) and a Spanish (`README.es.md`)
version; keep them in sync.

## Before opening a PR

1. `make fmt`
2. `make test`
3. `make lint`
4. Update `CHANGELOG.md` (Unreleased section, EN + ES bullet).
5. If you added a user-facing string, update both i18n catalogs.

## Releases

Releases are tag-driven. Maintainers tag `vX.Y.Z` on `main`, which
triggers goreleaser to publish binaries to GitHub Releases and update
the Homebrew tap. Contributors don't need to do anything special.

---

## Español

¡Gracias por tu interés en `aka`! Es un gestor de aliases de shell con
CLI + TUI, diseñado para ser seguro y reversible.

### Puesta en marcha

Requisitos: Go 1.22+, `make`, `golangci-lint` (opcional), `goreleaser`
(opcional para builds locales), `vhs` (opcional para el gif).

```sh
git clone https://github.com/aaangelmartin/aka.git
cd aka
make build
./bin/aka version
```

Tests / lint:

```sh
make test
make lint
```

### Flujo de trabajo

- Conventional Commits: `feat:`, `fix:`, `docs:`, `chore:`, `refactor:`,
  `test:`, `ci:`, `build:`. Cambios incompatibles con `!`:
  `feat!: eliminar X`.
- Trabaja en una rama (`feat/…`, `fix/…`) y abre un PR a `main`.
- Nunca hagas push directo a `main`.
- Mensajes de commit en inglés.

### Regla bilingüe

Toda cadena visible para el usuario tiene que estar en **inglés y
español**. Las cadenas del CLI y la TUI viven en
`internal/i18n/catalog.go` con mapas `EN` y `ES`. Un test de paridad
obliga a mantener ambas listas alineadas.

El README tiene versión inglesa (`README.md`) y española
(`README.es.md`); mantén ambas sincronizadas.

### Antes de abrir un PR

1. `make fmt`
2. `make test`
3. `make lint`
4. Actualiza `CHANGELOG.md` (sección Unreleased, bullet EN + ES).
5. Si añadiste una cadena visible, actualiza los dos catálogos i18n.

### Releases

Las releases se disparan con tags. Los mantenedores etiquetan
`vX.Y.Z` sobre `main`; goreleaser publica los binarios en GitHub
Releases y actualiza el tap de Homebrew. Quien contribuye no tiene
que hacer nada especial.
