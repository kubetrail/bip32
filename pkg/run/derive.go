package run

import (
	"fmt"

	"github.com/kubetrail/bip32/pkg/flags"
	"github.com/kubetrail/bip32/pkg/keys"
	"github.com/kubetrail/bip39/pkg/prompts"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Derive(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.DerivationPath, cmd.Flag(flags.DerivationPath))
	derivationPath := viper.GetString(flags.DerivationPath)

	prompt, err := prompts.Status()
	if err != nil {
		return fmt.Errorf("failed to get prompt status: %w", err)
	}

	var keyString string

	if len(args) == 0 {
		if prompt {
			if err := keys.Prompt(cmd.OutOrStdout()); err != nil {
				return fmt.Errorf("failed to prompt for key: %w", err)
			}
		}

		keyString, err = keys.Read(cmd.InOrStdin())
		if err != nil {
			return fmt.Errorf("failed to read key from input: %w", err)
		}
	} else {
		keyString = args[0]
	}

	key, err := keys.Derive(keyString, derivationPath)
	if err != nil {
		return fmt.Errorf("failed to derive key: %w", err)
	}

	if prompt {
		if _, err := fmt.Fprintln(cmd.OutOrStdout(), key.Print()); err != nil {
			return fmt.Errorf("failed to write key to output: %w", err)
		}

		return nil
	}

	if _, err := fmt.Fprintln(cmd.OutOrStdout(), key); err != nil {
		return fmt.Errorf("failed to write key to output: %w", err)
	}

	return nil
}
