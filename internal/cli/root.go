// Package cli wires the Cobra command tree.
package cli

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/aaangelmartin/aka/internal/config"
	"github.com/aaangelmartin/aka/internal/i18n"
	"github.com/aaangelmartin/aka/internal/tui"
)

func newRoot() *cobra.Command {
	root := &cobra.Command{
		Use:           "aka",
		Short:         i18n.T("cli.root.short"),
		Long:          i18n.T("cli.root.long"),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(_ *cobra.Command, _ []string) error {
			return tui.Run()
		},
	}
	root.PersistentFlags().String("lang", "", "force language (en|es); default: auto-detect")
	root.AddCommand(
		newAddCmd(),
		newLsCmd(),
		newRmCmd(),
		newEditCmd(),
		newConfigCmd(),
		newImportCmd(),
		newExportCmd(),
		newInstallCmd(),
		newUninstallCmd(),
		newVersionCmd(),
	)
	return root
}

// Execute runs the CLI. Returns a non-nil error on failure.
func Execute() error {
	setupI18n()
	return newRoot().Execute()
}

// setupI18n resolves the active language from (in order) the --lang flag,
// the config file, and $LANG. It runs before the Cobra tree is built so
// that command Short/Long strings resolve to the chosen language.
func setupI18n() {
	flagLang := peekLangFlag(os.Args[1:])
	cfgLang := ""
	if path, err := config.ConfigPath(); err == nil {
		if cfg, err := config.Load(path); err == nil {
			cfgLang = cfg.Language
		}
	}
	i18n.Set(i18n.Detect(cfgLang, flagLang))
}

// peekLangFlag finds `--lang X` or `--lang=X` in args without depending on
// cobra. Returns "" if not present.
func peekLangFlag(args []string) string {
	for i, a := range args {
		switch {
		case a == "--lang" && i+1 < len(args):
			return args[i+1]
		case len(a) > len("--lang=") && a[:len("--lang=")] == "--lang=":
			return a[len("--lang="):]
		}
	}
	return ""
}
