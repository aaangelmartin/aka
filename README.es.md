<div align="center">

```
         _
   __ _ | | __ __ _
  / _` || |/ // _` |
 | (_| ||   <| (_| |
  \__,_||_|\_\\__,_|
```

### Gestiona los aliases de tu shell — desde la terminal, con TUI.

[![CI](https://img.shields.io/github/actions/workflow/status/aaangelmartin/aka/ci.yml?branch=main&label=CI&style=for-the-badge&labelColor=000000&color=22D3A6)](https://github.com/aaangelmartin/aka/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/aaangelmartin/aka?sort=semver&style=for-the-badge&labelColor=000000&color=00B5E2)](https://github.com/aaangelmartin/aka/releases/latest)
[![Licencia: Apache 2.0](https://img.shields.io/badge/Licencia-Apache%202.0-D22128?style=for-the-badge&labelColor=000000)](./LICENSE)
[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=for-the-badge&labelColor=000000&logo=go&logoColor=white)](https://go.dev)
[![Hecho con Charm](https://img.shields.io/badge/hecho%20con-Charm-FF79C6?style=for-the-badge&labelColor=000000)](https://charm.sh)
[![Idiomas](https://img.shields.io/badge/lang-EN%20%C2%B7%20ES-00B5E2?style=for-the-badge&labelColor=000000)](./README.md)
[![Homebrew](https://img.shields.io/badge/brew-aaangelmartin%2Ftap%2Faka-F9C900?style=for-the-badge&labelColor=000000)](https://github.com/aaangelmartin/homebrew-tap)

[English](./README.md) · **Español**

</div>

---

## ¿Por qué `aka`?

Tu `.zshrc` (o `.bashrc`, o `config.fish`) es un cementerio de aliases
que ya nadie recuerda haber añadido, sin búsqueda, sin categorías, sin
vista previa. Editarlo a mano es frágil — una comilla mal puesta y tu
shell no carga.

`aka` (a.k.a. — *also known as*) resuelve eso:

- **Un único lugar para todos tus aliases**, respaldado por un archivo
  JSON.
- **TUI + CLI** para añadir, editar, borrar, buscar, etiquetar.
- **Multi-shell**: genera `aliases.zsh`, `aliases.bash` y
  `aliases.fish` desde la misma fuente.
- **Seguro por diseño**: tu rc se modifica una sola vez (una línea
  `source …` entre marcas `# >>> aka >>>`). `aka uninstall` limpia.
- **Bilingüe** (inglés · español).

## Instalación

### Homebrew (macOS / Linux)

```sh
brew install aaangelmartin/tap/aka
```

### Go

```sh
go install github.com/aaangelmartin/aka/cmd/aka@latest
```

### Manual

Descarga un binario desde la
[última release](https://github.com/aaangelmartin/aka/releases/latest),
extráelo y coloca `aka` en tu `$PATH`.

## Empezar rápido

```sh
aka install zsh                # añade `source ~/.config/aka/aliases.zsh` a ~/.zshrc
aka import --from-rc ~/.zshrc  # migra los aliases existentes
aka add cc claude              # añade el alias: cc='claude'
aka ls                         # lista aliases
aka                            # lanza la TUI
```

## Estado

En desarrollo temprano. Las funciones llegan en fases:

- **v0.1** — Núcleo CLI (`add`, `ls`, `rm`, `edit`, generadores shell).
- **v0.2** — TUI.
- **v0.3** — Install/uninstall en shell + importar desde rc.
- **v0.4** — i18n completo (EN / ES).
- **v0.5** — Fuzzy, etiquetas, temas, GIF demo.
- **v1.0** — Release pulida.

Mira [CHANGELOG.md](./CHANGELOG.md) para saber qué va en cada versión.

## Licencia

Apache 2.0 — ver [LICENSE](./LICENSE) y [NOTICE](./NOTICE).

## Contribuir

Ver [CONTRIBUTING.md](./CONTRIBUTING.md). La paridad de cadenas en
inglés y español se valida con un test; mantén ambas sincronizadas.
