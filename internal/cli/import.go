package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/aaangelmartin/aka/internal/alias"
)

func newImportCmd() *cobra.Command {
	var (
		merge bool
	)
	cmd := &cobra.Command{
		Use:   "import <file>",
		Short: "Import aliases from a JSON file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			var in []alias.Alias
			if err := json.Unmarshal(b, &in); err != nil {
				return fmt.Errorf("parse %s: %w", path, err)
			}
			sess, err := openSession()
			if err != nil {
				return err
			}
			added, updated := 0, 0
			for _, a := range in {
				if err := alias.Validate(a); err != nil {
					return fmt.Errorf("invalid alias %q: %w", a.Name, err)
				}
				if a.CreatedAt.IsZero() {
					a.CreatedAt = time.Now().UTC()
				}
				if _, err := sess.store.Get(a.Name); err == nil {
					if !merge {
						return fmt.Errorf("alias %q already exists (use --merge to overwrite)", a.Name)
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
			fmt.Fprintf(cmd.OutOrStdout(), "imported %d new, updated %d\n", added, updated)
			return nil
		},
	}
	cmd.Flags().BoolVar(&merge, "merge", false, "overwrite aliases with the same name")
	return cmd
}
