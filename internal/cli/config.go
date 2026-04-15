package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"github.com/aaangelmartin/aka/internal/alias"
	"github.com/aaangelmartin/aka/internal/config"
)

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config [key] [value]",
		Short: "Read or write a config key",
		Long: `Read or write a config key (no args prints everything).

Supported keys:
  language         auto | en | es
  default_shells   comma-separated subset of zsh,bash,fish
  confirm_delete   true | false
  theme            free-form string (validated by TUI)`,
		Args: cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfgPath, err := config.ConfigPath()
			if err != nil {
				return err
			}
			cfg, err := config.Load(cfgPath)
			if err != nil {
				return err
			}
			switch len(args) {
			case 0:
				fmt.Fprintf(cmd.OutOrStdout(), "config: %s\n", cfgPath)
				fmt.Fprintf(cmd.OutOrStdout(), "  language        = %s\n", cfg.Language)
				fmt.Fprintf(cmd.OutOrStdout(), "  default_shells  = %s\n", strings.Join(cfg.DefaultShells, ","))
				fmt.Fprintf(cmd.OutOrStdout(), "  confirm_delete  = %t\n", cfg.ConfirmDelete)
				fmt.Fprintf(cmd.OutOrStdout(), "  theme           = %s\n", cfg.Theme)
				return nil
			case 1:
				v, err := getKey(cfg, args[0])
				if err != nil {
					return err
				}
				fmt.Fprintln(cmd.OutOrStdout(), v)
				return nil
			case 2:
				if err := setKey(&cfg, args[0], args[1]); err != nil {
					return err
				}
				if err := config.Save(cfgPath, cfg); err != nil {
					return err
				}
				fmt.Fprintf(cmd.OutOrStdout(), "set %s = %s\n", args[0], args[1])
				return nil
			}
			return nil
		},
	}
	return cmd
}

func getKey(cfg config.Config, key string) (string, error) {
	switch key {
	case "language":
		return cfg.Language, nil
	case "default_shells":
		return strings.Join(cfg.DefaultShells, ","), nil
	case "confirm_delete":
		return strconv.FormatBool(cfg.ConfirmDelete), nil
	case "theme":
		return cfg.Theme, nil
	}
	return "", fmt.Errorf("unknown config key %q", key)
}

func setKey(cfg *config.Config, key, val string) error {
	switch key {
	case "language":
		if val != "auto" && val != "en" && val != "es" {
			return fmt.Errorf("language must be auto|en|es")
		}
		cfg.Language = val
	case "default_shells":
		parts := strings.Split(val, ",")
		for i, p := range parts {
			p = strings.TrimSpace(p)
			if !alias.IsValidShell(p) {
				return fmt.Errorf("unknown shell %q", p)
			}
			parts[i] = p
		}
		cfg.DefaultShells = parts
	case "confirm_delete":
		b, err := strconv.ParseBool(val)
		if err != nil {
			return fmt.Errorf("confirm_delete must be true|false")
		}
		cfg.ConfirmDelete = b
	case "theme":
		cfg.Theme = val
	default:
		return fmt.Errorf("unknown config key %q", key)
	}
	return nil
}
