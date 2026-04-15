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

	// TUI — header / footers
	"tui.header.count":    "%d aliases",
	"tui.footer.list":     "a add · e edit · d delete · enter copy · / filter · t by tag · o settings · L lang · ? help · q quit",
	"tui.footer.filter":   "type to filter · enter to confirm · esc to clear",
	"tui.footer.form":     "tab/shift+tab move · enter next · ctrl+s submit · esc cancel",
	"tui.footer.confirm":  "y/enter yes · n/esc no · ←/→/tab toggle",
	"tui.footer.back":     "press any key to return",
	"tui.footer.settings": "↑/↓ pick · ←/→ cycle · esc back",

	// TUI — empty / placeholders
	"tui.empty":      "no aliases yet — press a to add one",
	"tui.no_matches": "no matches",

	// TUI — preview pane
	"tui.preview.created": "created %s",
	"tui.preview.last":    "last used %s",

	// TUI — form
	"tui.form.add":            "Add alias",
	"tui.form.edit":           "Edit alias",
	"tui.field.name":          "name",
	"tui.field.command":       "command",
	"tui.field.shells":        "shells",
	"tui.field.tags":          "tags",
	"tui.field.desc":          "description",
	"tui.placeholder.name":    "(alias name)",
	"tui.placeholder.command": "(shell command)",
	"tui.placeholder.shells":  "(zsh,bash,fish — blank = all)",
	"tui.placeholder.tags":    "(comma-separated)",
	"tui.placeholder.desc":    "(optional description)",

	// TUI — confirm
	"tui.confirm.title": "Delete alias?",
	"tui.confirm.yes":   "Yes",
	"tui.confirm.no":    "No",

	// TUI — help screen
	"tui.help.title": "aka — keyboard shortcuts",
	"help.move":      "move up / down",
	"help.jump":      "jump to top / bottom",
	"help.filter":    "filter aliases",
	"help.esc":       "clear filter / go back",
	"help.add":       "add a new alias",
	"help.edit":      "edit selected alias",
	"help.delete":    "delete selected alias",
	"help.copy":      "copy command to clipboard",
	"help.tag":       "filter by selected tag",
	"help.settings":  "open settings",
	"help.lang":      "toggle language (en ↔ es)",
	"help.toggle":    "toggle this help",
	"help.quit":      "quit",
	"help.theme":     "theme:",
	"help.lang_cur":  "language:",

	// TUI — settings
	"tui.settings.title":      "Settings",
	"tui.settings.desc":       "Cycle values with ← / → (or enter); changes are saved to config.toml.",
	"settings.language":       "language",
	"settings.theme":          "theme",
	"settings.confirm_delete": "confirm before delete",
	"settings.on":             "on",
	"settings.off":            "off",

	// TUI — status messages
	"tui.status.saved":     "saved %s",
	"tui.status.deleted":   "deleted %s",
	"tui.status.copied":    "copied: %s",
	"tui.status.copyfail":  "clipboard error: %s",
	"tui.status.delfail":   "delete failed: %s",
	"tui.status.tag_set":   "tag filter: #%s",
	"tui.status.tag_clear": "tag filter cleared",
	"tui.status.lang":      "language: %s",

	// TUI — form errors (i18n keys used as errMsg)
	"err.empty_name":    "name is required",
	"err.empty_command": "command is required",
	"err.validation":    "validation failed — see status line",
	"err.rename":        "rename failed — see status line",
	"err.exists":        "an alias with that name already exists",
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

	// TUI — header / footers
	"tui.header.count":    "%d aliases",
	"tui.footer.list":     "a añadir · e editar · d borrar · enter copiar · / filtrar · t por tag · o ajustes · L idioma · ? ayuda · q salir",
	"tui.footer.filter":   "escribe para filtrar · enter para confirmar · esc para limpiar",
	"tui.footer.form":     "tab/shift+tab mover · enter siguiente · ctrl+s enviar · esc cancelar",
	"tui.footer.confirm":  "s/enter sí · n/esc no · ←/→/tab alternar",
	"tui.footer.back":     "pulsa cualquier tecla para volver",
	"tui.footer.settings": "↑/↓ elegir · ←/→ cambiar · esc volver",

	// TUI — empty / placeholders
	"tui.empty":      "aún no hay aliases — pulsa a para añadir uno",
	"tui.no_matches": "sin resultados",

	// TUI — preview pane
	"tui.preview.created": "creado %s",
	"tui.preview.last":    "última vez %s",

	// TUI — form
	"tui.form.add":            "Añadir alias",
	"tui.form.edit":           "Editar alias",
	"tui.field.name":          "nombre",
	"tui.field.command":       "comando",
	"tui.field.shells":        "shells",
	"tui.field.tags":          "etiquetas",
	"tui.field.desc":          "descripción",
	"tui.placeholder.name":    "(nombre del alias)",
	"tui.placeholder.command": "(comando shell)",
	"tui.placeholder.shells":  "(zsh,bash,fish — vacío = todos)",
	"tui.placeholder.tags":    "(separadas por coma)",
	"tui.placeholder.desc":    "(descripción opcional)",

	// TUI — confirm
	"tui.confirm.title": "¿Eliminar alias?",
	"tui.confirm.yes":   "Sí",
	"tui.confirm.no":    "No",

	// TUI — help screen
	"tui.help.title": "aka — atajos de teclado",
	"help.move":      "mover arriba / abajo",
	"help.jump":      "ir al principio / final",
	"help.filter":    "filtrar aliases",
	"help.esc":       "limpiar filtro / volver",
	"help.add":       "añadir un alias",
	"help.edit":      "editar alias seleccionado",
	"help.delete":    "borrar alias seleccionado",
	"help.copy":      "copiar comando al portapapeles",
	"help.tag":       "filtrar por etiqueta seleccionada",
	"help.settings":  "abrir ajustes",
	"help.lang":      "cambiar idioma (en ↔ es)",
	"help.toggle":    "mostrar/ocultar esta ayuda",
	"help.quit":      "salir",
	"help.theme":     "tema:",
	"help.lang_cur":  "idioma:",

	// TUI — settings
	"tui.settings.title":      "Ajustes",
	"tui.settings.desc":       "Cambia los valores con ← / → (o enter); los cambios se guardan en config.toml.",
	"settings.language":       "idioma",
	"settings.theme":          "tema",
	"settings.confirm_delete": "confirmar antes de borrar",
	"settings.on":             "sí",
	"settings.off":            "no",

	// TUI — status messages
	"tui.status.saved":     "guardado %s",
	"tui.status.deleted":   "eliminado %s",
	"tui.status.copied":    "copiado: %s",
	"tui.status.copyfail":  "error de portapapeles: %s",
	"tui.status.delfail":   "error al borrar: %s",
	"tui.status.tag_set":   "filtro de tag: #%s",
	"tui.status.tag_clear": "filtro de tag limpiado",
	"tui.status.lang":      "idioma: %s",

	// TUI — form errors (i18n keys used as errMsg)
	"err.empty_name":    "el nombre es obligatorio",
	"err.empty_command": "el comando es obligatorio",
	"err.validation":    "validación fallida — revisa la línea de estado",
	"err.rename":        "renombrado fallido — revisa la línea de estado",
	"err.exists":        "ya existe un alias con ese nombre",
}
