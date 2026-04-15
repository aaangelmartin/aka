// Package cli wires the Cobra command tree.
package cli

import (
	"github.com/spf13/cobra"

	"github.com/aaangelmartin/aka/internal/tui"
)

func newRoot() *cobra.Command {
	root := &cobra.Command{
		Use:           "aka",
		Short:         "Manage shell aliases from the terminal with a TUI",
		Long:          "aka — manage shell aliases from the terminal with a TUI. Multi-shell (zsh, bash, fish).",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(_ *cobra.Command, _ []string) error {
			return tui.Run()
		},
	}
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
	return newRoot().Execute()
}
