package run

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/kubetrail/bip32/pkg/flags"
	"github.com/kubetrail/bip32/pkg/keys"
	"github.com/kubetrail/bip39/pkg/prompts"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func Decode(cmd *cobra.Command, args []string) error {
	persistentFlags := getPersistentFlags(cmd)

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

	var keyFormat string
	if keys.IsValidBase58String(keyString) {
		keyFormat = keys.KeyFormatB58
	}

	if len(keyFormat) == 0 {
		if _, err := hex.DecodeString(keyString); err == nil {
			keyFormat = keys.KeyFormatHex
		}
	}

	var key *keys.Key
	switch keyFormat {
	case keys.KeyFormatB58:
		switch len(base58.Decode(keyString)) {
		case 38: // treat input as a wif private key
			key, err = keys.DecodePrivateWifKey(keyString)
			if err != nil {
				return fmt.Errorf("failed to decode private wif key: %w", err)
			}
		case 82: // treat input as extended key, private or public
			key, err = keys.DecodeExtendedKey(keyString)
			if err != nil {
				return fmt.Errorf("failed to serialize key: %w", err)
			}
		default:
			return fmt.Errorf("invalid input key length, needs to be either 38 bytes (prvKeyWif) or 82 bytes (xPrv, xPub) long")
		}
	case keys.KeyFormatHex:
		key, err = keys.DecodePublicHex(keyString)
		if err != nil {
			return fmt.Errorf("failed to decode public key hex: %w", err)
		}
	default:
		return fmt.Errorf("invalid base58 or hex key")
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
