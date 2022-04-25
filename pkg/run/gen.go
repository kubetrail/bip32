package run

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"syscall"

	"github.com/kubetrail/bip32/pkg/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"github.com/wemeetagain/go-hdwallet"
	"golang.org/x/term"
)

func Gen(cmd *cobra.Command, args []string) error {
	_ = viper.BindPFlag(flags.UsePassphrase, cmd.Flag(flags.UsePassphrase))
	_ = viper.BindPFlag(flags.SkipMnemonicValidation, cmd.Flag(flags.SkipMnemonicValidation))
	_ = viper.BindPFlag(flags.DerivationPath, cmd.Flag(flags.DerivationPath))
	_ = viper.BindPFlag(flags.InputHexSeed, cmd.Flag(flags.InputHexSeed))
	_ = viper.BindPFlag(flags.Network, cmd.Flag(flags.Network))

	usePassphrase := viper.GetBool(flags.UsePassphrase)
	skipMnemonicValidation := viper.GetBool(flags.SkipMnemonicValidation)
	derivationPath := viper.GetString(flags.DerivationPath)
	inputHexSeed := viper.GetBool(flags.InputHexSeed)
	network := viper.GetString(flags.Network)

	// setup key versions based on network
	switch network {
	case NetworkMainnet:
		bip32.PrivateWalletVersion = hdwallet.Private
		bip32.PublicWalletVersion = hdwallet.Public
	case NetworkTestnet:
		bip32.PrivateWalletVersion = hdwallet.TestPrivate
		bip32.PublicWalletVersion = hdwallet.TestPublic
		hdwallet.Private = hdwallet.TestPrivate
		hdwallet.Public = hdwallet.TestPublic
	default:
		return fmt.Errorf("invalid network: %s", network)
	}

	prompt, err := getPromptStatus()
	if err != nil {
		return fmt.Errorf("failed to get prompt status: %w", err)
	}

	var passphrase []byte
	var seed []byte

	if inputHexSeed && usePassphrase {
		return fmt.Errorf("cannot use passphrase when entering seed")
	}

	if inputHexSeed && skipMnemonicValidation {
		return fmt.Errorf("dont use --skip-mnemonic-validation when entering seed")
	}

	if !inputHexSeed {
		if prompt {
			if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Enter mnemonic: "); err != nil {
				return fmt.Errorf("failed to write to output: %w", err)
			}
		}

		inputReader := bufio.NewReader(cmd.InOrStdin())
		mnemonic, err := inputReader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read mnemonic from input: %w", err)
		}
		mnemonic = strings.Trim(mnemonic, "\n")

		if !skipMnemonicValidation && !bip39.IsMnemonicValid(mnemonic) {
			return fmt.Errorf("mnemonic is invalid or please use --skip-mnemonic-validation flag")
		}

		if usePassphrase {
			if prompt {
				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Enter secret passphrase: "); err != nil {
					return fmt.Errorf("failed to write to output: %w", err)
				}
			}

			passphrase, err = term.ReadPassword(syscall.Stdin)
			if err != nil {
				return fmt.Errorf("failed to read secret passphrase from input: %w", err)
			}
			if _, err := fmt.Fprintln(cmd.OutOrStdout()); err != nil {
				return fmt.Errorf("failed to write to output: %w", err)
			}

			if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Enter secret passphrase again: "); err != nil {
				return fmt.Errorf("failed to write to output: %w", err)
			}

			passphraseConfirm, err := term.ReadPassword(syscall.Stdin)
			if err != nil {
				return fmt.Errorf("failed to read secret passphrase from input: %w", err)
			}
			if _, err := fmt.Fprintln(cmd.OutOrStdout()); err != nil {
				return fmt.Errorf("failed to write to output: %w", err)
			}

			if !bytes.Equal(passphrase, passphraseConfirm) {
				return fmt.Errorf("passphrases do not match")
			}
		}
		seed = bip39.NewSeed(mnemonic, string(passphrase))
	} else {
		if prompt {
			if _, err := fmt.Fprintf(cmd.OutOrStdout(), "Enter seed in hex: "); err != nil {
				return fmt.Errorf("failed to write to output: %w", err)
			}
		}

		inputReader := bufio.NewReader(cmd.InOrStdin())
		hexSeed, err := inputReader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read hex seed from input: %w", err)
		}
		hexSeed = strings.Trim(hexSeed, "\n")

		seed, err = hex.DecodeString(hexSeed)
		if err != nil {
			return fmt.Errorf("invalid hex seed string: %w", err)
		}
	}

	key, err := bip32.NewMasterKey(seed)
	if err != nil {
		return fmt.Errorf("failed to generate root key: %w", err)
	}

	derivationPath = strings.Trim(derivationPath, "/")
	parts := strings.Split(derivationPath, "/")
	if len(parts) == 0 || parts[0] != "m" {
		return fmt.Errorf("invalid derivation path: %s", derivationPath)
	}

	for i, part := range parts {
		if i == 0 {
			continue
		}

		if len(part) == 0 {
			return fmt.Errorf("invalid derivation path at index %d: %s", i, derivationPath)
		}

		var idx uint32
		if part[len(part)-1] == '\'' || part[len(part)-1] == 'h' || part[len(part)-1] == 'H' {
			idx = bip32.FirstHardenedChild
			part = part[:len(part)-1]
		}

		index, err := strconv.ParseInt(part, 10, 64)
		if err != nil || index < 0 {
			return fmt.Errorf("invalid derivation path at index %d: %s, %w", i, derivationPath, err)
		}

		idx += uint32(index)
		key, err = key.NewChildKey(idx)
		if err != nil {
			return fmt.Errorf("failed to generate %d child key: %w", i, err)
		}
	}

	outPrv := fmt.Sprintf("%s", key)
	outPub := fmt.Sprintf("%s", key.PublicKey())

	wallet, err := hdwallet.StringWallet(outPrv)
	if err != nil {
		return fmt.Errorf("failed to create hdwallet: %w", err)
	}

	if wallet.String() != outPrv {
		return fmt.Errorf("private key mismatch, \nwallet: %s\n bip32: %s", wallet.String(), outPrv)
	}

	if wallet.Pub().String() != outPub {
		return fmt.Errorf("private key mismatch, \nwallet: %s\n bip32: %s", wallet.Pub().String(), outPub)
	}

	if prompt {
		if _, err := fmt.Fprintln(cmd.OutOrStdout(), "addr:", wallet.Address()); err != nil {
			return fmt.Errorf("failed to write key to output: %w", err)
		}

		if _, err := fmt.Fprintln(cmd.OutOrStdout(), "pub:", outPub); err != nil {
			return fmt.Errorf("failed to write key to output: %w", err)
		}

		if _, err := fmt.Fprintln(cmd.OutOrStdout(), "prv:", outPrv); err != nil {
			return fmt.Errorf("failed to write key to output: %w", err)
		}

		return nil
	}

	jb, err := json.Marshal(
		struct {
			Addr string `json:"addr,omitempty"`
			Prv  string `json:"prv,omitempty"`
			Pub  string `json:"pub,omitempty"`
		}{
			Addr: wallet.Address(),
			Prv:  outPrv,
			Pub:  outPub,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to serialize output: %w", err)
	}

	if _, err := fmt.Fprintln(cmd.OutOrStdout(), string(jb)); err != nil {
		return fmt.Errorf("failed to write key to output: %w", err)
	}

	return nil
}
