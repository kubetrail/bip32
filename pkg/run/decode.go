package run

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
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

	var keyFormat string
	if keys.IsValidBase58String(key) {
		keyFormat = keys.KeyFormatB58
	}

	if len(keyFormat) == 0 {
		if _, err := hex.DecodeString(key); err == nil {
			keyFormat = keys.KeyFormatHex
		}
	}

	var jb []byte
	switch keyFormat {
	case keys.KeyFormatB58:
		switch len(base58.Decode(key)) {
		case 38: // treat input as a wif private key
			jb, err = keys.DecodePrivateWifKey(key)
			if err != nil {
				return fmt.Errorf("failed to decode private wif key: %w", err)
			}
		case 82: // treat input as extended key, private or public
			jb, err = keys.DecodeExtendedKey(key)
			if err != nil {
				return fmt.Errorf("failed to serialize key: %w", err)
			}
		default:
			return fmt.Errorf("invalid input key length, needs to be either 38 bytes (prvKeyWif) or 82 bytes (xPrv, xPub) long")
		}
	case keys.KeyFormatHex:
		jb, err = keys.DecodePublicHex(key)
		if err != nil {
			return fmt.Errorf("failed to decode public key hex: %w", err)
		}
	default:
		return fmt.Errorf("invalid base58 or hex key")
	}

	if _, err := fmt.Fprintln(cmd.OutOrStdout(), string(jb)); err != nil {
		return fmt.Errorf("failed to write serialized key to output: %w", err)
	}

	return nil
}
