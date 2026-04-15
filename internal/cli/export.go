package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newExportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export [file]",
		Short: "Export aliases as JSON (stdout if no file given)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			sess, err := openSession()
			if err != nil {
				return err
			}
			b, err := json.MarshalIndent(sess.store.List(), "", "  ")
			if err != nil {
				return err
			}
			if len(args) == 0 {
				fmt.Fprintln(cmd.OutOrStdout(), string(b))
				return nil
			}
			if err := os.WriteFile(args[0], b, 0o644); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "wrote %s (%d aliases)\n", args[0], sess.store.Len())
			return nil
		},
	}
	return cmd
}
