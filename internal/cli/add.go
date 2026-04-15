package cli

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/aaangelmartin/aka/internal/alias"
)

func newAddCmd() *cobra.Command {
	var (
		shells      []string
		tags        []string
		description string
		force       bool
	)
	cmd := &cobra.Command{
		Use:   "add <name> <command>...",
		Short: "Add a new alias",
		Long:  "Add a new alias. The command may be passed as multiple arguments; they are joined with a single space.",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			command := strings.Join(args[1:], " ")
			for _, sh := range shells {
				if !alias.IsValidShell(sh) {
					return fmt.Errorf("unknown shell %q (allowed: zsh, bash, fish)", sh)
				}
			}
			a := alias.Alias{
				Name:        name,
				Command:     command,
				Shells:      shells,
				Tags:        tags,
				Description: description,
				CreatedAt:   time.Now().UTC(),
			}
			if err := alias.Validate(a); err != nil {
				return err
			}
			sess, err := openSession()
			if err != nil {
				return err
			}
			if force {
				sess.store.Set(a)
			} else if err := sess.store.Put(a); err != nil {
				return err
			}
			if err := sess.commit(); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "added %s → %s\n", name, command)
			return nil
		},
	}
	cmd.Flags().StringSliceVarP(&shells, "shell", "s", nil, "shells to emit for (zsh,bash,fish); default: all")
	cmd.Flags().StringSliceVarP(&tags, "tag", "t", nil, "tags (repeatable or comma-separated)")
	cmd.Flags().StringVarP(&description, "desc", "d", "", "description")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "overwrite if an alias with the same name exists")
	// Flags must precede positional args so the command can contain its own
	// flags (e.g. `aka add ccauto 'claude --dangerously-skip-permissions'`).
	cmd.Flags().SetInterspersed(false)
	return cmd
}
