package i18n

// en is the English catalog. Every key in en MUST also appear in es (and
// vice versa); the parity test enforces this.
var en = map[string]string{
	// CLI — short descriptions
	"cli.root.short":      "Manage shell aliases from the terminal with a TUI",
	"cli.root.long":       "aka — manage shell aliases from the terminal with a TUI. Multi-shell (zsh, bash, fish).",
	"cli.add.short":       "Add a new alias",
	"cli.ls.short":        "List aliases",
	"cli.rm.short":        "Remove an alias",
	"cli.edit.short":      "Edit an alias in place",
	"cli.config.short":    "Read or write a config key",
	"cli.import.short":    "Import aliases from JSON (or an rc file with --from-rc)",
	"cli.export.short":    "Export aliases as JSON (stdout if no file given)",
	"cli.install.short":   "Add the aka source line to your shell rc file(s)",
	"cli.uninstall.short": "Remove the aka source block from your shell rc file(s)",
	"cli.version.short":   "Print version, commit, and build date",

	// CLI — messages
	"msg.no_aliases":     "no aliases yet — try `aka add <name> <command>`",
	"msg.added":          "added %s → %s",
	"msg.removed":        "removed %s",
	"msg.updated":        "updated %s",
	"msg.aborted":        "aborted",
	"msg.no_changes":     "no changes",
	"msg.confirm_delete": "delete %s (%s)? [y/N] ",
	"msg.install_hint":   "open a new terminal or `source` the rc file to pick up your aliases.",
	"msg.import_summary": "imported %d new, updated %d, skipped %d",

	// TUI
	"tui.title":                "aka — aliases",
	"tui.form.add":             "Add alias",
	"tui.form.edit":            "Edit alias",
	"tui.form.field.name":      "name",
	"tui.form.field.command":   "command",
	"tui.form.field.shells":    "shells (zsh,bash,fish — blank = all)",
	"tui.form.field.tags":      "tags (comma-separated)",
	"tui.form.field.desc":      "description",
	"tui.form.hint":            "tab/shift+tab move · enter next/submit · ctrl+s submit · esc cancel",
	"tui.list.hint":            "a add · e edit · d delete · enter copy · / filter · ? help · q quit",
	"tui.confirm.title":        "Delete %s?",
	"tui.confirm.hint":         "y = yes   ·   n/esc = no",
	"tui.help.title":           "aka — keys",
	"tui.help.return":          "press any key to return",
	"tui.status.saved":         "saved: %s",
	"tui.status.deleted":       "deleted: %s",
	"tui.status.copied":        "copied: %s",
	"tui.status.clipboard_err": "clipboard error: %s",
}

// es is the Spanish catalog. Keep every key in sync with en.
var es = map[string]string{
	// CLI — short descriptions
	"cli.root.short":      "Gestiona los aliases de tu shell desde la terminal con TUI",
	"cli.root.long":       "aka — gestiona los aliases de tu shell desde la terminal con TUI. Multi-shell (zsh, bash, fish).",
	"cli.add.short":       "Añade un alias",
	"cli.ls.short":        "Lista los aliases",
	"cli.rm.short":        "Elimina un alias",
	"cli.edit.short":      "Edita un alias existente",
	"cli.config.short":    "Lee o escribe una clave de configuración",
	"cli.import.short":    "Importa aliases desde JSON (o desde un rc con --from-rc)",
	"cli.export.short":    "Exporta aliases en JSON (a stdout si no se indica archivo)",
	"cli.install.short":   "Añade la línea `source` de aka a tu(s) rc de shell",
	"cli.uninstall.short": "Elimina el bloque de aka del rc de tu(s) shell",
	"cli.version.short":   "Muestra versión, commit y fecha de compilación",

	// CLI — messages
	"msg.no_aliases":     "aún no hay aliases — prueba `aka add <nombre> <comando>`",
	"msg.added":          "añadido %s → %s",
	"msg.removed":        "eliminado %s",
	"msg.updated":        "actualizado %s",
	"msg.aborted":        "cancelado",
	"msg.no_changes":     "sin cambios",
	"msg.confirm_delete": "¿eliminar %s (%s)? [s/N] ",
	"msg.install_hint":   "abre una terminal nueva o haz `source` del rc para cargar los aliases.",
	"msg.import_summary": "importados %d nuevos, actualizados %d, omitidos %d",

	// TUI
	"tui.title":                "aka — aliases",
	"tui.form.add":             "Añadir alias",
	"tui.form.edit":            "Editar alias",
	"tui.form.field.name":      "nombre",
	"tui.form.field.command":   "comando",
	"tui.form.field.shells":    "shells (zsh,bash,fish — vacío = todos)",
	"tui.form.field.tags":      "etiquetas (separadas por coma)",
	"tui.form.field.desc":      "descripción",
	"tui.form.hint":            "tab/shift+tab mover · enter siguiente/enviar · ctrl+s enviar · esc cancelar",
	"tui.list.hint":            "a añadir · e editar · d borrar · enter copiar · / filtrar · ? ayuda · q salir",
	"tui.confirm.title":        "¿Eliminar %s?",
	"tui.confirm.hint":         "s = sí   ·   n/esc = no",
	"tui.help.title":           "aka — teclas",
	"tui.help.return":          "pulsa cualquier tecla para volver",
	"tui.status.saved":         "guardado: %s",
	"tui.status.deleted":       "eliminado: %s",
	"tui.status.copied":        "copiado: %s",
	"tui.status.clipboard_err": "error de portapapeles: %s",
}
