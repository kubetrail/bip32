/*
Copyright © 2022 kubetrail.io authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/kubetrail/bip32/pkg/flags"
	"github.com/kubetrail/bip32/pkg/keys"
	"github.com/kubetrail/bip32/pkg/run"
	"github.com/kubetrail/bip39/pkg/mnemonics"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate keys from mnemonic",
	Long: `This command generates private/public keys
from mnemonic and optional secret passphrase per BIP-32 spec.

Alternatively, a seed in hex format can be provided bypassing
all mnemonic related computation and be directly used for
key generation

The keys are generated based on a chain derivation path
Path     |   Remark
---------|--------------------------------------------------------------
m        |   Master key (aka root key)
m/0      |   First child of master key
m/0'     |   First hardened child of master key
m/0/0    |   First child of first child of master key
m/0'/0   |   First child of first hardened child of master key
m/0/0'   |   First hardened child of first child of master key
m/0'/0'  |   First hardened child of first hardened child of master key'

Mnemonic language can be specified from the following list:
1. English (default)
2. Japanese
3. ChineseSimplified
4. ChineseTraditional
5. Czech
6. French
7. Italian
8. Korean
9. Spanish

BIP-39 proposal: https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki

Please note that same keys will be generated for mnemonics from different languages
if the underlying entropy is the same. In other words, keys are always
generated after translating input mnemonic to English.

Verify output using: http://bip32.org/

`,
	RunE: run.Gen,
	Args: cobra.MaximumNArgs(24),
}

func init() {
	rootCmd.AddCommand(genCmd)
	f := genCmd.Flags()

	f.String(flags.DerivationPath, flags.DerivationPathAuto, "Chain Derivation path")
	f.Bool(flags.UsePassphrase, false, "Prompt for secret passphrase")
	f.Bool(flags.InputHexSeed, false, "Treat input as hex seed instead of mnemonic")
	f.String(flags.MnemonicLanguage, mnemonics.LanguageEnglish, "Mnemonic language")
	f.Bool(flags.SkipMnemonicValidation, false, "Skip mnemonic validation")
	// https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#serialization-format
	f.String(flags.Network, flags.NetworkMainnet, "Network: mainnet or testnet")
	f.String(flags.AddrType, keys.AddrTypeP2pkhOrP2sh, "Script type")
	f.Bool(flags.ShowAllKeys, false, "Show all keys")

	_ = genCmd.RegisterFlagCompletionFunc(
		flags.Network,
		func(
			cmd *cobra.Command,
			args []string,
			toComplete string,
		) (
			[]string,
			cobra.ShellCompDirective,
		) {
			return []string{
					flags.NetworkMainnet,
					flags.NetworkTestnet,
				},
				cobra.ShellCompDirectiveDefault
		},
	)

	_ = genCmd.RegisterFlagCompletionFunc(
		flags.MnemonicLanguage,
		func(
			cmd *cobra.Command,
			args []string,
			toComplete string,
		) (
			[]string,
			cobra.ShellCompDirective,
		) {
			return []string{
					mnemonics.LanguageEnglish,
					mnemonics.LanguageJapanese,
					mnemonics.LanguageChineseSimplified,
					mnemonics.LanguageChineseTraditional,
					mnemonics.LanguageCzech,
					mnemonics.LanguageFrench,
					mnemonics.LanguageItalian,
					mnemonics.LanguageKorean,
					mnemonics.LanguageSpanish,
				},
				cobra.ShellCompDirectiveDefault
		},
	)

	_ = genCmd.RegisterFlagCompletionFunc(
		flags.DerivationPath,
		func(
			cmd *cobra.Command,
			args []string,
			toComplete string,
		) (
			[]string,
			cobra.ShellCompDirective,
		) {
			return []string{
					flags.DerivationPath0,
					flags.DerivationPath1,
					flags.DerivationPath2,
					flags.DerivationPath3,
					flags.DerivationPath4,
					flags.DerivationPath5,
					flags.DerivationPath6,
					flags.DerivationPath7,
					flags.DerivationPath8,
				},
				cobra.ShellCompDirectiveDefault
		},
	)

	_ = genCmd.RegisterFlagCompletionFunc(
		flags.AddrType,
		func(
			cmd *cobra.Command,
			args []string,
			toComplete string,
		) (
			[]string,
			cobra.ShellCompDirective,
		) {
			return []string{
					keys.AddrTypeLegacy,
					keys.AddrTypeP2sh,
					keys.AddrTypeSegWitCompatible,
					keys.AddrTypeSegWitNative,
					keys.AddrTypeBech32,
					keys.AddrTypeP2pkhOrP2sh,
					keys.AddrTypeP2wpkhP2sh,
					keys.AddrTypeP2wshP2sh,
					keys.AddrTypeP2wpkh,
					keys.AddrTypeP2wsh,
				},
				cobra.ShellCompDirectiveDefault
		},
	)
}
