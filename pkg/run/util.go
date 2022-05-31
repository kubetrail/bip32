package run

import (
	"strings"

	"github.com/kubetrail/bip32/pkg/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type persistentFlagValues struct {
	OutputFormat string `json:"outputFormat,omitempty"`
}

func getPersistentFlags(cmd *cobra.Command) persistentFlagValues {
	rootCmd := cmd.Root().PersistentFlags()

	_ = viper.BindPFlag(flags.OutputFormat, rootCmd.Lookup(flags.OutputFormat))
	outputFormat := strings.ToLower(viper.GetString(flags.OutputFormat))

	return persistentFlagValues{
		OutputFormat: outputFormat,
	}
}
