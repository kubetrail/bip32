package run

import (
	"encoding/json"
	"fmt"

	"github.com/kubetrail/bip32/pkg/flags"
	"github.com/kubetrail/bip32/pkg/keys"
	"github.com/kubetrail/bip39/pkg/prompts"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func Derive(cmd *cobra.Command, args []string) error {
	persistentFlags := getPersistentFlags(cmd)

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

	switch persistentFlags.OutputFormat {
	case flags.OutputFormatNative, flags.OutputFormatYaml:
		jb, err := yaml.Marshal(key)
		if err != nil {
			return fmt.Errorf("failed to serialize output to json: %w", err)
		}
		if _, err := fmt.Fprint(cmd.OutOrStdout(), string(jb)); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}
	case flags.OutputFormatJson:
		jb, err := json.Marshal(key)
		if err != nil {
			return fmt.Errorf("failed to serialize output to json: %w", err)
		}
		if _, err := fmt.Fprintln(cmd.OutOrStdout(), string(jb)); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}
	}

	return nil
}
