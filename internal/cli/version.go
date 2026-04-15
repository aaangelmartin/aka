package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aaangelmartin/aka/internal/buildinfo"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version, commit, and build date",
		RunE: func(cmd *cobra.Command, _ []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), "aka %s (commit %s, built %s)\n",
				buildinfo.Version, buildinfo.Commit, buildinfo.Date)
			return nil
		},
	}
}
