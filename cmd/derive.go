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
	Long: `
Read more about usage on https://github.com/kubetrail/bip32
`,
	RunE: run.Derive,
}

func init() {
	rootCmd.AddCommand(deriveCmd)
	f := deriveCmd.Flags()

	f.String(flags.DerivationPath, "m", "Relative chain Derivation path")
}
