package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/aaangelmartin/aka/internal/i18n"
)

func newRmCmd() *cobra.Command {
	var yes bool
	cmd := &cobra.Command{
		Use:     "rm <name>",
		Aliases: []string{"remove", "delete"},
		Short:   i18n.T("cli.rm.short"),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			sess, err := openSession()
			if err != nil {
				return err
			}
			a, err := sess.store.Get(name)
			if err != nil {
				return err
			}
			if !yes && sess.cfg.ConfirmDelete {
				fmt.Fprint(cmd.OutOrStdout(), i18n.Tf("msg.confirm_delete", a.Name, a.Command))
				reader := bufio.NewReader(os.Stdin)
				line, _ := reader.ReadString('\n')
				trimmed := strings.ToLower(strings.TrimSpace(line))
				// Accept "y"/"yes" (English) or "s"/"sí"/"si" (Spanish).
				if !strings.HasPrefix(trimmed, "y") && !strings.HasPrefix(trimmed, "s") {
					fmt.Fprintln(cmd.OutOrStdout(), i18n.T("msg.aborted"))
					return nil
				}
			}
			if err := sess.store.Delete(a.Name); err != nil {
				return err
			}
			if err := sess.commit(); err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), i18n.Tf("msg.removed", a.Name))
			return nil
		},
	}
	cmd.Flags().BoolVarP(&yes, "yes", "y", false, "skip confirmation")
	return cmd
}
