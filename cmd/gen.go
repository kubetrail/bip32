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
	"github.com/kubetrail/bip32/pkg/run"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate key from mnemonic",
	Long: `This command generates private/public keys
from mnemonic and optional secret passphrase per BIP-32 spec.

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

`,
	RunE: run.Gen,
}

func init() {
	rootCmd.AddCommand(genCmd)
	f := genCmd.Flags()

	f.String(flags.DerivationPath, "m", "Chain Derivation path")
	f.Bool(flags.UsePassphrase, false, "Prompt for secret passphrase")
	f.Bool(flags.InputHexSeed, false, "Treat input as hex seed instead of mnemonic")
	f.Bool(flags.SkipMnemonicValidation, false, "Skip mnemonic validation")
}
