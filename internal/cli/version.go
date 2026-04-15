package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aaangelmartin/aka/internal/buildinfo"
	"github.com/aaangelmartin/aka/internal/i18n"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: i18n.T("cli.version.short"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			fmt.Fprintf(cmd.OutOrStdout(), "aka %s (commit %s, built %s)\n",
				buildinfo.Version, buildinfo.Commit, buildinfo.Date)
			return nil
		},
	}
}
