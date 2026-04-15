<div align="center">

```
         _
   __ _ | | __ __ _
  / _` || |/ // _` |
 | (_| ||   <| (_| |
  \__,_||_|\_\\__,_|
```

### Manage your shell aliases — from the terminal, with a TUI.

[![CI](https://img.shields.io/github/actions/workflow/status/aaangelmartin/aka/ci.yml?branch=main&label=CI&style=for-the-badge&labelColor=000000&color=22D3A6)](https://github.com/aaangelmartin/aka/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/aaangelmartin/aka?sort=semver&style=for-the-badge&labelColor=000000&color=00B5E2)](https://github.com/aaangelmartin/aka/releases/latest)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-D22128?style=for-the-badge&labelColor=000000)](./LICENSE)
[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=for-the-badge&labelColor=000000&logo=go&logoColor=white)](https://go.dev)
[![Made with Charm](https://img.shields.io/badge/made%20with-Charm-FF79C6?style=for-the-badge&labelColor=000000)](https://charm.sh)
[![Languages](https://img.shields.io/badge/lang-EN%20%C2%B7%20ES-00B5E2?style=for-the-badge&labelColor=000000)](./README.es.md)
[![Homebrew](https://img.shields.io/badge/brew-aaangelmartin%2Ftap%2Faka-F9C900?style=for-the-badge&labelColor=000000)](https://github.com/aaangelmartin/homebrew-tap)

**English** · [Español](./README.es.md)

</div>

---

## Why `aka`?

Your `.zshrc` (or `.bashrc`, or `config.fish`) is a graveyard of aliases
nobody remembers adding, with no search, no categories, no preview.
Editing it by hand is fragile — one stray quote and your shell won't
load.

`aka` (a.k.a. — *also known as*) fixes that:

- **One place for all your aliases**, backed by a JSON file.
- **TUI + CLI** to add, edit, delete, search, tag.
- **Multi-shell**: generates `aliases.zsh`, `aliases.bash`, and
  `aliases.fish` from the same source.
- **Safe by design**: your rc file is modified exactly once (a single
  `source …` line between `# >>> aka >>>` markers). `aka uninstall`
  cleans up.
- **Bilingual** (English · Spanish).

## Install

### Homebrew (macOS / Linux)

```sh
brew install aaangelmartin/tap/aka
```

### Go

```sh
go install github.com/aaangelmartin/aka/cmd/aka@latest
```

### Manual

Grab a binary from the [latest release](https://github.com/aaangelmartin/aka/releases/latest),
extract, and drop `aka` on your `$PATH`.

## Quick start

```sh
aka install zsh                # add `source ~/.config/aka/aliases.zsh` to ~/.zshrc
aka import --from-rc ~/.zshrc  # migrate existing aliases
aka add cc claude              # add alias: cc='claude'
aka ls                         # list aliases
aka                            # launch the TUI
```

## Status

Early development. Features land in phases:

- **v0.1** — CLI core (`add`, `ls`, `rm`, `edit`, shell emitters).
- **v0.2** — TUI.
- **v0.3** — Shell install/uninstall + import from rc.
- **v0.4** — Full i18n (EN / ES).
- **v0.5** — Fuzzy, tags, themes, demo GIF.
- **v1.0** — Polished release.

See [CHANGELOG.md](./CHANGELOG.md) for what ships where.

## License

Apache 2.0 — see [LICENSE](./LICENSE) and [NOTICE](./NOTICE).

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md). Bilingual strings are enforced
by a parity test; please keep English and Spanish in sync.
