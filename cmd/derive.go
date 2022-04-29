/*
Copyright © 2022 kubetrail.io authors

*/
package cmd

import (
	"github.com/kubetrail/bip32/pkg/flags"
	"github.com/kubetrail/bip32/pkg/run"
	"github.com/spf13/cobra"
)

// deriveCmd represents the derive command
var deriveCmd = &cobra.Command{
	Use:   "derive",
	Short: "Derive a child key",
	Long: `Derive a child key from private or public key.
Please note that derivation of hardened keys is only allowed
for private keys

The keys are generated based on a chain derivation path
Path     |   Remark
---------|--------------------------------------------------------------
0        |   First child of key
0'       |   First hardened child of private key
0/0      |   First child of first child of key
0'/0     |   First child of first hardened child of private key
0/0'     |   First hardened child of first child of private key
0'/0'    |   First hardened child of first hardened child of private key'

`,
	RunE: run.Derive,
}

func init() {
	rootCmd.AddCommand(deriveCmd)
	f := deriveCmd.Flags()

	f.String(flags.DerivationPath, "m", "Chain Derivation path")
}
