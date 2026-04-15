package cli

import (
	"encoding/json"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/aaangelmartin/aka/internal/alias"
)

func newLsCmd() *cobra.Command {
	var (
		jsonOut bool
		tag     string
		shell   string
	)
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List aliases",
		RunE: func(cmd *cobra.Command, _ []string) error {
			sess, err := openSession()
			if err != nil {
				return err
			}
			list := sess.store.List()
			list = filterByTag(list, tag)
			list = filterByShell(list, shell)

			if jsonOut {
				return json.NewEncoder(cmd.OutOrStdout()).Encode(list)
			}
			if len(list) == 0 {
				fmt.Fprintln(cmd.OutOrStdout(), "no aliases yet — try `aka add <name> <command>`")
				return nil
			}
			w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
			fmt.Fprintln(w, "NAME\tSHELLS\tTAGS\tCOMMAND")
			for _, a := range list {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
					a.Name,
					strings.Join(a.ShellsOrAll(), ","),
					strings.Join(a.Tags, ","),
					a.Command,
				)
			}
			return w.Flush()
		},
	}
	cmd.Flags().BoolVarP(&jsonOut, "json", "j", false, "output JSON")
	cmd.Flags().StringVarP(&tag, "tag", "t", "", "filter by tag")
	cmd.Flags().StringVarP(&shell, "shell", "s", "", "filter by shell (zsh/bash/fish)")
	return cmd
}

func filterByTag(in []alias.Alias, tag string) []alias.Alias {
	if tag == "" {
		return in
	}
	out := in[:0:0]
	for _, a := range in {
		for _, t := range a.Tags {
			if t == tag {
				out = append(out, a)
				break
			}
		}
	}
	return out
}

func filterByShell(in []alias.Alias, shell string) []alias.Alias {
	if shell == "" {
		return in
	}
	out := in[:0:0]
	for _, a := range in {
		if a.TargetsShell(shell) {
			out = append(out, a)
		}
	}
	return out
}
