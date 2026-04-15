// Package cli wires the Cobra command tree.
package cli

import "github.com/spf13/cobra"

func newRoot() *cobra.Command {
	root := &cobra.Command{
		Use:           "aka",
		Short:         "Manage shell aliases from the terminal with a TUI",
		Long:          "aka — manage shell aliases from the terminal with a TUI. Multi-shell (zsh, bash, fish).",
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.AddCommand(newVersionCmd())
	return root
}

// Execute runs the CLI. Returns a non-nil error on failure.
func Execute() error {
	return newRoot().Execute()
}
