package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/aaangelmartin/aka/internal/alias"
	"github.com/aaangelmartin/aka/internal/i18n"
	"github.com/aaangelmartin/aka/internal/shellint"
)

func newImportCmd() *cobra.Command {
	var (
		merge  bool
		fromRC bool
	)
	cmd := &cobra.Command{
		Use:   "import <file>",
		Short: i18n.T("cli.import.short"),
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			var in []alias.Alias
			if fromRC {
				in = shellint.ParseRC(string(b))
			} else {
				if err := json.Unmarshal(b, &in); err != nil {
					return fmt.Errorf("parse %s: %w", path, err)
				}
			}
			sess, err := openSession()
			if err != nil {
				return err
			}
			added, updated, skipped := 0, 0, 0
			for _, a := range in {
				if err := alias.Validate(a); err != nil {
					skipped++
					fmt.Fprintf(cmd.ErrOrStderr(), "skip %q: %v\n", a.Name, err)
					continue
				}
				if a.CreatedAt.IsZero() {
					a.CreatedAt = time.Now().UTC()
				}
				if _, err := sess.store.Get(a.Name); err == nil {
					if !merge {
						skipped++
						fmt.Fprintf(cmd.ErrOrStderr(), "skip %q: already exists (use --merge)\n", a.Name)
						continue
					}
					sess.store.Set(a)
					updated++
				} else {
					if err := sess.store.Put(a); err != nil {
						return err
					}
					added++
				}
			}
			if err := sess.commit(); err != nil {
				return err
			}
			fmt.Fprintln(cmd.OutOrStdout(), i18n.Tf("msg.import_summary", added, updated, skipped))
			return nil
		},
	}
	cmd.Flags().BoolVar(&merge, "merge", false, "overwrite aliases with the same name")
	cmd.Flags().BoolVar(&fromRC, "from-rc", false, "parse the file as a shell rc (zsh/bash/fish) and extract `alias ...` lines")
	return cmd
}
