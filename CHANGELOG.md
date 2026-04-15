# Changelog / Registro de cambios

All notable changes to this project will be documented here. Format follows
[Keep a Changelog](https://keepachangelog.com/en/1.1.0/). Versions follow
[SemVer](https://semver.org/).

Todos los cambios relevantes se documentan aquí. El formato sigue
[Keep a Changelog](https://keepachangelog.com/es-ES/1.1.0/). Las versiones
siguen [SemVer](https://semver.org/lang/es/).

## [Unreleased]

## [0.5.0] — 2026-04-15

### Added / Añadido
- EN — Four TUI themes (`default`, `dracula`, `nord`, `gruvbox`). Selected with `aka config theme <name>`.
- ES — Cuatro temas para la TUI (`default`, `dracula`, `nord`, `gruvbox`). Se eligen con `aka config theme <nombre>`.
- EN — VHS demo tape (`demo/demo.tape`) so `make demo` regenerates the README GIF.
- ES — Script VHS (`demo/demo.tape`) para regenerar el GIF del README con `make demo`.

## [0.4.0] — 2026-04-15

### Added / Añadido
- EN — Bilingual strings (EN/ES) across CLI help and TUI labels, driven by `internal/i18n` with a `TestCatalogParity` guard.
- ES — Cadenas bilingües (EN/ES) en la ayuda CLI y etiquetas TUI, gestionadas en `internal/i18n` con el test `TestCatalogParity`.
- EN — Global `--lang {en,es}` flag; default honours `language` in config and falls back to `$LANG`.
- ES — Flag global `--lang {en,es}`; por defecto respeta `language` en config y usa `$LANG` como último recurso.
- EN — Confirm prompt accepts both `y` (English) and `s` (Spanish).
- ES — El prompt de confirmación acepta `y` (inglés) y `s` (español).

## [0.3.0] — 2026-04-15

### Added / Añadido
- EN — `aka install [zsh|bash|fish|all]`: inserts a managed block into the rc file (markers `# >>> aka >>>` / `# <<< aka <<<`), with a timestamped backup.
- ES — `aka install [zsh|bash|fish|all]`: inserta un bloque gestionado en el rc (marcas `# >>> aka >>>` / `# <<< aka <<<`), con copia de seguridad.
- EN — `aka uninstall`: removes the managed block cleanly; user content is preserved.
- ES — `aka uninstall`: elimina el bloque gestionado; el contenido del usuario se preserva.
- EN — `aka import --from-rc <file>`: parses `alias` lines from an existing rc (handles single/double quotes and `'\''` escape).
- ES — `aka import --from-rc <archivo>`: parsea líneas `alias` de un rc existente (soporta comillas y escape `'\''`).

## [0.2.0] — 2026-04-15

### Added / Añadido
- EN — Interactive TUI (Bubble Tea): list with filter, add/edit form, delete confirm, help overlay.
- ES — TUI interactiva (Bubble Tea): lista con filtro, formulario add/edit, confirmación de borrado, ayuda.
- EN — `Enter` copies the highlighted command to the clipboard.
- ES — `Enter` copia el comando destacado al portapapeles.
- EN — Running `aka` with no args launches the TUI.
- ES — Ejecutar `aka` sin argumentos lanza la TUI.

## [0.1.0] — 2026-04-15

### Added / Añadido
- EN — CLI core: `add`, `ls`, `rm`, `edit`, `config`, `import`, `export`, `version`.
- ES — Núcleo CLI: `add`, `ls`, `rm`, `edit`, `config`, `import`, `export`, `version`.
- EN — JSON store with atomic writes (temp + rename) as the source of truth.
- ES — Almacén JSON con escritura atómica (temp + rename) como fuente de verdad.
- EN — Shell emitters for zsh / bash / fish, regenerated on every mutation.
- ES — Generadores para zsh / bash / fish, regenerados en cada cambio.
- EN — XDG-aware paths (`$AKA_CONFIG`, `$AKA_DATA`, `$AKA_OUTDIR` to override).
- ES — Rutas XDG (`$AKA_CONFIG`, `$AKA_DATA`, `$AKA_OUTDIR` para sobreescribir).
- EN — Repository bootstrap: scaffolding, Makefile, goreleaser, CI/release workflows, Apache 2.0 license, bilingual READMEs.
- ES — Bootstrap del repositorio: estructura, Makefile, goreleaser, workflows CI/release, licencia Apache 2.0, READMEs bilingües.
