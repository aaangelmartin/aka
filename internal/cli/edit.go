package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aaangelmartin/aka/internal/alias"
)

func newEditCmd() *cobra.Command {
	var (
		newName     string
		newCommand  string
		newDesc     string
		setTags     []string
		setShells   []string
		clearTags   bool
		clearDesc   bool
		clearShells bool
	)
	cmd := &cobra.Command{
		Use:   "edit <name>",
		Short: "Edit an alias in place",
		Args:  cobra.ExactArgs(1),
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
			changed := false
			if newCommand != "" {
				a.Command = newCommand
				changed = true
			}
			if cmd.Flags().Changed("desc") {
				a.Description = newDesc
				changed = true
			}
			if clearDesc {
				a.Description = ""
				changed = true
			}
			if len(setTags) > 0 {
				a.Tags = setTags
				changed = true
			}
			if clearTags {
				a.Tags = nil
				changed = true
			}
			if len(setShells) > 0 {
				for _, sh := range setShells {
					if !alias.IsValidShell(sh) {
						return fmt.Errorf("unknown shell %q", sh)
					}
				}
				a.Shells = setShells
				changed = true
			}
			if clearShells {
				a.Shells = nil
				changed = true
			}
			if newName != "" && newName != a.Name {
				if err := sess.store.Rename(a.Name, newName); err != nil {
					return err
				}
				a.Name = newName
				changed = true
			}
			sess.store.Set(a)
			if err := alias.Validate(a); err != nil {
				return err
			}
			if !changed {
				fmt.Fprintf(cmd.OutOrStdout(), "no changes\n")
				return nil
			}
			if err := sess.commit(); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "updated %s\n", a.Name)
			return nil
		},
	}
	cmd.Flags().StringVarP(&newName, "name", "n", "", "new name")
	cmd.Flags().StringVarP(&newCommand, "command", "c", "", "new command")
	cmd.Flags().StringVarP(&newDesc, "desc", "d", "", "new description (empty string clears)")
	cmd.Flags().BoolVar(&clearDesc, "no-desc", false, "clear the description")
	cmd.Flags().StringSliceVarP(&setTags, "tag", "t", nil, "replace tags")
	cmd.Flags().BoolVar(&clearTags, "no-tags", false, "clear tags")
	cmd.Flags().StringSliceVarP(&setShells, "shell", "s", nil, "replace shells (zsh,bash,fish)")
	cmd.Flags().BoolVar(&clearShells, "all-shells", false, "target every shell (clears shells list)")
	return cmd
}
