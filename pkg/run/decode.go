package run

import (
	"fmt"

	"github.com/kubetrail/bip32/pkg/keys"
	"github.com/kubetrail/bip39/pkg/prompts"
	"github.com/spf13/cobra"
)

func Decode(cmd *cobra.Command, args []string) error {
	prompt, err := prompts.Status()
	if err != nil {
		return fmt.Errorf("failed to get prompt status: %w", err)
	}

	var key string

	if len(args) == 0 {
		if prompt {
			if err := keys.Prompt(cmd.OutOrStdout()); err != nil {
				return fmt.Errorf("failed to prompt for key: %w", err)
			}
		}

		key, err = keys.Read(cmd.InOrStdin())
		if err != nil {
			return fmt.Errorf("failed to read key from input: %w", err)
		}
	} else {
		key = args[0]
	}

	jb, err := keys.DecodeToJson(key)
	if err != nil {
		return fmt.Errorf("failed to serialize key: %w", err)
	}

	if _, err := fmt.Fprintln(cmd.OutOrStdout(), string(jb)); err != nil {
		return fmt.Errorf("failed to write serialized key to output: %w", err)
	}

	return nil
}
